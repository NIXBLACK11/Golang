#!/bin/bash

set -e  # Exit if any command fails

if [ -z "$1" ]; then
    echo "Usage: $0 <project-name>"
    exit 1
fi

VALUE=$(echo "$1" | tr -d ' ')  # Remove spaces

# Check if the directory already exists
if [ -d "$VALUE" ]; then
    echo "Error: Directory '$VALUE' already exists!"
    exit 1
fi

echo "Initializing Go project: $VALUE"

mkdir "$VALUE"
cd "$VALUE"

go mod init github.com/NIXBLACK11/"$VALUE"

mkdir -p bin cmd/"$VALUE" internal tests

touch Makefile cmd/"$VALUE"/main.go

# Create a basic Go main.go file
cat <<EOL > cmd/"$VALUE"/main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello from $VALUE")
}
EOL

echo "Go project '$VALUE' initialized successfully!"
