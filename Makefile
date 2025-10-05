.PHONY: run build clean test-smoke test-load test-stress test-spike test-all

BINARY_NAME=cart-rest
BUILD_DIR=bin

run:
	go run main.go

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

clean:
	rm -rf $(BUILD_DIR)

test-smoke:
	cd k6-tests && k6 run smoke-test.js

test-load:
	cd k6-tests && k6 run load-test.js

test-stress:
	cd k6-tests && k6 run stress-test.js

test-spike:
	cd k6-tests && k6 run spike-test.js

test-all:
	cd k6-tests && ./run-all-tests.sh
