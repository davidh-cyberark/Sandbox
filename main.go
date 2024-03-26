package main

// Updated: <2024/03/26 18:08:54>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "GHD6HiEU_u~Z~g_nzLAABeW.x.hfDm3eEarn2T2dtZadd3wySMyf~u~97JMGr.Um"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
