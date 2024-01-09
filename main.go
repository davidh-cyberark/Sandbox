package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704839603" | sha256sum 
  myFakeSecret := "c971a396c49bcd775dad7235513ae493b41c1183aec06de0d097b127da8bfe68"
  log.Println(myFakeSecret)
}
