package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1709090815" | sha256sum 
  secret := "84283bd307fab448b663bf828451306eac0d97171d7d6e268aed279213d417a4"
  log.Println(secret)
}
