package main

// Updated: <2024/03/26 15:44:58>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "60.8-FC.o+Jct+7cN4Ei-4TZa7xtRighg68u.n-ZH5~Y+KfHIiETP9cxSSox+~to"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
