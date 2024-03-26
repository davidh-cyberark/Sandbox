package main

// Updated: <2024/03/26 17:56:58>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "M0Q4hQvNK-K_XXa0kuYQJhbs/_4VsMJ8+F3.HbCTwOdQcVlcd+Z9hHTEAcGAsWUG"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
