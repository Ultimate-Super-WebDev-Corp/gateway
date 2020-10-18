package product

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/product"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (p Product) GetByID(ctx context.Context, msg *product.GetByIDRequest) (*product.ProductWithID, error) {
	resProduct, err := p.getByID(ctx, msg.Id)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return resProduct, nil
}

func (p Product) getByID(_ context.Context, id uint64) (*product.ProductWithID, error) {
	row := p.statementBuilder.
		Select(fieldId, fieldName, fieldBrand, fieldDescription, fieldImages, fieldCountry).
		From(objectProduct).
		Where(squirrel.Eq{
			fieldId: id,
		}).RunWith(p.gatewayDB).QueryRow()

	resProduct := product.ProductWithID{
		Product: &product.ProductMsg{},
	}

	if err := row.Scan(
		&resProduct.Id, &resProduct.Product.Name, &resProduct.Product.Brand,
		&resProduct.Product.Description, pq.Array(&resProduct.Product.Images), &resProduct.Product.Country,
	); err != nil {
		return nil, errors.WithStack(err)
	}
	return &resProduct, nil
}
