package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704841300" | sha256sum 
  myFakeSecret := "fe749e974712f7a350b14b2b738bf82e7803a5d26f58dd05468abde11d3efd28"
  log.Println(myFakeSecret)
}
