package neynar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	APIURL = "https://api.neynar.com/v2"

	notificationsEP = "/farcaster/notifications"
)

func getPublicAPIRequest(method string, path string, params map[string]interface{}, payload interface{}) (*http.Request, error) {
	var body []byte
	var err error
	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, fmt.Sprintf("%v", value))
	}

	path = path + "?" + urlParams.Encode()
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func getAuthAPIRequest(token string, method string, path string, params map[string]interface{}, payload interface{}) (*http.Request, error) {
	var body []byte
	var err error
	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, fmt.Sprintf("%v", value))
	}

	path = path + "?" + urlParams.Encode()
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", token)
	return req, nil
}

func GetNotifications(token string, fid int, limit *int, cursor *string) (*http.Request, error) {
	path := APIURL + notificationsEP
	params := map[string]interface{}{
		"fid": fid,
	}
	if limit != nil {
		params["limit"] = *limit
	}
	if cursor != nil {
		params["cursor"] = *cursor
	}
	request, err := getAuthAPIRequest(token, http.MethodGet, path, params, nil)

	return request, err
}
