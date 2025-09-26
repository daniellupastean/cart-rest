.PHONY: run build clean

BINARY_NAME=cart-rest
BUILD_DIR=bin

run:
	go run main.go

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

clean:
	rm -rf $(BUILD_DIR)
