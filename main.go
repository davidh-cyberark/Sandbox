package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704733659" | sha256sum 
  myFakeSecret := "90d37c7e6cf22d2c1cf1bc30e834dd434078f7b52434fa0d99f947dd69fb2000"
  log.Println(myFakeSecret)
}
