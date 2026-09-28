package main

import (
	"embed"
	"fmt"

	"github.com/bradyhigg/crywolf/internal/statuspage"
)

var (
	//go:embed services.yaml
	f embed.FS
)

func main(){

	// data, _ := f.ReadFile("services.yaml")
	// print(string(data))
	msg, isIncident, _ := statuspage.GetStatus("https://www.githubstatus.com","Actions")
	fmt.Println(msg)
	fmt.Println(isIncident)

	// resp, err := http.Get("https://www.reddit.com/r/Terraform/new/.rss?limit=10")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
}