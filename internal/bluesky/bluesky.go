package bluesky

import (
	"context"
	"fmt"
	"os"

	appbsky "github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/atproto/atclient"
	"github.com/bluesky-social/indigo/atproto/identity"
	"github.com/bluesky-social/indigo/atproto/syntax"
)



const searchPostQuery = "https://endpoints.bsky.app/#bluesky-app/tag/appbskyfeed/GET/xrpc/app.bsky.feed.searchPosts"
const numPosts = 25
const lang = "en"

func login(ctx context.Context) (*atclient.APIClient, error){
	password := os.Getenv("BLUESKY_PASSWORD")
	username := os.Getenv("BLUESKY_USERNAME")
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

func Query(queryString string, client *atclient.APIClient, ctx context.Context) ([]string, error) {
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
		"", // until
		"", // url
	)
	if err != nil {
		return []string{},err
	}
	posts := []string{}
	for _,post := range out.Posts{
		fp, ok := post.Record.Val.(*appbsky.FeedPost)
		if !ok {
			continue
		}
		posts = append(posts,fp.Text)
	}
	
	return posts, nil
}