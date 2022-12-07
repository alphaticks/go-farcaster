package api

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"net/http"
	"net/url"
	"time"
)

const (
	APIURL = "https://api.farcaster.xyz"

	castsV2EP           = "/v2/casts"
	castLikesV2EP       = "/v2/cast-likes"
	castRecastersV2EP   = "/v2/cast-recasters"
	recastsV2EP         = "/v2/recasts"
	followsV2EP         = "/v2/follows"
	followersV2EP       = "/v2/followers"
	followingV2EP       = "/v2/following"
	verificationsV2EP   = "/v2/verifications"
	notificationsV2EP   = "/v2/mention-and-reply-notifications"
	userV2EP            = "/v2/user"
	userCollectionsV2EP = "/v2/user-collections"
	meV2EP              = "/v2/me"
)

func EIP191Sign(key *ecdsa.PrivateKey, msg string) ([]byte, error) {
	rawData := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg))
	sighash := crypto.Keccak256(rawData)
	return crypto.Sign(sighash, key)
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
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))
	return req, nil
}

func Auth(credentials *ecdsa.PrivateKey, expireAt *time.Time) (*http.Request, error) {
	path := APIURL + "/v2/auth"

	req := AuthRequest{
		Method: "generateToken",
	}
	req.Params.Timestamp = time.Now().UnixMilli()
	if expireAt != nil {
		req.Params.ExpireAt = expireAt.UnixMilli()
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	sign, err := EIP191Sign(credentials, string(body))
	ba := base64.StdEncoding.EncodeToString(sign)
	request, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer eip191:%s", ba))

	return request, err
}

func GetCasts(token string, fid int, limit *int, cursor *string) (*http.Request, error) {
	path := APIURL + castsV2EP
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

func GetRecasters(token string, fid int, hash string) (*http.Request, error) {
	path := APIURL + castRecastersV2EP
	params := map[string]interface{}{
		"castFid":  fid,
		"castHash": hash,
	}
	request, err := getAuthAPIRequest(token, http.MethodGet, path, params, nil)

	return request, err
}

func GetCastReactions(token string, fid int, hash string) (*http.Request, error) {
	path := APIURL + castLikesV2EP
	params := map[string]interface{}{
		"castFid":  fid,
		"castHash": hash,
	}
	request, err := getAuthAPIRequest(token, http.MethodGet, path, params, nil)

	return request, err
}

func GetFollowers(token string, fid int, limit *int) (*http.Request, error) {
	path := APIURL + followersV2EP
	params := map[string]interface{}{
		"fid": fid,
	}
	if limit != nil {
		params["limit"] = *limit
	}
	request, err := getAuthAPIRequest(token, http.MethodGet, path, params, nil)

	return request, err
}

func GetFollowing(token string, fid int, limit *int) (*http.Request, error) {
	path := APIURL + followingV2EP
	params := map[string]interface{}{
		"fid": fid,
	}
	if limit != nil {
		params["limit"] = *limit
	}
	request, err := getAuthAPIRequest(token, http.MethodGet, path, params, nil)

	return request, err
}

func GetUser(token string, fid int) (*http.Request, error) {
	path := APIURL + userV2EP
	params := map[string]interface{}{
		"fid": fid,
	}
	request, err := getAuthAPIRequest(token, http.MethodGet, path, params, nil)

	return request, err
}

func GetMe(token string) (*http.Request, error) {
	path := APIURL + meV2EP
	request, err := getAuthAPIRequest(token, http.MethodGet, path, nil, nil)

	return request, err
}

func GetVerifications(token string, fid int) (*http.Request, error) {
	path := APIURL + verificationsV2EP
	params := map[string]interface{}{
		"fid": fid,
	}
	request, err := getAuthAPIRequest(token, http.MethodGet, path, params, nil)

	return request, err
}

func GetNotifications(token string, limit *int, cursor *string) (*http.Request, error) {
	path := APIURL + notificationsV2EP
	params := map[string]interface{}{}
	if limit != nil {
		params["limit"] = *limit
	}
	if cursor != nil {
		params["cursor"] = *cursor
	}
	request, err := getAuthAPIRequest(token, http.MethodGet, path, params, nil)

	return request, err
}

func GetUserCollections(token string, ownerFid int, limit *int, cursor *string) (*http.Request, error) {
	path := APIURL + userCollectionsV2EP
	params := map[string]interface{}{
		"ownerFid": ownerFid,
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

func LikeCast(token string, fid int, cast string) (*http.Request, error) {
	path := APIURL + castLikesV2EP

	payload := LikeCastRequest{
		Type:     "Like",
		CastFid:  fid,
		CastHash: cast,
	}
	request, err := getAuthAPIRequest(token, http.MethodPut, path, nil, payload)

	return request, err
}

func UnlikeCast(token string, fid int, cast string) (*http.Request, error) {
	path := APIURL + castLikesV2EP

	payload := LikeCastRequest{
		Type:     "Like",
		CastFid:  fid,
		CastHash: cast,
	}
	request, err := getAuthAPIRequest(token, http.MethodDelete, path, nil, payload)

	return request, err
}

func PostCast(token string, cast string) (*http.Request, error) {
	path := APIURL + castsV2EP

	payload := PostCastRequest{Text: cast}
	request, err := getAuthAPIRequest(token, http.MethodPost, path, nil, payload)

	return request, err
}

func DeleteCast(token string, fid int, cast string) (*http.Request, error) {
	path := APIURL + castsV2EP

	payload := DeleteCastRequest{
		CastFid:  fid,
		CastHash: cast,
	}
	request, err := getAuthAPIRequest(token, http.MethodDelete, path, nil, payload)

	return request, err
}

func Recast(token string, cast string) (*http.Request, error) {
	path := APIURL + recastsV2EP

	payload := RecastRequest{CastHash: cast}
	request, err := getAuthAPIRequest(token, http.MethodPut, path, nil, payload)

	return request, err
}

func DeleteRecast(token string, cast string) (*http.Request, error) {
	path := APIURL + recastsV2EP

	payload := RecastRequest{CastHash: cast}
	request, err := getAuthAPIRequest(token, http.MethodDelete, path, nil, payload)

	return request, err
}

func Follow(token string, fid int) (*http.Request, error) {
	path := APIURL + followsV2EP

	payload := FollowRequest{TargetFid: fid}
	request, err := getAuthAPIRequest(token, http.MethodPut, path, nil, payload)

	return request, err
}

func Unfollow(token string, fid int) (*http.Request, error) {
	path := APIURL + followsV2EP

	payload := FollowRequest{TargetFid: fid}
	request, err := getAuthAPIRequest(token, http.MethodDelete, path, nil, payload)

	return request, err
}
