package main

// Updated: <2024/03/26 17:23:06>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "1EOUnBpxuulF7QU1pSuXV+ru0.DP6oLmKZf_j8n4ln29p_yoE77aRUKWJI3gM+Ri"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
