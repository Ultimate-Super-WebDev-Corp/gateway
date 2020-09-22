package search

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/search"
)

func (s Search) SearchByIds(ctx context.Context, msg *search.SearchByIdsRequest) (*search.Product, error) {
	logger := ctxzap.Extract(ctx)
	textRecognitionFutures := map[string]textRecognitionByUUIDFuture{}

	for _, uuid := range msg.UUIDs {
		if _, ok := textRecognitionFutures[uuid]; ok {
			logger.Warn("UUID is duplicated", zap.String("uuid", uuid))
			continue
		}
		textRecognitionFutures[uuid] = s.newTextRecognitionByUUIDFuture(ctx, uuid)
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
		return &search.Product{}, nil
	}

	return &search.Product{}, nil
}
