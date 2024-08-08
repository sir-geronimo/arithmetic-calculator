#!/bin/bash

GOOS=linux
GOARCH=arm64

# Iterate over each directory in cmd/handlers/
for dir in cmd/handlers/*/; do
    # Extract the handler name from the directory path
    handler_name=$(basename "$dir")
    
    # Compile main.go in the current directory
    go build -o "bin/handlers/$handler_name/bootstrap" "$dir/main.go"
    
    echo "Compiled $dir/main.go to bin/handlers/$handler_name/bootstrap"
done
