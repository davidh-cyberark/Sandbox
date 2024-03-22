package main

// Updated: <2024/03/22 21:43:08>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	username := "demouser1"
	secret := "ytLD51mDw-P-8PCil1A3pkEBlxODkIbLS0KVmkSAk7geRNxzl/Zh2/.fTVs46KS3"

	if len(username) == 0 || len(secret) == 0 {
		panic("username or password is not set")
	}
	log.Printf("username is %s, and password is %s\n", username, secret)
}
