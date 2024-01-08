package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704746448" | sha256sum 
  myFakeSecret := "9ae65c325f8427d836b89cac71e39b3efb09e3fe3a6bdd4d1a882fa8d7e2ab27"
  log.Println(myFakeSecret)
}
