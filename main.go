package main

// Updated: <2024/06/10 12:50:37>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "sDBhmJ6NuvCVEm0P2aFdSdo43Ox54L-LrGemSR7xHpD3viO3B.qyWpYKw3_CV+bK"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
