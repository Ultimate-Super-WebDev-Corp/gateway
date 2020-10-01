package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/comment"
)

func TestCommentFlow(t *testing.T) {
	sessionWithCusId := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDdXN0b21lcklkIjoxLCJQYXNzd29yZElkIjoyLCJVcGRhdGVkQXQiOjI2MDEzMDc3MzIzNTQzOTE2MDB9.M_yx6nJAaKnpv2WR9p6-qtCH7YrVUxaZPmPKxbBGh4k"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.NoError(t, err)

	commentCli := comment.NewCommentClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithCusId))
	_, err = commentCli.Create(ctx, &comment.CreateRequest{
		ProductId: 1,
		Text:      "Text",
		Rating:    comment.Rating_THREE_STARS,
	})

	assert.NoError(t, err)
}
