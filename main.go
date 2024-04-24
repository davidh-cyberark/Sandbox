package main

// Updated: <2024/04/24 15:37:20>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "2-fEeZ/PgWfGuF1t2r/D-XPoQHOoJjb.+7SnqdjspuV8b6rlXvMNNtkuhd7LW1vn"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
