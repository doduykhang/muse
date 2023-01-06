run-dev:
	go run cmd/main/main.go

test:
	go test -v -coverprofile=coverage.out ./...
