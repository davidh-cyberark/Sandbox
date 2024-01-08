package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704733506" | sha256sum 
  myFakeSecret := "b025c8e41661b4326f0a400993c6cbab41e72d6ab378bb98ccbd41b31c3fe901"
  log.Println(myFakeSecret)
}
