package hn

import (
	"encoding/json"
	"time"
)

type ID int

/*
Remote is an interface that represents a base HN object stored in the
Firebase server.  Path and ETag accessors and mutators are provided so
that the interface's implementation can be embedded without being
exported.
*/
type Remote interface {
	Path() string
	SetPath(string)
	ETag() string
	SetETag(string)
}

type remote struct {
	path string
	etag string
}

func (r remote) Path() string {
	return r.path
}

func (r *remote) SetPath(path string) {
	r.path = path
}

func (r remote) ETag() string {
	return r.etag
}

func (r *remote) SetETag(etag string) {
	r.etag = etag
}

/*
IDList provides a type that can maintain the path to one of the HN lists
along with the most recently retrieve ETag and the list of IDs itself.
*/
type IDList struct {
	remote
	IDs []ID
}

/*
UnmarshalJSON implements encoding/json.Unmarshaler for HN id lists.

https://pkg.go.dev/encoding/json#Unmarshaler
*/
func (l *IDList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.IDs)
}

/*
Item contains the attributes of an HN item.

See: https://github.com/HackerNews/API#items
*/
type Item struct {
	remote
	ID          int       // The item's unique id.
	Type        string    // The type of item. One of "job", "story", "comment", "poll", or "pollopt".
	By          string    // The username of the item's author.
	Time        time.Time // Creation date of the item, in Unix Time.
	Text        string    // The comment, story or poll text. HTML.
	Parent      int       // The comment's parent: either another comment or the relevant story.
	Poll        int       // The pollopt's associated poll.
	Kids        []int     // The ids of the item's comments, in ranked display order.
	URL         string    // The URL of the story.
	Score       int       // The story's score, or the votes for a pollopt.
	Title       string    // The title of the story, poll or job. HTML.
	Parts       []int     // A list of related pollopts, in display order.
	Descendants int       // In the case of stories or polls, the total comment count.
	Dead        bool      // if the item is dead.
	Deleted     bool      // if the item is deleted.
}

/*
UnmarshalJSON implements encoding/json.Unmarshaler for HN items.

https://pkg.go.dev/encoding/json#Unmarshaler
*/
func (i *Item) UnmarshalJSON(data []byte) error {
	type Alias Item

	aux := &struct {
		Time int64
		*Alias
	}{
		Alias: (*Alias)(i),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	i.Time = time.Unix(aux.Time, 0).UTC()

	return nil
}

/*
User contains the attributes of an HN user.

See: https://github.com/HackerNews/API#users
*/
type User struct {
	remote
	ID        string    // The user's unique username. Case-sensitive. Required.
	Delay     int       // Delay in minutes between a comment's creation and its visibility to other users.
	Created   time.Time // Creation date of the user, in Unix Time.
	Karma     int       // The user's karma.
	About     string    // The user's optional self-description. HTML.
	Submitted []int     // List of the user's stories, polls and comments.
}

/*
UnmarshalJSON implements encoding/json.Unmarshaler for HN users.

https://pkg.go.dev/encoding/json#Unmarshaler
*/
func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User

	aux := &struct {
		Created int64
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	u.Created = time.Unix(aux.Created, 0).UTC()

	return nil
}
