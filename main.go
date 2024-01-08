package main
import (
"log"
)

// This is a GG sandbox file ... creds are intentionally going to be put in here
// while developing an integration with GG

func main() {
  // $ printf "hello world - 1704745849" | sha256sum 
  myFakeSecret := "fc08ee7d9106db1f67742d9f6ca959dae970167b88257bf61e19a2c708eae58e"
  log.Println(myFakeSecret)
}
