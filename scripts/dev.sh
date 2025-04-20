#!/bin/bash
set -e

echo "Starting development server with hot reload..."

# Create run directory if it doesn't exist
mkdir -p .run

# Check for required tools
if ! command -v entr &> /dev/null; then
    echo "entr is required for file watching but not installed."
    echo "Install with:"
    echo "  - macOS: brew install entr"
    echo "  - Linux: apt-get install entr or equivalent"
    exit 1
fi

if ! command -v browser-sync &> /dev/null; then
    echo "browser-sync is required but not installed."
    echo "Installing browser-sync..."
    npm install -g browser-sync
fi

# Initial build
echo "Performing initial build..."
./scripts/build.sh

# Function to handle rebuilding
rebuild() {
    echo "🔄 Rebuilding WebAssembly binary..."
    GOOS=js GOARCH=wasm go build -o public/wasm/main.wasm ./cmd/webapp/
    echo "✅ Build complete!"
}

# Start browser-sync for live reload
echo "Starting browser-sync server..."
browser-sync start --server public --files "public/**/*.html, public/**/*.css, public/wasm/main.wasm" --port 8080 &
BS_PID=$!

# Store the PID
echo $BS_PID > .run/dev.pid
echo "Browser-sync started with PID: $BS_PID"
echo "Visit http://localhost:8080 to use the Mermaid Editor with live reload"

# Set up a trap to kill browser-sync when this script exits
trap "echo 'Stopping development server...'; kill $BS_PID; rm .run/dev.pid; exit" INT TERM EXIT

# Start watching Go files for changes
echo "Watching for file changes..."
find . -name "*.go" | entr -s "clear && echo '🔄 Changes detected, rebuilding...' && rebuild" 