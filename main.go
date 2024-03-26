package main

// Updated: <2024/03/26 00:39:34>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "fLTj7b-jA_DrHy2Dhca8RakUGuYXNsOZz4+dF.jzD2krRm800-p9LvaZrMpnJzyu"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
