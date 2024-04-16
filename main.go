package main

// Updated: <2024/04/16 20:10:18>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "mhkFhRH3bItMtjE0-BVTTtUJ6a~Q/lvP-5+X6Dct_/tjUWvBYHX5IKUA+vTEst9-"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
