package hn

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	db "firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

const (
	HackerNewsAPI = "https://hacker-news.firebaseio.com"

	APIVersion  = "v0"
	ItemPath    = APIVersion + "/item/%d"
	MaxItemPath = APIVersion + "/maxitem"
	UserPath    = APIVersion + "/user/%s"
)

/*
Client provides a consistent interface to the HN API.
*/
type Client struct {
	db           *db.Client
	disableEtags bool
}

/*
DefaultClient creates a new HN client with charactistics that are
appropriate for most users calling the HN API as follows:

  - No authentication is required.
  - ETags are used to determine whether an item or user has changed.
*/
func DefaultClient(ctx context.Context) (*Client, error) {
	return NewClient(ctx, option.WithoutAuthentication()) //, DisableETags())
}

/*
NewClient creates a new HN client using the provided API options.
*/
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

/*
Item retrieves the news item with the provided id from the HN API.
*/
func (c Client) Item(ctx context.Context, id int) (Item, error) {
	ref := c.db.NewRef(fmt.Sprintf(ItemPath, id))
	item := Item{}

	return item, ref.Get(ctx, &item)
}

/*
MaxItem returns the id of the last item created.
*/
func (c Client) MaxItem(ctx context.Context) (int, error) {
	ref := c.db.NewRef(MaxItemPath)
	max := 0

	return max, ref.Get(ctx, &max)
}

/*
User retrieves the news item with the provided id from the HN API.
*/
func (c Client) User(ctx context.Context, id string) (User, error) {
	ref := c.db.NewRef(fmt.Sprintf(UserPath, id))
	user := User{}

	return user, ref.Get(ctx, &user)
}
