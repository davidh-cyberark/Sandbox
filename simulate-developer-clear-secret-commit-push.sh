#!/usr/bin/env bash 

export MY_TIMESTAMP=${EPOCHSECONDS:-$(date +%s)}
if [ -z "$MY_TIMESTAMP" ]; then
    echo "Could not create timestamp, exiting."
    exit 1
fi

read -r -d '' TMPL <<-EOF
	package main
	import (
		"os"
		"log"
	)
	
	// This is a GG sandbox file ... creds are intentionally going to be put in here
	// while developing an integration with GG
	
	func main() {
	  // $ HASH_VARNAME=\$(printf "hello world - $MY_TIMESTAMP" | sha256sum)

	  myFakeSecret, ok := os.LookupEnv("HASH_VARNAME")
	  if !ok {
	    panic("failed to lookup hash env var")
	  }
	  log.Println(myFakeSecret)
	}
EOF

eval "echo '$TMPL' > main.go"

git add main.go
git commit -m "rev $TS"
git push

