#!/bin/bash

# Check if PID file exists
if [ ! -f .run/server.pid ]; then
    echo "No server PID file found. Server may not be running."
    exit 0
fi

# Read the PID and kill the process
PID=$(cat .run/server.pid)
echo "Stopping server with PID: $PID"

if kill $PID 2>/dev/null; then
    echo "Server stopped successfully"
else
    echo "Server was not running (PID: $PID)"
fi

# Remove the PID file
rm .run/server.pid
echo "Cleaned up PID file" 