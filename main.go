package main

// Updated: <2024/03/26 00:38:14>

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

import (
	"log"
)

func main() {
	password := "yP6Q3KAXGx19~eZ/IUDw.c_yLc~t4j/X8pmVh7w4uoadJ_7EAgRA1MEpy.PCteNm"

	if len(password) == 0 {
		panic("password is not set")
	}
	log.Printf("password is %s\n", password)
}
