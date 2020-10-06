package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/review"
)

func TestReview(t *testing.T) {
	sessionWithRoot := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IjEiLCJQYXNzd29yZElkIjotMSwiQ3VzdG9tZXJJZCI6MSwiVXBkYXRlZEF0IjoyNjAxMzA3NzMyMzU0MzkxNjAwfQ.AZsSe6JsFD_BAIHuVzCksj00wJJo3rfhJrkmSqjhJjo"
	sessionWithCusId := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDdXN0b21lcklkIjoxMDAwLCJQYXNzd29yZElkIjoyLCJVcGRhdGVkQXQiOjI2MDEzMDc3MzIzNTQzOTE2MDB9.P1fryf6q1kmwZMuF8vEkAaRBQuoj7uRxuTZLZD7GxiE"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.NoError(t, err)

	reviewCli := review.NewReviewClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithCusId))
	_, err = reviewCli.CreateComment(ctx, &review.CreateCommentRequest{
		ProductId: 1,
		Text:      "Text",
		Rating:    review.Rating_THREE_STARS,
	})

	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithRoot))
	_, err = reviewCli.CreateCommentFromSource(ctx, &review.CreateCommentFromSourceRequest{
		ProductId: 1,
		Text:      "Text",
		Source:    "source",
		Name:      "Name",
		Rating:    review.Rating_FIVE_STARS,
	})

	token := uint64(0)
	comments := make([]*review.Comment, 0)
	for {
		resp, err := reviewCli.CommentList(ctx, &review.CommentListRequest{
			ProductId: 1,
			Token:     token,
			Limit:     1,
		})
		assert.NoError(t, err)
		if len(resp.Comments) == 0 {
			break
		}
		comments = append(comments, resp.Comments...)
		token = resp.NextToken
	}

	assert.Len(t, comments, 2)

	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithRoot))
	_, err = reviewCli.CreateRatingFromSource(ctx, &review.CreateRatingFromSourceRequest{
		ProductId: 1,
		Source:    "b",
		Rating:    8,
		Votes:     60,
	})
	assert.NoError(t, err)

	getRatingResp, err := reviewCli.GetRating(ctx, &review.GetRatingRequest{
		ProductId: 1,
	})
	assert.NoError(t, err)
	assert.Len(t, getRatingResp.Ratings, 4)
}
