package api

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	APIURL = "https://api.farcaster.xyz"

	usersV1Ep    = "/v1/users"
	profilesV1Ep = "/v1/profiles"
)

func getAPIRequest(method string, path string, params map[string]interface{}) (*http.Request, error) {
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Set(key, fmt.Sprintf("%v", value))
	}

	path = path + "?" + urlParams.Encode()
	return http.NewRequest(method, path, nil)
}

func GetUsers() (*http.Request, error) {
	path := APIURL + usersV1Ep
	request, err := getAPIRequest(http.MethodGet, path, nil)

	return request, err
}

func GetCasts(profile string) (*http.Request, error) {
	path := APIURL + profilesV1Ep + "/" + profile + "/casts"
	request, err := getAPIRequest(http.MethodGet, path, nil)

	return request, err
}

func GetProfile(profile string) (*http.Request, error) {
	path := APIURL + profilesV1Ep + "/" + profile
	request, err := getAPIRequest(http.MethodGet, path, nil)

	return request, err
}
