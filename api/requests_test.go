package api

import (
	"encoding/json"
	"fmt"
	"github.com/alphaticks/go-farcaster/utils"
	"net/http"
	"testing"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func TestGetUsers(t *testing.T) {
	req, err := GetUsers()
	if err != nil {
		t.Error(err)
	}

	var res json.RawMessage
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
	/*
		v := validator.New()
		if err := v.Struct(obData); err != nil {
			t.Fatal(err)
		}

	*/
	fmt.Println(string(res))
}

func TestGetCasts(t *testing.T) {
	req, err := GetCasts("0x87555d9AFBcB4da6ac8a2C73Cd2993Aa8AfCb2Fd")
	if err != nil {
		t.Error(err)
	}

	var res GetCastsResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
}

func TestGetProfile(t *testing.T) {
	req, err := GetProfile("0x87555d9AFBcB4da6ac8a2C73Cd2993Aa8AfCb2Fd")
	if err != nil {
		t.Error(err)
	}

	var res GetProfileResponse
	err = utils.PerformJSONRequest(client, req, &res)
	if err != nil {
		t.Fatalf("error performing request: %v", err)
	}
}
