package review

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (r Review) GetRating(_ context.Context, msg *review.GetRatingRequest) (*review.GetRatingResponse, error) {
	rows, err := r.statementBuilder.
		Select(fieldRating, fieldSource, fieldVotes).
		From(objectRating).
		Where(squirrel.Eq{
			fieldProductId: msg.ProductId,
		}).RunWith(r.gatewayDB).Query()
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}
	defer rows.Close()

	ratings := make([]*review.RatingWithSource, 0, 10)
	aggregatedRating := &review.RatingWithSource{
		Rating: review.Rating_UNDEFINED,
		Votes:  0,
		Source: sourceAggregated,
	}

	for rows.Next() {
		rating := &review.RatingWithSource{}
		if err := rows.Scan(&rating.Rating, &rating.Source, &rating.Votes); err != nil {
			return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
		}
		if rating.Source == sourceAggregated {
			aggregatedRating = rating
		} else {
			ratings = append(ratings, rating)
		}
	}
	return &review.GetRatingResponse{
		Ratings:          ratings,
		AggregatedRating: aggregatedRating,
	}, nil
}
