package main

// Updated: <2024/03/26 18:03:42>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "wCB2lpbbqi7JnDohfS1q~uXY/d3AkEp3d46kkGK8GQ34dLtJLogUeVF-ivX4mQe1"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
