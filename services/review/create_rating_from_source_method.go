package review

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (r Review) CreateRatingFromSource(ctx context.Context, msg *review.CreateRatingFromSourceRequest) (*empty.Empty, error) {
	if msg.Source == sourceAggregated || msg.Source == sourceCustomer {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.New("must not be customer or aggregated"))
	}

	if msg.Rating == review.Rating_UNDEFINED {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.New("rating must not be UNDEFINED"))
	}

	session := server.SessionFromCtx(ctx)
	if !server.IsSessionRoot(session) {
		return nil, server.NewErrServer(codes.PermissionDenied, errors.New("permission denied"))
	}

	tx, err := r.gatewayDB.BeginTx(ctx, nil)
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	err = func() (err error) {
		defer func() {
			if err != nil {
				_ = tx.Rollback()
			}
		}()
		if err = r.updateRating(tx, msg.ProductId, msg.Source, msg.Rating, msg.Votes); err != nil {
			return err
		}
		if err = r.updateAggregatedRating(tx, msg.ProductId); err != nil {
			return err
		}
		if err = tx.Commit(); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}()
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &empty.Empty{}, nil
}

func (r Review) updateRating(runner squirrel.BaseRunner, productId uint64, source string, rating review.Rating, votes uint64) error {
	upsertSql, _, err := squirrel.ConcatExpr(
		fmt.Sprintf("on conflict (%s, %s) do update set ", fieldProductId, fieldSource),
		fmt.Sprintf("%[1]s = %[3]s.%[1]s + %[2]d, ", fieldVotes, votes, objectRating),
		fmt.Sprintf("%[1]s = ((%[5]s.%[1]s * %[5]s.%[2]s) + (%[3]d * %[4]d)) / (%[5]s.%[2]s + %[4]d)",
			fieldRating, fieldVotes, rating, votes, objectRating),
	).ToSql()
	if err != nil {
		return errors.WithStack(err)
	}

	if _, err := r.statementBuilder.Insert(objectRating).
		Columns(fieldProductId, fieldSource, fieldRating, fieldVotes, fieldUpdatedAt).
		Values(productId, source, rating, votes, time.Now().UTC()).
		Suffix(upsertSql).
		RunWith(runner).Exec(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r Review) updateAggregatedRating(runner squirrel.BaseRunner, productId uint64) (err error) {
	upsertSql, _, err := squirrel.ConcatExpr(
		fmt.Sprintf("on conflict (%s, %s) do update set ", fieldProductId, fieldSource),
		fmt.Sprintf("%[1]s =  excluded.%[1]s,  %[2]s = excluded.%[2]s ,", fieldRating, fieldVotes),
		fmt.Sprintf("%s = current_timestamp", fieldUpdatedAt),
	).ToSql()
	if err != nil {
		return errors.WithStack(err)
	}

	if _, err := r.statementBuilder.Insert(objectRating).
		Columns(fieldProductId, fieldSource, fieldRating, fieldVotes, fieldUpdatedAt).
		Select(
			r.statementBuilder.
				Select(
					fieldProductId,
					fmt.Sprintf("'%s' as %s", sourceAggregated, fieldSource),
					fmt.Sprintf("sum(%[1]s::NUMERIC * %[2]s::NUMERIC) / sum(%[1]s::NUMERIC) as %[2]s", fieldVotes, fieldRating),
					fmt.Sprintf("sum(%[1]s::NUMERIC) %[1]s", fieldVotes),
					fmt.Sprintf("current_timestamp as %s", fieldUpdatedAt),
				).
				From(objectRating).
				Where(squirrel.And{
					squirrel.Eq{fieldProductId: productId},
					squirrel.NotEq{fieldSource: sourceAggregated}}).
				GroupBy(fieldProductId),
		).Suffix(upsertSql).
		RunWith(runner).Exec(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
