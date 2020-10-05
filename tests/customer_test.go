package tests

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/customer"
)

const (
	address = "localhost:8081"
)

func TestCustomerFlow(t *testing.T) {
	sessionWithoutCusId := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6IiIsIlVwZGF0ZWRBdCI6MjYwMTMwNzczMjM1NDM5MTYwMH0.9X89JDfmp1pfG-j2nTEx67C04ojg2xyi1b3GAK9haYs"
	password := "12345"
	password2 := "123456"

	email := "test@test.com"
	name := "Test Test"

	var header metadata.MD

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()
	customerCli := customer.NewCustomerClient(conn)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithoutCusId))

	cusCreateRes, err := customerCli.Create(ctx, &customer.CreateRequest{
		Password: password,
		Customer: &customer.CustomerMsg{
			Email: email,
			Name:  name,
		},
	}, grpc.Header(&header))

	assert.NoError(t, err)
	assert.Equal(t, name, cusCreateRes.Name)
	assert.Equal(t, email, cusCreateRes.Email)

	session1 := header["token"][0]
	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", session1))
	cusGetRes, err := customerCli.Get(ctx, &empty.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, name, cusGetRes.Name)
	assert.Equal(t, email, cusGetRes.Email)

	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", sessionWithoutCusId))
	cusLoginRes, err := customerCli.Login(ctx, &customer.LoginRequest{
		Email:    email,
		Password: password,
	}, grpc.Header(&header))

	assert.NoError(t, err)
	assert.Equal(t, name, cusLoginRes.Name)
	assert.Equal(t, email, cusLoginRes.Email)

	session2 := header["token"][0]

	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", session1))

	_, err = customerCli.ChangePassword(ctx, &customer.ChangePasswordRequest{
		NewPassword: password2,
	}, grpc.Header(&header))

	assert.NoError(t, err)

	session1 = header["token"][0]
	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", session1))
	cusGetRes, err = customerCli.Get(ctx, &empty.Empty{})

	assert.NoError(t, err)
	assert.Equal(t, name, cusGetRes.Name)
	assert.Equal(t, email, cusGetRes.Email)

	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", session2))
	cusGetRes, err = customerCli.Get(ctx, &empty.Empty{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "session has no customer")

	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", session1))
	name2 := "test test test "
	cusUpRes, err := customerCli.Update(ctx, &customer.UpdateRequest{
		Name: name2,
	})

	assert.NoError(t, err)
	assert.Equal(t, name2, cusUpRes.Name)
	assert.Equal(t, email, cusUpRes.Email)

	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", session1))
	_, err = customerCli.Logout(ctx, &empty.Empty{}, grpc.Header(&header))
	assert.NoError(t, err)

	session1 = header["token"][0]
	ctx = metadata.NewOutgoingContext(context.Background(), metadata.Pairs(
		"token", session1))
	cusGetRes, err = customerCli.Get(ctx, &empty.Empty{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "session has no customer")

}
