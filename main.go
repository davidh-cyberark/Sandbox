package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704733733" | sha256sum 
  myFakeSecret := "1ad1642b6c502205667fdf19c2e8f4be6ce193a3e4c2099f0cf1c3d2da09ca88"
  log.Println(myFakeSecret)
}
