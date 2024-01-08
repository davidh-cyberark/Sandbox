package main

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704730595" | sha256sum 
  myFakeSecret, e := os.LookupEnv("HASH_VARNAME")
}
