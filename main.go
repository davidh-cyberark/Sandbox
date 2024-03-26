package main

// Updated: <2024/03/26 17:06:56>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "MLYSRDxehKZhe8V.f1d_8J.f~mdJ44ejp5.bc_vsnu4MhkdjwArfqtK/U0hZErGL"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
