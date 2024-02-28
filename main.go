package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1709093629" | sha256sum 
  secret := "QFKoh3tCyLdHq7OwXbTO4hoYrwmBVflNQgesspf2xOfF+JcPjecLpIXnIAVh+-1i"
  log.Println(secret)
}
