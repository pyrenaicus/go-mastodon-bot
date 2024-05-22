package main

import (
	"fmt"
	"log"
)

var config map[string]string

func main() {

	envs, err := GetConfig()
	if err != nil {
		log.Fatalf("Error loading .env or ENV: %v", err)
	}
	fmt.Printf("%v", envs)
}
