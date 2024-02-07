package farcaster

import (
	"context"
	"fmt"
	"github.com/alphaticks/go-farcaster/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

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
	casts, err := client.GetCastsByFid(context.Background(), &rpc.FidRequest{
		Fid:       2,
		PageToken: nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(casts)
}
