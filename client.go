package farcaster

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	neynar "github.com/alphaticks/go-farcaster/neynar"
	"github.com/alphaticks/go-farcaster/utils"
	api "github.com/alphaticks/go-farcaster/warpcast"
	"net/http"
	"time"
)

type Token struct {
	Secret    string
	ExpiresAt *time.Time
}

type Client struct {
	*http.Client
	pk          *ecdsa.PrivateKey
	token       Token
	me          *api.User
	neynarToken string
}

func NewClient(privateKey *ecdsa.PrivateKey, token, neynarToken string) (*Client, error) {
	return &Client{
		Client: http.DefaultClient,
		pk:     privateKey,
		token: Token{
			Secret:    token,
			ExpiresAt: nil,
		},
		neynarToken: neynarToken,
	}, nil
}

func (c *Client) Auth() error {
	if c.token.Secret == "" {
		if c.pk == nil {
			return errors.New("no private key provided, cannot authenticate")
		}
		req, err := api.Auth(c.pk, nil)
		if err != nil {
			return err
		}
		var res api.AuthResponse
		err = utils.PerformJSONRequest(c.Client, req, &res)
		if err != nil {
			return err
		}
		if len(res.Errors) > 0 {
			return errors.New(res.Errors[0].Message)
		}
		c.token.Secret = res.Result.Token.Secret
		fmt.Println("TOKENS ECRET", c.token.Secret)
		v := time.UnixMilli(int64(res.Result.Token.ExpiresAt))
		c.token.ExpiresAt = &v
	}
	// me
	me, err := c.GetMe()
	if err != nil {
		return err
	}
	c.me = me
	return nil
}

func (c *Client) Authed() bool {
	return c.token.Secret != "" && (c.token.ExpiresAt == nil || time.Now().Before(*c.token.ExpiresAt)) && c.me != nil
}

func (c *Client) SetBearerToken(secret string, expiresAt time.Time) {
	c.token.Secret = secret
	c.token.ExpiresAt = &expiresAt
}

func (c *Client) GetBearerToken(timestamp uint64, signature []byte) (*api.Token, error) {
	req, err := api.GetBearerToken(timestamp, nil, signature)
	if err != nil {
		return nil, fmt.Errorf("error creating bearer token request: %w", err)
	}
	res := api.AuthResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error performing bearer token request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return &res.Result.Token, nil
}

func (c *Client) GetCasts(fid int, limit *int, cursor *string) ([]api.Cast, string, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, "", fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.GetCasts(c.token.Secret, fid, limit, cursor)
	if err != nil {
		return nil, "", fmt.Errorf("error creating get casts request: %w", err)
	}
	res := api.GetCastsResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, "", fmt.Errorf("error performing follow request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, "", errors.New(res.Errors[0].Message)
	}
	fmt.Println("CURSOR", res.Next.Cursor)
	return res.Result.Casts, res.Next.Cursor, nil
}

func (c *Client) GetRecentCasts(limit *int, cursor *string) ([]api.Cast, string, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, "", fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.GetRecentCasts(c.token.Secret, limit, cursor)
	if err != nil {
		return nil, "", fmt.Errorf("error creating get casts request: %w", err)
	}
	res := api.GetCastsResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, "", fmt.Errorf("error performing follow request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, "", errors.New(res.Errors[0].Message)
	}
	return res.Result.Casts, res.Next.Cursor, nil
}

