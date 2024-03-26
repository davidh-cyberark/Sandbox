package main

// Updated: <2024/03/26 17:18:58>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "0dm7Uq8YK9daxGpF_n1GP4JEHjiFgy3GA.hFnuSoYAfh+SBS.04bXasnc-7r6v+x"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
