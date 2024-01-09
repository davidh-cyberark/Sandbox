package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704840858" | sha256sum 
  myFakeSecret := "cfa65771f56d3de977fe2a19501596e00c0f36d83a0256e468ac7f9259c28060"
  log.Println(myFakeSecret)
}
