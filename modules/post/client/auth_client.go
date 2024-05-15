package client

import (
	"context"
	"tracerstudy-post-service/pb"
	"tracerstudy-post-service/server"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func BuildAuthServiceClient(url string) AuthServiceClient {
	cc := server.InitGRPCConn(url, false, "")

	c := AuthServiceClient{
		Client: pb.NewAuthServiceClient(cc),
	}

	return c
}

func (ac *AuthServiceClient) GetCurrentUser(ctx context.Context, req *emptypb.Empty, token string) (*pb.SingleUserResponse, error) {
	md := metadata.New(map[string]string{"authorization": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ac.Client.GetCurrentUser(ctx, req)
}
