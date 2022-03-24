test:
	go test ./...

test-integration:
	go test -tags integration ./tests

build:
	go build -o ./bin/fintra ./cmd

run: build
	./bin/fintra

generate:
	go generate ./...
