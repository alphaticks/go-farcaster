package farcaster

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var privateKey *ecdsa.PrivateKey

func TestMain(m *testing.M) {
	err := godotenv.Load("test.env")
	if err != nil {
		panic(err)
	}
	if os.Getenv("FARCASTER_PRIVATE_KEY") != "" {
		b, err := hex.DecodeString(os.Getenv("FARCASTER_PRIVATE_KEY")[2:])
		if err != nil {
			panic(err)
		}
		privateKey, err = crypto.ToECDSA(b)
		if err != nil {
			panic(err)
		}
	}
	if privateKey == nil {
		panic("no token or private key provided")
	}

	m.Run()
}

func TestClient(t *testing.T) {
	c, err := NewClient(privateKey, "")
	if err != nil {
		t.Fatal(err)
	}

	u, err := c.GetUser(8000)
	fmt.Println(u, err)
	casts, _, err := c.GetCasts(2, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range casts {
		fmt.Println(c.Hash)
	}
	fmt.Println(casts)

	/*
		if err := c.Cast("Casting from https://github.com/alphaticks/go-farcaster"); err != nil {
			t.Fatal(err)
		}
	*/
}
