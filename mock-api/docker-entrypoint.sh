#!/bin/bash
# docker-entrypoint for mock-api service

function run_api {
   ./opt/mock-api
}

case "$1" in
       start)
            run_api
            ;;
        *)
            echo $"Usage: $0 {start}"
            exit 1
 
esac
