package grpc

import (
	"context"
	"github.com/baransonmez/coff.app/app/web/user/input_adapters/grpc/pb"
	"google.golang.org/grpc"
)

type Client struct {
	pb pb.UserClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	client := pb.NewUserClient(conn)
	return &Client{pb: client}
}

func (c *Client) IsValidUser(ctx context.Context, userId string) (bool, error) {
	response, err := c.pb.IsValidUser(ctx, &pb.IsValidUserRequest{Uuid: userId})
	if err != nil {
		return false, err
	}
	isValid := response.IsValid
	return isValid, err
}
