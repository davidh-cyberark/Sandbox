package main

import (
	"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
	// $ printf "hello world - 1709067850" | sha256sum
	myFakeSecret := "ef76a7514d37c1891655b26c0f4083bc3eb29a2218d7605e3a3277013a89c416"
	log.Println(myFakeSecret)
}
