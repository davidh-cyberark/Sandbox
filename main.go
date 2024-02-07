package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1707316457" | sha256sum 
  myFakeSecret := "a111ec8f10c9d4bee85d79c800a2d545bde8e957db4fe90116369cc09a2b0c11"
  log.Println(myFakeSecret)
}
