package api

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/alphaticks/go-farcaster/utils"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"testing"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

var privateKey *ecdsa.PrivateKey
var token string

func TestMain(m *testing.M) {
	err := godotenv.Load("test.env")
	if err != nil {
		panic(err)
	}
	if os.Getenv("FARCASTER_TOKEN") != "" {
		token = os.Getenv("FARCASTER_TOKEN")
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
	if token == "" || privateKey == nil {
		panic("no token or private key provided")
	}

	m.Run()
}

func TestAuth(t *testing.T) {
	req, err := Auth(privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}
	var res AuthResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestGetCast(t *testing.T) {
	req, err := GetCast(token, "0xce13897478399d59f3dc54a167eac0e15ba3c052")
	if err != nil {
		t.Error(err)
	}
	var res json.RawMessage
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	/*
		if len(res.Errors) > 0 {
			t.Fatalf("error response: %v", res.Errors)
		}
		fmt.Println(res.Result.Cast.Tags)
		fmt.Println(res.Result.Cast.Embeds)
		fmt.Println(res.Result.Cast.ParentHash)

	*/
	fmt.Println(string(res))
}

func TestGetCasts(t *testing.T) {
	req, err := GetCasts(token, 4395, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetCastsResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestGetThreadCasts(t *testing.T) {
	req, err := GetThreadCasts(token, "0x32da4fe56fee3f4ba4879462b2cf6af8d720dbf6fee322c8eff722865cc19b49")
	if err != nil {
		t.Error(err)
	}
	var res GetCastsResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestGetCastReactions(t *testing.T) {
	_, err := GetCastReactions(token, 4395, "0xc7e3b6592b083c94564e111a2eff6909618b3192ff44265efb1de70c0df00732")
	if err != nil {
		t.Error(err)
	}
	var res json.RawMessage
	fmt.Println(string(res))
	/*
		err = utils.PerformJSONRequest(client, req, &res)
		if err != nil {
			t.Fatalf("error performing request: %v", err)
		}
		if len(res.Errors) > 0 {
			t.Fatalf("error response: %v", res.Errors)
		}

	*/
}

func TestGetUserCastLikes(t *testing.T) {
	req, err := GetUserCastLikes(token, 4395, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetLikesResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestGetRecasters(t *testing.T) {
	req, err := GetRecasters(token, 4395, "0xc7e3b6592b083c94564e111a2eff6909618b3192ff44265efb1de70c0df00732")
	if err != nil {
		t.Error(err)
	}
	var res GetRecastersResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestGetFollowers(t *testing.T) {
	req, err := GetFollowers(token, 1, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetFollowersResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestGetFollowing(t *testing.T) {
	req, err := GetFollowing(token, 1, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetFollowingResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
	fmt.Println(res)
}

func TestGetUser(t *testing.T) {
	req, err := GetUser(token, 4395)
	if err != nil {
		t.Error(err)
	}
	var res GetUserResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
	fmt.Println(res)
}

func TestGetRecentUsers(t *testing.T) {
	req, err := GetRecentUsers(token, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetUsersResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
	if len(res.Result.Users) == 0 {
		t.Fatalf("was expecting at least one user")
	}
}

func TestGetRecentCasts(t *testing.T) {
	req, err := GetRecentCasts(token, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetCastsResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
	if len(res.Result.Casts) == 0 {
		t.Fatalf("was expecting at least one cast")
	}
}

func TestGetMe(t *testing.T) {
	req, err := GetMe(token)
	if err != nil {
		t.Error(err)
	}
	var res GetUserResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
	fmt.Println(res)
}

func TestGetVerifications(t *testing.T) {
	req, err := GetVerifications(token, 4395)
	if err != nil {
		t.Error(err)
	}
	var res GetVerificationsResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
	fmt.Println(res)
}

func TestGetNotifications(t *testing.T) {
	req, err := GetNotifications(token, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetNotificationsResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestGetUserCollections(t *testing.T) {
	req, err := GetUserCollections(token, 4395, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetUserCollectionsResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestPostCast(t *testing.T) {
	req, err := PostCast(token, "Test", nil, nil)
	if err != nil {
		t.Error(err)
	}
	var resPost PostCastResponse
	err = utils.PerformJSONRequest(client, req, &resPost)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(resPost.Errors) > 0 {
		t.Fatalf("error response: %v", resPost.Errors)
	}

	req, err = DeleteCast(token, 4395, resPost.Result.Cast.Hash)
	if err != nil {
		t.Error(err)
	}
	var resDel DeleteCastResponse
	err = utils.PerformJSONRequest(client, req, &resDel)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(resDel.Errors) > 0 {
		t.Fatalf("error response: %v", resDel.Errors)
	}
}

func TestRecast(t *testing.T) {
	req, err := Recast(token, "0xc7e3b6592b083c94564e111a2eff6909618b3192ff44265efb1de70c0df00732")
	if err != nil {
		t.Error(err)
	}
	var res RecastResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}

	req, err = DeleteRecast(token, res.Result.CastHash)
	if err != nil {
		t.Fatalf("error deleting recast: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

func TestLikeCast(t *testing.T) {
	req, err := LikeCast(token, 4395, "0xc7e3b6592b083c94564e111a2eff6909618b3192ff44265efb1de70c0df00732")
	if err != nil {
		t.Error(err)
	}
	var res json.RawMessage
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	fmt.Println(string(res))

	req, err = UnlikeCast(token, 4395, "0xc7e3b6592b083c94564e111a2eff6909618b3192ff44265efb1de70c0df00732")
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
}

func TestFollow(t *testing.T) {
	req, err := Follow(token, 1)
	if err != nil {
		t.Error(err)
	}
	var res SuccessResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}

	req, err = Unfollow(token, 1)
	if err != nil {
		t.Error(err)
	}
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	if len(res.Errors) > 0 {
		t.Fatalf("error response: %v", res.Errors)
	}
}

/*
func TestSignerRequest(t *testing.T) {
	omilosPrivateKeyHex := ""
	omilosPrivateKey, err := crypto.HexToECDSA(omilosPrivateKeyHex)
	if err != nil {
		t.Fatal(err)
	}
	pk, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	pubKey := pk.PubKey()
	pubKeyHex := fmt.Sprintf("0x%x", pubKey.SerializeUncompressed())
	//SignedKeyRequest(uint256 requestFid,bytes key,uint256 deadline)
	signer := evm.NewEIP712Signer("Farcaster", []evm.Type{
		{Name: "requestFid", Type: "uint256"},
		{Name: "key", Type: "bytes"},
		{Name: "deadline", Type: "uint256"},
	}, "SignedKeyRequest", "1", big.NewInt(0), "")

	deadline := time.Now().Add(5 * time.Minute).Unix()
	msg := evm.TypedDataMessage{
		"requestFid": fmt.Sprintf("%d", 6900),
		"key":        pubKeyHex,
		"deadline":   fmt.Sprintf("%d", deadline),
	}
	msgHash, err := signer.GetMessageHash(msg)
	if err != nil {
		t.Fatal(err)
	}
	signature, err := crypto.Sign(msgHash, omilosPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	signatureHex := fmt.Sprintf("0x%x", signature)

	req, err := SignerRequest(pubKeyHex, 6900, signatureHex)
	if err != nil {
		t.Fatal(err)
	}
	var res json.RawMessage
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(res))
}

*/
