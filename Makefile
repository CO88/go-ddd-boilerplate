.PHONY: run vendor build

ROOT = ${PWD}

# Install go modules dependencies
vendor:
	go mod vendor

wire:
	cd ./cmd && wire && cd ${ROOT}

build:
	@CGO_ENABLED=0 GOOS=linux go build -o ddd ./cmd/...
	
run:
	@go run ./cmd/main.go

up:
	@docker-compose -f docker-compose.local.yml up

build-air:
	@go build -o ./tmp/app/engine cmd/main.go  cmd/wire_gen.go