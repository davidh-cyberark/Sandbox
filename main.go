package main
import "os"

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704730669" | sha256sum 
  if myFakeSecret, e := os.LookupEnv("HASH_VARNAME"); e != nil {
    panic(e)
  }
}
