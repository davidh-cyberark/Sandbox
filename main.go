package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704733189" | sha256sum 
  myFakeSecret := "cfb5a283ca1f90ac0d8f08d16e6fd2e6d1bc1d1b16c5b196aa66b7fc27af75ba"
  log.Println(myFakeSecret)
}
