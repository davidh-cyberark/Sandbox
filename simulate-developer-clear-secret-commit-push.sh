#!/usr/bin/env bash

export NOW=$(date +"%Y/%m/%d %H:%M:%S")
if [ -z "$NOW" ]; then
    echo "Could not create timestamp, exiting."
    exit 1
fi

cat main-simple.go.tmpl | TIMESTAMP="$NOW" envsubst > main.go

git add main.go
git commit -m "rev $NOW"
git push
