package product

import (
	"context"
	"encoding/json"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
)

const (
	elasticIndexName = "product"
)

func (p Product) SearchByIds(ctx context.Context, msg *product.SearchByIdsRequest) (*product.ProductMsg, error) {
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
			logger.Warn(err.Error(), zap.String("uuid", uuid))
			continue
		}
		texts = append(texts, t)
	}

	if len(texts) == 0 {
		return nil, status.Error(codes.NotFound, "text not recognized")
	}

	text := ""
	for _, t := range texts {
		text += " " + t.text
	}

	elasticQuery := elastic.NewBoolQuery().Should(
		elastic.NewMatchQuery("brand", text).Fuzziness("2"),
		elastic.NewMatchQuery("name", text).Fuzziness("2"),
	).MinimumNumberShouldMatch(1)

	searchRes, err := p.elasticCli.Search(elasticIndexName).
		Query(elasticQuery).
		Size(1).Do(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if searchRes.Hits == nil || len(searchRes.Hits.Hits) == 0 {
		return nil, status.Error(codes.NotFound, "product not found")
	}

	resProduct := product.ProductMsg{}
	if err := json.Unmarshal(searchRes.Hits.Hits[0].Source, &resProduct); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &resProduct, nil
}
