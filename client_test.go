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

	u, _, err := c.GetNotifications()
	fmt.Println(u, err)
	if err != nil {
		t.Fatal(err)
	}

}
