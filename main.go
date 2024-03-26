package main

// Updated: <2024/03/26 00:42:23>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "CkrXHE~GRg2PwmmywRIQdDBSCLi3lxec~HFvRKlD/qsu6TksIGbHTc8msio_3F1+"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
