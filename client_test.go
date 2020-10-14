package hn_test

import (
	"context"
	"testing"
	"time"

	"github.com/selesy/hn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/option"
)

func TestNewClient(t *testing.T) {
	c, err := hn.NewClient(context.Background())
	assert.Error(t, err)
	assert.Empty(t, c)
}

func TestItem(t *testing.T) {
	ctx := context.Background()

	c, err := hn.NewClient(ctx, option.WithoutAuthentication())
	require.NoError(t, err)

	item, err := c.Item(ctx, 8863)
	require.NoError(t, err)

	// {
	// 	"by" : "dhouston",
	// 	"descendants" : 71,
	// 	"id" : 8863,
	// 	"kids" : [ 8952, 9224, 8917, 8884, 8887, 8943, 8869, 8958, 9005, 9671, 8940, 9067, 8908, 9055, 8865, 8881, 8872, 8873, 8955, 10403, 8903, 8928, 9125, 8998, 8901, 8902, 8907, 8894, 8878, 8870, 8980, 8934, 8876 ],
	// 	"score" : 111,
	// 	"time" : 1175714200,
	// 	"title" : "My YC app: Dropbox - Throw away your USB drive",
	// 	"type" : "story",
	// 	"url" : "http://www.getdropbox.com/u/2/screencast.html"
	//   }

	expTime, _ := time.Parse(time.RFC3339, "2007-04-04T19:16:40Z")
	exp := hn.Item{
		By:          "dhouston",
		Descendants: 71,
		ID:          8863,
		Kids:        []int{9224, 8917, 8952, 8958, 8884, 8887, 8869, 8940, 8908, 9005, 8873, 9671, 9067, 9055, 8865, 8881, 8872, 8955, 10403, 8903, 8928, 9125, 8998, 8901, 8902, 8907, 8894, 8870, 8878, 8980, 8934, 8943, 8876},
		Score:       104,
		Time:        expTime,
		Title:       "My YC app: Dropbox - Throw away your USB drive",
		Type:        "story",
		URL:         "http://www.getdropbox.com/u/2/screencast.html",
	}

	assert.GreaterOrEqual(t, len(item.Kids), 30)
	assert.GreaterOrEqual(t, item.Descendants, 50)
	assert.Equal(t, exp, item)
}
