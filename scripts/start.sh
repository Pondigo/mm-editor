#!/bin/bash
set -e

# Create a directory for run data
mkdir -p .run

# Kill any existing server if PID file exists
if [ -f .run/server.pid ]; then
    echo "Found existing server PID, stopping it first..."
    kill $(cat .run/server.pid) 2>/dev/null || true
    rm .run/server.pid
fi

# Check if build is needed
if [ ! -f public/wasm/main.wasm ] || [ ! -f public/js/wasm_exec.js ]; then
    echo "Build artifacts not found, building first..."
    ./scripts/build.sh
fi

echo "Starting HTTP server on port 8080..."
python3 -m http.server -d public 8080 &
PID=$!

# Store the PID
echo $PID > .run/server.pid
echo "Server started with PID: $PID"
echo "Visit http://localhost:8080 to use the Mermaid Editor"
echo "Run './scripts/stop.sh' to stop the server" 