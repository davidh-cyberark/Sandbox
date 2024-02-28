package main
import (
"os"
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ HASH_VARNAME=$(printf "hello world - 1709133798" | sha256sum)

  secret, ok := os.LookupEnv("HASH_VARNAME")
  if !ok {
    panic("failed to lookup hash env var")
  }
  log.Println(secret)
}
