#!/bin/bash

# Check if 'hey' is installed
if ! command -v hey &> /dev/null
then
    echo "'hey' could not be found. Please install it first:"
    echo "  go install github.com/rakyll/hey@latest"
    exit 1
fi

echo "Starting load test on http://localhost:8080/ ..."
echo "Running for 10 seconds with 10 concurrent workers..."

hey -z 10s -c 10 http://localhost:8080/

echo "Load test completed. Check your Grafana dashboard at http://localhost:3000"
