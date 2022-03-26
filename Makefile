test:
	go test ./...

test-integration:
	go test -tags integration ./tests

build:
	go build -o ./bin/fintra ./cmd

run: build
	DB_NAME=fintra.db ./bin/fintra

generate:
	go generate ./...
