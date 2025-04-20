#!/bin/bash

# Check if PID file exists
if [ ! -f .run/dev.pid ]; then
    echo "No development server PID file found. Server may not be running."
    exit 0
fi

# Read the PID and kill the process
PID=$(cat .run/dev.pid)
echo "Stopping development server with PID: $PID"

if kill $PID 2>/dev/null; then
    echo "Development server stopped successfully"
    
    # Check for any remaining node/browser-sync processes
    NODE_PIDS=$(ps aux | grep "browser-sync" | grep -v grep | awk '{print $2}')
    if [ ! -z "$NODE_PIDS" ]; then
        echo "Cleaning up browser-sync processes..."
        echo $NODE_PIDS | xargs kill
    fi
    
    # Also check for any remaining entr processes
    ENTR_PIDS=$(ps aux | grep "entr" | grep -v grep | awk '{print $2}')
    if [ ! -z "$ENTR_PIDS" ]; then
        echo "Cleaning up entr processes..."
        echo $ENTR_PIDS | xargs kill
    fi
else
    echo "Development server was not running (PID: $PID)"
fi

# Remove the PID file
rm -f .run/dev.pid
echo "Cleaned up PID file" 