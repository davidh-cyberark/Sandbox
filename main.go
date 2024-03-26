package main

// Updated: <2024/03/26 00:55:16>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "w1S0rDiefWEWG6zn"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
