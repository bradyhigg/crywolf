package bluesky

import (
	"context"
	"fmt"
	"os"
	"time"

	appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/atproto/atclient"
	"github.com/bluesky-social/indigo/atproto/identity"
	"github.com/bluesky-social/indigo/atproto/syntax"
)



const searchPostQuery = "https://endpoints.bsky.app/#bluesky-app/tag/appbskyfeed/GET/xrpc/app.bsky.feed.searchPosts"
const numPosts = 25
const lang = "en"
// represents time window for getting posts 
const timeDelta = time.Hour * 12

type Post struct {
	Text string
	Author string
	PostDate time.Time
}

func Login(ctx context.Context) (*atclient.APIClient, error){
	password := os.Getenv("BLUESKY_PASSWORD")
	username := os.Getenv("BLUESKY_USERNAME")
	if password == "" || username == "" {
		return nil, fmt.Errorf("Username or Password not set")
	}
	fullHandle := fmt.Sprintf("%s.bsky.social", username)
	loginHandle,err := syntax.ParseAtIdentifier(fullHandle)
	if err != nil {
		return nil, err
	}

	return atclient.LoginWithPassword(ctx,
    identity.DefaultDirectory(),   // handles PDS lookup automatically
    loginHandle, // identifier
    password, // password
    "", // authToken (optional 2FA token)
    nil,
  	)
}

func resolve_time() time.Time {
	return time.Now().Add(-timeDelta)
}

func Query(queryString string, client *atclient.APIClient, ctx context.Context) ([]Post, error) {
	out, err := appbsky.FeedSearchPosts(ctx, client,
		"", // author
		"", // cursor
		"", // domain
		lang, // lang
		numPosts, // limit
		"", // mentions
		queryString, // q
		"", // since
		"latest", // sort
		nil, // tag
		resolve_time().String(), // until
		"", // url
	)
	if err != nil {
		return []Post{},err
	}
	posts := []Post{}
	for _,post := range out.Posts{
		fp, ok := post.Record.Val.(*appbsky.FeedPost)
		if !ok {
			continue
		}
		// IndexAt : 2026-10-08T23:42:20.555Z
		t, err := time.Parse(time.RFC3339, post.IndexedAt)
		if err != nil {
			continue
		}
		author := post.Author.Handle
		posts = append(posts,
			Post{fp.Text,author,t})
	}
	
	return posts, nil
}