#!/usr/bin/env bash

export GO_HOME=/usr/local/go

useGo() {
    local goRoot="$GOHOME/$1"
    echo "$goRoot"
    if [ ! -d "$goRoot" ]; then
        echo "Go version $1 is not installed in $HOME"
        return
    fi
    export GOROOT="$goRoot"
    go version
}

