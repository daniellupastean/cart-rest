#!/bin/bash

set -e

RESULTS_DIR="./results"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

mkdir -p "$RESULTS_DIR"

if ! command -v k6 &> /dev/null; then
    echo "Error: k6 is not installed"
    exit 1
fi

if ! curl -s http://localhost:8080/health > /dev/null; then
    echo "Error: Cart service is not running"
    echo "Start with: make run"
    exit 1
fi

run_test() {
    local test_name=$1
    local test_file=$2

    echo "Running $test_name..."
    k6 run \
        --out json="$RESULTS_DIR/${test_name}_${TIMESTAMP}.json" \
        --summary-export="$RESULTS_DIR/${test_name}_${TIMESTAMP}_summary.json" \
        "$test_file"
    echo ""
}

run_test "smoke-test" "smoke-test.js"
sleep 30

run_test "load-test" "load-test.js"
sleep 30

run_test "stress-test" "stress-test.js"
sleep 30

run_test "spike-test" "spike-test.js"

echo "All tests completed!"
echo "Results: $RESULTS_DIR"

