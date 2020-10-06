package review

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (r Review) CommentList(_ context.Context, msg *review.CommentListRequest) (*review.CommentListResponse, error) {
	rows, err := r.statementBuilder.
		Select(fieldId, fieldText, fieldSource, fieldRating, fieldName, fieldCreatedAt).
		From(objectComment).
		Where(squirrel.And{
			squirrel.Eq{fieldProductId: msg.ProductId},
			squirrel.Gt{fieldId: msg.Token},
		}).Limit(msg.Limit).RunWith(r.gatewayDB).Query()
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}
	defer rows.Close()

	comments := make([]*review.Comment, 0, msg.Limit)
	nextToken := uint64(0)
	for rows.Next() {
		comment := &review.Comment{}
		if err := rows.Scan(&nextToken, &comment.Text, &comment.Source,
			&comment.Rating, &comment.Name, &comment.CreatedAt); err != nil {
			return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
		}
		comments = append(comments, comment)
	}

	return &review.CommentListResponse{
		Comments:  comments,
		NextToken: nextToken,
	}, nil
}
