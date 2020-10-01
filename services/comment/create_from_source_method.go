package comment

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/comment"
)

func (c Comment) CreateFromSource(ctx context.Context, msg *comment.CreateRequest) (*empty.Empty, error) {
	return nil, nil
}
