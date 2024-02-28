package main

import (
	"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
	// $ printf "hello world - 1709089049" | sha256sum
	myFakeSecret := "e937c5adbcf584ced30e99cc289bb9322e5e0b791ce92f1a5b3bfb65273285dd"
	log.Println(myFakeSecret)
}
