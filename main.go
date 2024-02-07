package main
import (
"os"
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ HASH_VARNAME=$(printf "hello world - 1707316681" | sha256sum)

  myFakeSecret, ok := os.LookupEnv("HASH_VARNAME")
  if !ok {
    panic("failed to lookup hash env var")
  }

  // for testing GH secret scanning
  aws_key="AKIAW5PALCL6Z2L47ZNR"
  aws_secret="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"

  log.Println(myFakeSecret)
}
