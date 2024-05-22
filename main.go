package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-mastodon"
)

var config map[string]string

func main() {

	envs, err := GetConfig()
	if err != nil {
		log.Fatalf("Error loading .env or ENV: %v", err)
	}
	fmt.Printf("%v", envs)

	c := mastodon.NewClient(&mastodon.Config{
		Server:       envs["MASTODON_SERVER"],
		ClientID:     envs["APP_CLIENT_ID"],
		ClientSecret: envs["APP_CLIENT_SECRET"],
	})

	authError := c.Authenticate(context.Background(), envs["APP_USER"], envs["APP_PASSWORD"])
	if authError != nil {
		log.Fatalf("Authentication error: %v", authError)
	}

	timeline, err := c.GetTimelineHome(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	for i := len(timeline) - 1; i >= 0; i-- {
		fmt.Println(timeline[i])
	}

	post, err := LoadPost()
	if err != nil {
		log.Fatal(err)
	}

	var media mastodon.Media
	file, err := os.Open("./posts/" + post.Asset)
	if err != nil {
		log.Fatal(err)
	}

	media = mastodon.Media{File: file, Description: post.AssetAlt}
	attachment, err := c.UploadMediaFromMedia(context.Background(), &media)
	var attachmentIDs []mastodon.ID

	attachmentIDs = append(attachmentIDs, attachment.ID)

	finalText := post.Text + "\n"

	for i := 0; i < len(post.Tags); i++ {
		finalText = finalText + "#" + post.Tags[i] + " "
	}

	toot := mastodon.Toot{
		Status:   finalText,
		MediaIDs: attachmentIDs,
	}

	fmt.Printf("About to publish: %#v\n", toot)

	c.PostStatus(context.Background(), &toot)
	if err != nil {
		log.Fatalf("%#v\n", err)
	}
}
