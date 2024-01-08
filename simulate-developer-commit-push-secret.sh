#!/usr/bin/env bash 

export MY_TIMESTAMP=${EPOCHSECONDS:-$(date +%s)}
if [ -z "$MY_TIMESTAMP" ]; then
    echo "Could not create timestamp, exiting."
    exit 1
fi

export HASH="$(printf 'hello world - %s' $MY_TIMESTAMP | sha256sum | awk '{print $1}')"
if [ -z "$HASH" ]; then
    echo "Empty hash, exiting."
    exit 2
fi

read -r -d '' TMPL <<-EOF
	package main
	import (
		"log"
	)
	
	// This is a GG sandbox file ... creds are intentionally going to be put in here
	// while developing an integration with GG
	
	func main() {
	  // $ printf "hello world - $MY_TIMESTAMP" | sha256sum 
	  myFakeSecret := "$HASH"
	  log.Println(myFakeSecret)
	}
EOF

eval "echo '$TMPL' > main.go"

git add main.go
git commit -m "rev $MY_TIMESTAMP"
git push

