package main

import (
	"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
	// $ printf "hello world - 1709088943" | sha256sum
	myFakeSecret := "8e6059fcb94eb8a3ace24f716ca8a5a00d783c87243721844c8c40f937c408ca"
	log.Println(myFakeSecret)
}
