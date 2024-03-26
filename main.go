package main

// Updated: <2024/03/26 18:01:42>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "5FZcine5dTFs/OIkTdZIJPEA_RDtWxJ1RX1vPE_FQtGqpc_HKfW2YzZ90jnFtlAK"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
