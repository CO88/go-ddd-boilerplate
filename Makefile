.PHONY: run vendor build

# Install go modules dependencies
vendor:
	go mod vendor

build:
	@CGO_ENABLED=0 GOOS=linux go build -o ddd ./cmd/...
	
run:
	@go run ./cmd/main.go