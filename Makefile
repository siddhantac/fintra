test:
	go test ./...

test-integration:
	go test -tags integration ./tests
