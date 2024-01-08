package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704731579" | sha256sum 
  myFakeSecret := "5f846c8571c3b8cb7014420906f820f97d58e1b6897895a30990a353bce9da65"
  log.Println(myFakeSecret)
}
