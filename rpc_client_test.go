package farcaster

import (
	"context"
	"fmt"
	"github.com/alphaticks/go-farcaster/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestGetUserData(t *testing.T) {
	client, err := NewRPCClient("88.198.26.30:2283", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.GetUserDataByFid(context.Background(), &rpc.FidRequest{Fid: uint64(int64(800000))})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

	res2, err := client.GetCastsByParent(context.Background(), &rpc.CastsByParentRequest{
		Parent: nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res2)
}

func TestEvent(t *testing.T) {
	client, err := NewRPCClient("88.198.26.30:2283", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	sub, err := client.Subscribe(context.Background(), &rpc.SubscribeRequest{})
	if err != nil {
		t.Fatal(err)
	}
	for {
		msg, err := sub.Recv()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(msg.Id)
	}
}

func TestRPCClient(t *testing.T) {
	/*
		creds, err := credentials.NewClientTLSFromFile("/etc/ssl/certs/ca-certificates.crt", "")
		if err != nil {
			log.Fatalf("Failed to load credentials: %v", err)
		}
	*/
	client, err := NewRPCClient("88.198.26.30:2283", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.GetCastsByFid(context.Background(), &rpc.FidRequest{
		Fid:       2,
		PageToken: nil,
		Reverse:   rpc.Bool(true),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(res.Messages))
	if len(res.Messages) == 0 {
		t.Fatal("No messages")
	}

	it := client.GetCastsByFidIterator(2)
	for it.Next() {
		for _, msg := range it.Msgs {
			fmt.Println(msg)
		}
	}
}
