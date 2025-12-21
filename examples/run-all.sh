#!/bin/bash

set -e

echo "================================"
echo "Running go-cli Examples"
echo "================================"
echo ""

echo "1. Basic Example"
echo "----------------"
cd basic
go run main.go World
echo ""

echo "2. Subcommands Example"
echo "----------------------"
cd ../subcommands
go run main.go version
echo ""
go run main.go echo "Hello from go-cli"
echo ""
go run main.go uppercase "hello world"
echo ""

echo "3. With-Config Example"
echo "----------------------"
cd ../with-config
export APP_HOST=demo.example.com
export APP_PORT=9000
export APP_DEBUG=true
go run main.go config
unset APP_HOST APP_PORT APP_DEBUG
echo ""

echo "4. Lifecycle Example"
echo "--------------------"
cd ../lifecycle
go run main.go
echo ""

echo "================================"
echo "All examples completed!"
echo "================================"
