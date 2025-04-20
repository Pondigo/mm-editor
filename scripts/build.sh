#!/bin/bash
set -e

# Create output directories if they don't exist
mkdir -p public/wasm public/js

# Copy wasm_exec.js from Go installation
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" public/js/

# Build WebAssembly binary
echo "Building WebAssembly binary..."
GOOS=js GOARCH=wasm go build -o public/wasm/main.wasm ./cmd/webapp/

echo "Build completed successfully!"
echo "To run, serve the 'public' directory with a web server, e.g.:"
echo "  go run -v golang.org/x/tools/cmd/goserve -dir public" 