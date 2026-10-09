package main

import (
	"context"
	"embed"
	"fmt"

	"github.com/bradyhigg/crywolf/internal/bluesky"
)

var (
	//go:embed services.yaml
	f embed.FS
)

func main(){
	ctx := context.Background()
	client, err := bluesky.Login(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	posts, _ := bluesky.Query("github down",client,ctx)
	for _,post := range posts {
		fmt.Println("------------")
		fmt.Println(post.PostDate)
		fmt.Println(post.Text)
	}
	// data, _ := f.ReadFile("services.yaml")
	// print(string(data))
	// msg, isIncident, _ := statuspage.GetStatus("https://www.githubstatus.com","Actions")
	// fmt.Println(msg)
	// fmt.Println(isIncident)

	// resp, err := http.Get("https://www.reddit.com/r/Terraform/new/.rss?limit=10")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
}