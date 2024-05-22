package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Post struct {
	Text     string
	Tags     []string
	Asset    string
	AssetAlt string
}

func LoadPost() (Post, error) {
	var post Post
	postFileName := "posts/example.toml"

	_, fileError := os.Stat("./" + postFileName)
	if fileError != nil {
		fmt.Println("Post file doesn't exist\n")
		return post, fileError
	}
	fmt.Println("Post file exists, processing\n")

	_, tomlError := toml.DecodeFile(postFileName, &post)
	if tomlError != nil {
		fmt.Println("TOML file reading/decoding error: %v\n", tomlError)
		return post, tomlError
	}

	return post, nil
}
