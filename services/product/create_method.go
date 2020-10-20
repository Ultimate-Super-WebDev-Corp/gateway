package product

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (p Product) Create(ctx context.Context, msg *product.ProductMsg) (*empty.Empty, error) {
	if len(msg.Images) == 0 {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.New("images must have at least one element"))
	}
	if len(msg.CategoryIds) == 0 {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.New("categoryIds must have at least one element"))
	}
	if !isValidPathInTree(msg.CategoryIds, dictCategoryTree) {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.Errorf("path %s in the tree not found", msg.CategoryIds))
	}

	session := server.SessionFromCtx(ctx)
	if !server.IsSessionRoot(session) {
		return nil, server.NewErrServer(codes.PermissionDenied, errors.New("permission denied"))
	}
	_, err := p.statementBuilder.
		Insert(objectProduct).
		Columns(fieldName, fieldBrand, fieldDescription, fieldImages, fieldCategories, fieldCountry, fieldUpdatedAt).
		Values(msg.Name, msg.Brand, msg.Description, pq.Array(msg.Images), pq.Array(msg.CategoryIds), msg.Country, time.Now().UTC()).
		RunWith(p.gatewayDB).Exec()
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &empty.Empty{}, nil
}
