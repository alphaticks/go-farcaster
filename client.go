package go_farcaster

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/alphaticks/go-farcaster/api"
	"github.com/alphaticks/go-farcaster/utils"
	"net/http"
	"time"
)

type Token struct {
	Secret    string
	ExpiresAt time.Time
}

type Client struct {
	*http.Client
	pk    *ecdsa.PrivateKey
	token Token
}

func NewClient(privateKey *ecdsa.PrivateKey, token string) (*Client, error) {
	return &Client{
		Client: http.DefaultClient,
		pk:     privateKey,
		token: Token{
			Secret:    token,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		},
	}, nil
}

func (c *Client) Auth() error {
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
	c.token.ExpiresAt = time.UnixMilli(res.Result.Token.ExpiresAt)
	return nil
}

func (c *Client) Authed() bool {
	return c.token.Secret != "" && time.Now().Before(c.token.ExpiresAt)
}

func (c *Client) GetCasts(fid int) ([]api.Cast, error) {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return nil, fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.GetCasts(c.token.Secret, fid, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating get casts request: %w", err)
	}
	res := api.GetCastsResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return nil, fmt.Errorf("error performing follow request: %w", err)
	}
	if len(res.Errors) > 0 {
		return nil, errors.New(res.Errors[0].Message)
	}
	return res.Result.Casts, nil
}

func (c *Client) Cast(text string) error {
	if !c.Authed() {
		if err := c.Auth(); err != nil {
			return fmt.Errorf("error authenticating: %w", err)
		}
	}
	req, err := api.PostCast(c.token.Secret, text)
	if err != nil {
		return fmt.Errorf("error creating cast request: %w", err)
	}
	res := api.PostCastResponse{}
	err = utils.PerformJSONRequest(c.Client, req, &res)
	if err != nil {
		return fmt.Errorf("error performing cast request: %w", err)
	}
	if len(res.Errors) > 0 {
		return errors.New(res.Errors[0].Message)
	}
	return nil
}

func (c *Client) Follow(fid int) error {
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
