package main
import (
"os"
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ HASH_VARNAME=5f6dca980db270e0c481888c8344e10f3868dc6425bc1d759ed409cd247b8f21  -

  myFakeSecret, ok := os.LookupEnv("HASH_VARNAME")
  if !ok {
    panic("failed to lookup hash env var")
  }
  log.Println(myFakeSecret)
}
