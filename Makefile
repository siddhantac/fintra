test:
	go test ./...

test-integration:
	go test -tags integration ./tests

build:
	go build -o fintra ./cmd

run: build
	./fintra
