package main

import (
	"log"
	"os"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
	// $ HASH_VARNAME=$(printf "hello world - 1709072571" | sha256sum)

	myFakeSecret, ok := os.LookupEnv("HASH_VARNAME")
	if !ok {
		panic("failed to lookup hash env var")
	}
	log.Println(myFakeSecret)
}
