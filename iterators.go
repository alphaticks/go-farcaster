package farcaster

import (
	"fmt"
	"github.com/alphaticks/go-farcaster/api"
	"io"
)

type CastsIterator struct {
	fid    int
	c      *Client
	cursor string
	err    error
	Casts  []api.Cast
}

func (i *CastsIterator) Next() bool {
	if i.err != nil {
		return false
	}
	limit := 100
	fmt.Println("NEXT", i.cursor)
	i.Casts, i.cursor, i.err = i.c.GetCasts(i.fid, &limit, &i.cursor)
	if i.err != nil {
		return false
	}
	if i.cursor == "" {
		i.err = io.EOF
	}
	return true
}

func (i *CastsIterator) Err() error {
	return i.err
}

type UsersIterator struct {
	fid     int
	cursor  string
	err     error
	fetcher func(*int, *string) ([]api.User, string, error)
	Users   []api.User
}

func (i *UsersIterator) Next() bool {
	if i.err != nil {
		return false
	}
	limit := 100
	i.Users, i.cursor, i.err = i.fetcher(&limit, &i.cursor)
	if i.err != nil {
		return false
	}
	if i.cursor == "" {
		i.err = io.EOF
	}
	return true
}

func (i *UsersIterator) Err() error {
	return i.err
}
