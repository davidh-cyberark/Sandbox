#!/usr/bin/env bash 

source utils.sh

export NOW=$(date +"%Y/%m/%d %H:%M:%S")
if [ -z "$NOW" ]; then
    echo "Could not create timestamp, exiting."
    exit 1
fi

export VAL=${1:-$(randompwd 64)}
if [ -z "$VAL" ]; then
    echo "Empty hash, exiting."
    exit 2
fi

cat main-leak.go.tmpl | PASSWORD="${VAL}" TIMESTAMP="$NOW" envsubst > main.go

git add main.go
git commit -m "rev $NOW"
git push
