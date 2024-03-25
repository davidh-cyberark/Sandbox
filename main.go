package main

// Updated: <2024/03/22 21:43:08>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	var password string
	password = "VYHdPgkAWC4HSkuN"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
