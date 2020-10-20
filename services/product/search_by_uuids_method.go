package product

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (p Product) SearchByUUIDs(ctx context.Context, msg *product.SearchByUUIDsRequest) (*product.ProductWithID, error) {
	logger := ctxzap.Extract(ctx)
	textRecognitionFutures := map[string]textRecognitionByUUIDFuture{}

	for _, uuid := range msg.UUIDs {
		if _, ok := textRecognitionFutures[uuid]; ok {
			logger.Warn("UUID is duplicated", zap.String("uuid", uuid))
			continue
		}
		textRecognitionFutures[uuid] = p.newTextRecognitionByUUIDFuture(ctx, uuid)
	}

	texts := make([]*recognizedText, 0, len(textRecognitionFutures))
	for uuid, f := range textRecognitionFutures {
		t, err := f.get()
		if err != nil {
			logger.Warn("text recognition future error", zap.String("uuid", uuid), zap.Error(err))
			continue
		}
		texts = append(texts, t)
	}

	if len(texts) == 0 {
		return nil, server.NewErrServer(codes.NotFound, errors.New("text not recognized"))
	}

	text := ""
	for _, t := range texts {
		text += " " + t.text
	}

	searchRes, err := p.elasticCli.Search(objectProduct).
		FetchSourceContext(
			elastic.NewFetchSourceContext(true).
				Include(fieldId)).
		Query(makeTextSearchQuery(text)).
		Size(eProductSize).Do(ctx)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	if searchRes.Hits == nil || len(searchRes.Hits.Hits) == 0 {
		return nil, server.NewErrServer(codes.NotFound, errors.New("product not found"))
	}

	eRespProduct := eProduct{}
	if err := json.Unmarshal(searchRes.Hits.Hits[0].Source, &eRespProduct); err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	resProduct, err := p.getByID(ctx, eRespProduct.Id)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, err)
	}

	return resProduct, nil
}

type eProduct struct {
	Id uint64 `json:"id"`
}

func makeTextSearchQuery(text string) *elastic.BoolQuery {
	return elastic.NewBoolQuery().Should(
		elastic.NewMatchQuery(fieldBrand, text).Fuzziness(eProductFuzziness),
		elastic.NewMatchQuery(fieldName, text).Fuzziness(eProductFuzziness),
	).MinimumNumberShouldMatch(1)
}
