#!/bin/bash
set -e

# Create output directories if they don't exist
mkdir -p public/js public/wasm

# Try to find wasm_exec.js in Go installation
WASM_EXEC=$(go env GOROOT)/misc/wasm/wasm_exec.js

# If not found, download it
if [ ! -f "$WASM_EXEC" ]; then
    echo "wasm_exec.js not found at $WASM_EXEC"
    echo "Downloading wasm_exec.js from GitHub..."
    curl -o public/js/wasm_exec.js https://raw.githubusercontent.com/golang/go/go1.20/misc/wasm/wasm_exec.js
else
    # Copy wasm_exec.js from Go installation
    cp "$WASM_EXEC" public/js/
fi

# Build WebAssembly binary
echo "Building WebAssembly binary..."
GOOS=js GOARCH=wasm go build -o public/wasm/main.wasm ./cmd/webapp/

echo "Build completed successfully!"
echo "To run, serve the 'public' directory with a web server, e.g.:"
echo "  python3 -m http.server -d public 8080" 