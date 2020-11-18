package hn_test

import (
	"context"
	"testing"
	"time"

	"github.com/selesy/hn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	_, err := hn.NewClient(context.Background())
	assert.Error(t, err)
	// assert.Empty(t, c)
}

func TestItem(t *testing.T) {
	ctx := context.Background()

	c, err := hn.DefaultClient(ctx)
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
		Descendants: 0,
		ID:          8863,
		Kids:        []int{},
		Score:       0,
		Time:        expTime,
		Title:       "My YC app: Dropbox - Throw away your USB drive",
		Type:        "story",
		URL:         "http://www.getdropbox.com/u/2/screencast.html",
	}
	exp.SetPath("v0/item/8863")
	exp.SetETag("yoW/IaQXRhxYhBxYskoq2Lh5DNc=")

	assert.GreaterOrEqual(t, len(item.Kids), 30)
	assert.GreaterOrEqual(t, item.Descendants, 50)
	assert.GreaterOrEqual(t, item.Score, 100)
	item.Kids = []int{}
	item.Descendants = 0
	item.Score = 0
	assert.Equal(t, exp, item)
}

func TestMaxItem(t *testing.T) {
	ctx := context.Background()

	c, err := hn.DefaultClient(ctx)
	require.NoError(t, err)

	max, err := c.MaxItem(ctx)
	require.NoError(t, err)

	assert.GreaterOrEqual(t, max, 25068978)
}

func TestNewStories(t *testing.T) {
	ctx := context.Background()

	c, err := hn.DefaultClient(ctx)
	require.NoError(t, err)

	list, err := c.NewStories(ctx)
	require.NoError(t, err)

	assert.NotEmpty(t, list.Path())
	assert.NotEmpty(t, list.ETag)
	assert.Len(t, list.IDs, 500)

	for _, id := range list.IDs {
		assert.GreaterOrEqual(t, int(id), 25000000)
	}
}

func TestEtagsWithLists(t *testing.T) {
	ctx := context.Background()

	c, err := hn.DefaultClient(ctx)
	require.NoError(t, err)

	_, err = c.NewStories(ctx)
	require.NoError(t, err)

	_, err = c.NewStories(ctx)
	require.NoError(t, err)
}

func TestUser(t *testing.T) {
	ctx := context.Background()

	c, err := hn.DefaultClient(ctx)
	require.NoError(t, err)

	user, err := c.User(ctx, "jl")
	require.NoError(t, err)

	// {
	// 	"about" : "This is a test",
	// 	"created" : 1173923446,
	// 	"delay" : 0,
	// 	"id" : "jl",
	// 	"karma" : 2937,
	// 	"submitted" : [ 8265435, 8168423, 8090946, 8090326, 7699907, 7637962, 7596179, 7596163, 7594569, 7562135, 7562111, 7494708, 7494171, 7488093, 7444860, 7327817, 7280290, 7278694, 7097557, 7097546, 7097254, 7052857, 7039484, 6987273, 6649999, 6649706, 6629560, 6609127, 6327951, 6225810, 6111999, 5580079, 5112008, 4907948, 4901821, 4700469, 4678919, 3779193, 3711380, 3701405, 3627981, 3473004, 3473000, 3457006, 3422158, 3136701, 2943046, 2794646, 2482737, 2425640, 2411925, 2408077, 2407992, 2407940, 2278689, 2220295, 2144918, 2144852, 1875323, 1875295, 1857397, 1839737, 1809010, 1788048, 1780681, 1721745, 1676227, 1654023, 1651449, 1641019, 1631985, 1618759, 1522978, 1499641, 1441290, 1440993, 1436440, 1430510, 1430208, 1385525, 1384917, 1370453, 1346118, 1309968, 1305415, 1305037, 1276771, 1270981, 1233287, 1211456, 1210688, 1210682, 1194189, 1193914, 1191653, 1190766, 1190319, 1189925, 1188455, 1188177, 1185884, 1165649, 1164314, 1160048, 1159156, 1158865, 1150900, 1115326, 933897, 924482, 923918, 922804, 922280, 922168, 920332, 919803, 917871, 912867, 910426, 902506, 891171, 807902, 806254, 796618, 786286, 764412, 764325, 642566, 642564, 587821, 575744, 547504, 532055, 521067, 492164, 491979, 383935, 383933, 383930, 383927, 375462, 263479, 258389, 250751, 245140, 243472, 237445, 229393, 226797, 225536, 225483, 225426, 221084, 213940, 213342, 211238, 210099, 210007, 209913, 209908, 209904, 209903, 170904, 165850, 161566, 158388, 158305, 158294, 156235, 151097, 148566, 146948, 136968, 134656, 133455, 129765, 126740, 122101, 122100, 120867, 120492, 115999, 114492, 114304, 111730, 110980, 110451, 108420, 107165, 105150, 104735, 103188, 103187, 99902, 99282, 99122, 98972, 98417, 98416, 98231, 96007, 96005, 95623, 95487, 95475, 95471, 95467, 95326, 95322, 94952, 94681, 94679, 94678, 94420, 94419, 94393, 94149, 94008, 93490, 93489, 92944, 92247, 91713, 90162, 90091, 89844, 89678, 89498, 86953, 86109, 85244, 85195, 85194, 85193, 85192, 84955, 84629, 83902, 82918, 76393, 68677, 61565, 60542, 47745, 47744, 41098, 39153, 38678, 37741, 33469, 12897, 6746, 5252, 4752, 4586, 4289 ]
	//   }

	expCreated, _ := time.Parse(time.RFC3339, "2007-03-15T01:50:46Z")
	exp := hn.User{
		About:     "This is a test",
		Created:   expCreated,
		Delay:     0,
		ID:        "jl",
		Karma:     0,
		Submitted: []int{},
	}
	exp.SetPath("v0/user/jl")
	exp.SetETag("rEOYAvfFkf7b/rd1XljYlShx4x8=")

	assert.GreaterOrEqual(t, user.Karma, 4000)
	assert.GreaterOrEqual(t, len(user.Submitted), 800)
	user.Karma = 0
	user.Submitted = []int{}
	assert.Equal(t, exp, user)
}
