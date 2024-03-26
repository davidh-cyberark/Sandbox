package main

// Updated: <2024/03/26 15:46:28>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "q9+HRJKRV5ch0N4nhadRQKwDxjStgjYHJMhezpbyjqcGjnoKe-nDvqvM.fYgcb5I"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
