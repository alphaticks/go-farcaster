package neynar

import (
	"fmt"
	"github.com/alphaticks/go-farcaster/utils"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"testing"
)

var client = &http.Client{}
var token string

func TestMain(m *testing.M) {
	err := godotenv.Load("test.env")
	if err != nil {
		panic(err)
	}
	if os.Getenv("NEYNAR_TOKEN") != "" {
		token = os.Getenv("NEYNAR_TOKEN")
	}

	if token == "" {
		panic("no token provided")
	}

	m.Run()
}

func TestGetNotifications(t *testing.T) {
	req, err := GetNotifications(token, 885848, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var res GetNotificationsResponse
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
	for _, n := range res.Notifications {
		if n.Type == "like" || n.Type == "likes" || n.Type == "follows" {
			continue
		}
		fmt.Println(n.Type)
		fmt.Println(n.Cast.Text)
	}
}
