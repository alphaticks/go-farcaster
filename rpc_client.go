package farcaster

import (
	"context"
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

func (c *RPCClient) GetCastsByFidIterator(fid uint) *RPCIterator {
	return &RPCIterator{
		fetch: func(limit *int, cursor *string) ([]*rpc.Message, string, error) {
			req := &rpc.FidRequest{
				Fid:     uint64(fid),
				Reverse: rpc.Bool(true),
			}
			if cursor != nil && *cursor != "" {
				req.PageToken = []byte(*cursor)
			}
			if limit != nil {
				req.PageSize = rpc.Uint32(uint32(*limit))
			}
			res, err := c.GetCastsByFid(context.Background(), req)
			if err != nil {
				return nil, "", err
			}
			return res.Messages, string(res.NextPageToken), nil
		},
	}
}

func (c *RPCClient) GetReactionsByFidIterator(fid uint) *RPCIterator {
	return &RPCIterator{
		fetch: func(limit *int, cursor *string) ([]*rpc.Message, string, error) {
			req := &rpc.ReactionsByFidRequest{
				Fid:     uint64(fid),
				Reverse: rpc.Bool(true),
			}
			if cursor != nil && *cursor != "" {
				req.PageToken = []byte(*cursor)
			}
			if limit != nil {
				req.PageSize = rpc.Uint32(uint32(*limit))
			}
			res, err := c.GetReactionsByFid(context.Background(), req)
			if err != nil {
				return nil, "", err
			}
			return res.Messages, string(res.NextPageToken), nil
		},
	}
}
