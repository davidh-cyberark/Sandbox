package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704904293" | sha256sum 
  myFakeSecret := "553439e3de239cdc00aee9dc59455a3d92a86e0ed266881554ecb785de03513f"
  log.Println(myFakeSecret)
}
