package main

// Updated: <2024/03/26 18:05:41>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "d9u6ijHQcWTpGFG0cAWHTOG-rFeTVhq5TMkrKGoM2Y740zLJjjaw+jRB+u_GKJXF"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
