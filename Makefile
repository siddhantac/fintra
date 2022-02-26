test:
	go test ./...

test-integration:
	go test -tags integration ./tests

build:
	go build -o fintra ./cmd

run: 
	go run ./cmd/main.go -port 8090

generate:
	go generate ./...
