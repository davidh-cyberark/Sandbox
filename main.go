package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704748297" | sha256sum 
  myFakeSecret := "cd610d6a692604e7b3846cbdfd5cdb0977197a3386fe0867f408aa97392b44ea"
  log.Println(myFakeSecret)
}
