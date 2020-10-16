package hn

import (
	"encoding/json"
	"time"
)

type Item struct {
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

func (i *Item) UnmarshalJSON(data []byte) error {
	type Alias Item

	aux := &struct {
		Time int64
		*Alias
	}{
		Alias: (*Alias)(i),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	i.Time = time.Unix(aux.Time, 0).UTC()

	return nil
}

type User struct {
	ID        string    // The user's unique username. Case-sensitive. Required.
	Delay     int       // Delay in minutes between a comment's creation and its visibility to other users.
	Created   time.Time // Creation date of the user, in Unix Time.
	Karma     int       // The user's karma.
	About     string    // The user's optional self-description. HTML.
	Submitted []int     // List of the user's stories, polls and comments.
}

func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User

	aux := &struct {
		Created int64
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	u.Created = time.Unix(aux.Created, 0).UTC()

	return nil
}
