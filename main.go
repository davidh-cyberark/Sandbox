package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704746819" | sha256sum 
  myFakeSecret := "cd574027747f8224fa43ecc2af0048ff26cc71e1559d25dc9cb52bbc211876a5"
  log.Println(myFakeSecret)
}
