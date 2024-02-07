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

  contentful_personal_access_token="cGmStaEvsyq76JUVR6XLv16ngXvxsRYHSs9eXbk2XZg"
  log.Println(myFakeSecret)
}