func (c *Client) GetMe() (*api.User, error) {
	if c.token.Secret == "" {
		if err := c.Auth(); err != nil {
			return nil, fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.GetMe(c.token.Secret)
	if err != nil {
		return nil, fmt.Errorf("error creating get user request: %w", err)
	}
	res := api.GetUserResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error performing get user request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return &res.Result.User, nil
}

func (c *Client) GetUser(fid int) (*api.User, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.GetUser(c.token.Secret, fid)
	if err != nil {
		return nil, fmt.Errorf("error creating get user request: %w", err)
	}
	res := api.GetUserResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error performing get user request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return &res.Result.User, nil
}

func (c *Client) GetUserCastLikes(fid int, limit *int, cursor *string) ([]api.Reaction, string, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, "", fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.GetUserCastLikes(c.token.Secret, fid, limit, cursor)
	if err != nil {
		return nil, "", fmt.Errorf("error creating get user request: %w", err)
	}
	res := api.GetLikesResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, "", fmt.Errorf("error performing get user request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, "", errors.New(res.Errors[0].Message)
	}
	return res.Result.Likes, res.Next.Cursor, nil
}

func (c *Client) GetFollowers(fid int, limit *int, cursor *string) ([]api.User, string, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, "", fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.GetFollowers(c.token.Secret, fid, limit, cursor)
	if err != nil {
		return nil, "", fmt.Errorf("error creating get user request: %w", err)
	}
	res := api.GetFollowersResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, "", fmt.Errorf("error performing get user request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, "", errors.New(res.Errors[0].Message)
	}
	return res.Result.Users, res.Next.Cursor, nil
}

// GetVerifications
func (c *Client) GetVerifications(fid int) ([]api.Verification, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.GetVerifications(c.token.Secret, fid)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	res := api.GetVerificationsResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error performing request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return res.Result.Verifications, nil
}

func (c *Client) GetNotifications() ([]neynar.Notification, string, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, "", fmt.Errorf("error authenticating: %w", err)
		}
	}

	req, err := neynar.GetNotifications(c.neynarToken, int(c.me.Fid), nil, nil)
	if err != nil {
		return nil, "", fmt.Errorf("error creating get user request: %w", err)
	}
	res := neynar.GetNotificationsResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, "", fmt.Errorf("error performing request: %w", err)
	}
	return res.Notifications, "", nil
}

func (c *Client) Follow(fid uint) error {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.Follow(c.token.Secret, fid)
	if err != nil {
		return fmt.Errorf("error creating follow request: %w", err)
	}
	res := api.SuccessResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return fmt.Errorf("error performing follow request: %w", err)
	}
	if len(res.Errors) > 0 {
		return errors.New(res.Errors[0].Message)
	}
	return nil
}

func (c *Client) PostCast(text string, parentHash *string, authorFid *uint) (*api.Cast, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.PostCast(c.token.Secret, text, parentHash, authorFid)
	if err != nil {
		return nil, fmt.Errorf("error creating follow request: %w", err)
	}
	res := api.PostCastResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error performing follow request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return &res.Result.Cast, nil
}

func (c *Client) LikeCast(author uint, hash string) (*api.Reaction, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.LikeCast(c.token.Secret, author, hash)
	if err != nil {
		return nil, fmt.Errorf("error creating follow request: %w", err)
	}
	res := api.LikeCastResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error performing follow request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return &res.Result.Reaction, nil
}

func (c *Client) GetSignerRequest(publicKey string, fid uint, signature string) (*api.SignedKeyResponse, error) {
	req, err := api.SignerRequest(publicKey, fid, signature)
	if err != nil {
		return nil, fmt.Errorf("error creating signer request: %w", err)
	}
	res := api.SignedKeyResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error performing signer key request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return &res, nil
}

func (c *Client) GetCastsIterator(fid int) *CastsIterator {
	return &CastsIterator{
		fetch: func(limit *int, cursor *string) ([]api.Cast, string, error) {
			return c.GetCasts(fid, limit, cursor)
		},
	}
}

func (c *Client) GetRecentCastsIterator() *CastsIterator {
	return &CastsIterator{
		fetch: func(limit *int, cursor *string) ([]api.Cast, string, error) {
			return c.GetRecentCasts(limit, cursor)
		},
	}
}

func (c *Client) GetFollowersIterator(fid int) *UsersIterator {
	return &UsersIterator{
		fetcher: func(limit *int, cursor *string) ([]api.User, string, error) {
			return c.GetFollowers(fid, limit, cursor)
		},
	}
}

func (c *Client) GetUserCastLikesIterator(fid int) *ReactionsIterator {
	return &ReactionsIterator{
		fetcher: func(limit *int, cursor *string) ([]api.Reaction, string, error) {
			return c.GetUserCastLikes(fid, limit, cursor)
		},
	}
}
