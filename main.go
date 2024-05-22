package main

import (
	"context"
	"fmt"
	"log"

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
}
