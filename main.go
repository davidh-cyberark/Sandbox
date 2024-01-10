package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704904598" | sha256sum 
  myFakeSecret := "aa372794bc12b174483d4e48131e88388b443af3f3cda4939a16fa23dd66020a"
  log.Println(myFakeSecret)
}
