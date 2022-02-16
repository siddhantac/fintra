test:
	go test ./...

test-integration:
	ENV=integration go test ./tests
