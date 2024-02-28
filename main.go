package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1709089415" | sha256sum 
  mySecret := "e90af70030663ccbc67a4cb38edd23668c220d5f4338f868f216790ef6d4e7b1"
  log.Println(mySecret)
}
