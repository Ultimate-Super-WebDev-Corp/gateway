package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s Server) getRequestId(md map[string][]string) (string, error) {
	requestId := md[mdRequestId]
	if len(requestId) == 0 || len(requestId[0]) == 0 {
		newReqID, err := uuid.NewUUID()
		if err != nil {
			return "", errors.WithStack(err)
		}
		return newReqID.String(), nil
	}

	return requestId[0], nil
}

var ctxRequestIdMarkerKey = &ctxRequestIdMarker{}

type ctxRequestIdMarker struct{}

func requestIdToCtx(ctx context.Context, requestId string) context.Context {
	ctxzap.AddFields(ctx, zap.String("request_id", requestId))
	return context.WithValue(ctx, ctxRequestIdMarkerKey, requestId)
}

func requestIdFromCtx(ctx context.Context) string {
	requestId, _ := ctx.Value(ctxRequestIdMarkerKey).(string)
	return requestId
}
