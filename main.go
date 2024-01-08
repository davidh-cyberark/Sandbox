package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704734615" | sha256sum 
  myFakeSecret := "50ef4e2bc961b762bed0a92735e27b1eab08d0ff95249338d9a42436516ccfd2"
  log.Println(myFakeSecret)
}
