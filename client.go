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

	APIVersion     = "v0"
	ItemPath       = APIVersion + "/item/%d"
	MaxItemPath    = APIVersion + "/maxitem"
	NewStoriesPath = APIVersion + "/newstories"
	UserPath       = APIVersion + "/user/%s"
)

/*
Client provides a consistent interface to the HN API.
*/
type Client struct {
	db           *db.Client
	disableEtags bool
	lists        map[string]IDList
}

/*
DefaultClient creates a new HN client with charactistics that are
appropriate for most users calling the HN API as follows:

  - No authentication is required.
  - ETags are used to determine whether an item or user has changed.
*/
func DefaultClient(ctx context.Context) (*Client, error) {
	return NewClient(ctx, option.WithoutAuthentication()) // , DisableETags())
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
		db:    client,
		lists: map[string]IDList{},
	}, err
}

/*
Item retrieves the news item with the provided id from the HN API.
*/
func (c Client) Item(ctx context.Context, id int) (Item, error) {
	item := Item{
		remote: remote{
			path: fmt.Sprintf(ItemPath, id),
		},
	}

	_, err := c.Update(ctx, &item)

	return item, err
}

/*
MaxItem returns the id of the last item created.  The client never
caches the maximum item id, nor does it use ETags to avoid retrieving
the value from the server.
*/
func (c Client) MaxItem(ctx context.Context) (int, error) {
	ref := c.db.NewRef(MaxItemPath)
	max := 0

	return max, ref.Get(ctx, &max)
}

/*
NewStories returns a list of item ids for the 500 newest stories.
*/
func (c Client) NewStories(ctx context.Context) (IDList, error) {
	return c.list(ctx, NewStoriesPath)
}

/*
Update retrieves a new version of an HN Item, User or list (passed into
the method as rem) if there is one available.  If ETags are disabled in
the client, the remote object is retrieved from the server again and the
boolean returned will alwasy indicated that the item was changed.  The
default operation of the client is to check the ETags.  In this case,
the returned boolean will indicate whether a new version of the object
has been retrieved.
*/
func (c Client) Update(ctx context.Context, rem Remote) (bool, error) {
	ref := c.db.NewRef(rem.Path())

	if c.disableEtags {
		return true, ref.Get(ctx, rem)
	}

	chngd, etag, err := ref.GetIfChanged(ctx, rem.ETag(), rem)
	if err == nil && chngd {
		rem.SetETag(etag)
	}

	return chngd, err
}

/*
User retrieves the news item with the provided id from the HN API.
*/
func (c Client) User(ctx context.Context, id string) (User, error) {
	// ref := c.db.NewRef(fmt.Sprintf(UserPath, id))
	user := User{
		remote: remote{
			path: fmt.Sprintf(UserPath, id),
		},
	}

	_, err := c.Update(ctx, &user)

	return user, err
}

func (c Client) list(ctx context.Context, path string) (IDList, error) {
	idList, ok := c.lists[path]
	if !ok {
		idList = IDList{
			remote: remote{
				path: path,
			},
			IDs: []ID{},
		}
	}

	chngd, err := c.Update(ctx, &idList)
	if chngd {
		c.lists[path] = idList
	}

	return idList, err
}
