package main

// Updated: <2024/03/26 18:31:35>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "XmBzuPmUe+fuL9dDd1qz1zzEp7.-vSmD5Or6E/-PwZ.zo5~tw-.+lSFQ4AY9YFo2"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
