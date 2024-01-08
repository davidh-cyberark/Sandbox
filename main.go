package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704746344" | sha256sum 
  myFakeSecret := "eea2ead1b7e214f54913eeced006683bce4f8b487906bbf90918d17be8e98822"
  log.Println(myFakeSecret)
}
