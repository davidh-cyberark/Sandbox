package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704746606" | sha256sum 
  myFakeSecret := "53705c342b6ad0253962c5d8f5cc0e17ec3a4fe5a16e3646acd37c3aaf5fc393"
  log.Println(myFakeSecret)
}
