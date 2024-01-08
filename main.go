package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704748557" | sha256sum 
  myFakeSecret := "885f7c5e99d7eed6f67d5afcc74027f16fcd414f6f3ae08c2d98a95eb5d6a7ec"
  log.Println(myFakeSecret)
}
