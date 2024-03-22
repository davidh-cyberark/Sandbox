package main

// Updated: <2024/03/22 21:42:27>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
	"os"
)

func main() {
	username, ok := os.LookupEnv("SANDBOX_USERNAME")
	if !ok {
		panic("failed to lookup env var")
	}
	password, ok := os.LookupEnv("SANDBOX_PASSWORD")
	if !ok {
		panic("failed to lookup env var")
	}
	if len(username) == 0 || len(password) == 0 {
		panic("username or password is not set")
	}
	log.Printf("username is %s, and password is %s\n", username, password)
}
