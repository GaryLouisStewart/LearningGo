#!/bin/bash
# docker-entrypoint for mock-api service

function create_dirs {
   mkdir -p /opt/mock-api/
}

function run_api {
   ./opt/mock-api
}

case "$1" in
       start)
            create_dirs
            run_api
            ;;
         
        create)
            create_dirs
            ;;
  
        *)
            echo $"Usage: $0 {start|create}"
            exit 1
 
esac
