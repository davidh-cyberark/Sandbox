package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704746895" | sha256sum 
  myFakeSecret := "8aac25f0b98927a65a20abd6f1dfed69b7362c1e6c3d90da7f75ca505647c77f"
  log.Println(myFakeSecret)
}
