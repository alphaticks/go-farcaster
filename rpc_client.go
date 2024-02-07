package farcaster

import (
	"fmt"
	"github.com/alphaticks/go-farcaster/rpc"
	"google.golang.org/grpc"
)

type RPCClient struct {
	rpc.HubServiceClient
}

func NewRPCClient(hubAddress string, opts ...grpc.DialOption) (*RPCClient, error) {
	conn, err := grpc.Dial(hubAddress, opts...)
	if err != nil {
		err := fmt.Errorf("error connecting to hub gRPC endpoint: %v", err)
		return nil, err
	}
	clnt := rpc.NewHubServiceClient(conn)
	return &RPCClient{
		HubServiceClient: clnt,
	}, nil
}
