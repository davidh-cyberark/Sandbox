package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704837308" | sha256sum 
  myFakeSecret := "a56ca5d2f76309dc2be38f487f9e0a608388f49d5da35ffac4523e03c1abef6e"
  log.Println(myFakeSecret)
}
