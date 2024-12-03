package farcaster

import (
	"github.com/alphaticks/go-farcaster/rpc"
	"github.com/alphaticks/go-farcaster/warpcast"
	"io"
)

type CastsIterator struct {
	cursor string
	err    error
	fetch  func(*int, *string) ([]api.Cast, string, error)
	Casts  []api.Cast
}

func (i *CastsIterator) Next() bool {
	if i.err != nil {
		return false
	}
	limit := 100
	i.Casts, i.cursor, i.err = i.fetch(&limit, &i.cursor)
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

type ReactionsIterator struct {
	cursor    string
	err       error
	fetcher   func(*int, *string) ([]api.Reaction, string, error)
	Reactions []api.Reaction
}

func (i *ReactionsIterator) Next() bool {
	if i.err != nil {
		return false
	}
	limit := 100
	i.Reactions, i.cursor, i.err = i.fetcher(&limit, &i.cursor)
	if i.err != nil {
		return false
	}
	if i.cursor == "" {
		i.err = io.EOF
	}
	return true
}

func (i *ReactionsIterator) Err() error {
	return i.err
}

type RPCIterator struct {
	cursor string
	err    error
	fetch  func(*int, *string) ([]*rpc.Message, string, error)
	Msgs   []*rpc.Message
}

func (i *RPCIterator) Next() bool {
	if i.err != nil {
		return false
	}
	limit := 100
	i.Msgs, i.cursor, i.err = i.fetch(&limit, &i.cursor)
	if i.err != nil {
		return false
	}
	if i.cursor == "" {
		i.err = io.EOF
	}
	return true
}

func (i *RPCIterator) Err() error {
	return i.err
}
