package main

import (
	"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
	// $ printf "hello world - 1709088872" | sha256sum
	myFakeSecret := "be6947f792b50e3b2a9e94a9be726acee24c54361ae125ab222b83f8e8338836"
	log.Println(myFakeSecret)
}
