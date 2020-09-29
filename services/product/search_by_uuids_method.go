package product

import (
	"context"
	"encoding/json"

	"github.com/Masterminds/squirrel"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
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
		elastic.NewMatchQuery(fieldBrand, text).Fuzziness(eProductFuzziness),
		elastic.NewMatchQuery(fieldName, text).Fuzziness(eProductFuzziness),
	).MinimumNumberShouldMatch(1)

	searchRes, err := p.elasticCli.Search(objectProduct).
		FetchSourceContext(
			elastic.NewFetchSourceContext(true).
				Include(fieldId)).
		Query(elasticQuery).
		Size(eProductSize).Do(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if searchRes.Hits == nil || len(searchRes.Hits.Hits) == 0 {
		return nil, status.Error(codes.NotFound, "product not found")
	}

	eRespProduct := eProduct{}
	if err := json.Unmarshal(searchRes.Hits.Hits[0].Source, &eRespProduct); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	row := p.gatewayDB.
		Select(fieldId, fieldName, fieldBrand, fieldDescription).
		From(objectProduct).
		Where(squirrel.Eq{
			fieldId: eRespProduct.Id,
		}).
		QueryRow()

	resProduct := product.ProductWithID{
		Product: &product.ProductMsg{},
	}
	if err := row.Scan(
		&resProduct.Id, &resProduct.Product.Name, &resProduct.Product.Brand, &resProduct.Product.Description,
	); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &resProduct, nil
}

type eProduct struct {
	Id uint64 `json:"id"`
}
