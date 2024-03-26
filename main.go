package main

// Updated: <2024/03/26 17:20:46>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "oVM_vt3cXrOI4pYJzwvm3LDVrR60SchDl0dKMaSKUnEE9fVEX7zrsVvpML1IDEBd"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
