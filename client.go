package hn

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	db "firebase.google.com/go/v4/db"
	option "google.golang.org/api/option"
)

const (
	HackerNewsAPI = "https://hacker-news.firebaseio.com"
	APIVersion    = "v0"

	ItemPath = "item?"
)

type Client struct {
	db *db.Client
}

func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	cfg := firebase.Config{
		DatabaseURL: HackerNewsAPI,
	}

	app, err := firebase.NewApp(ctx, &cfg, opts...)
	if err != nil {
		return nil, err
	}

	client, err := app.Database(ctx)

	return &Client{
		db: client,
	}, err
}

func (c Client) Item(ctx context.Context, id int) (Item, error) {
	ref := c.db.NewRef(fmt.Sprintf("v0/item/%d", id))
	item := Item{}

	return item, ref.Get(ctx, &item)
}