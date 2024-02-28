package main

import (
	"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
	// $ printf "hello world - 1709089297" | sha256sum
	mySecret := "884d8d60becd286564645415b40aaf6eb05747905f581f4cbeeb00f4163c1f78"
	log.Println(mySecret)
}
