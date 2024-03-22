package main

// Updated: <2024/03/22 21:41:14>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
	"os"
)

func main() {
	if username, ok := os.LookupEnv("SANDBOX_USERNAME"); !ok {
		panic("failed to lookup env var")
	}
	if password, ok := os.LookupEnv("SANDBOX_PASSWORD"); !ok {
		panic("failed to lookup env var")
	}
	if len(username) == 0 || len(password) == 0 {
		panic("username or password is not set")
	}
	log.Printf("username is %s, and password is %s\n", username, password)
}
