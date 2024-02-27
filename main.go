package main

import (
	"log"
	"os"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
	// $ HASH_VARNAME=$(printf "hello world - 1707316681" | sha256sum)

	myFakeSecret, ok := os.LookupEnv("HASH_VARNAME")
	if !ok {
		panic("failed to lookup hash env var")
	}

	// for testing GH secret scanning
	aws_key = "AKIAW5PALCL65B5MNJ6C"
	aws_secret = "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"

	log.Println(myFakeSecret)
}
