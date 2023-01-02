run dev:
	export ENVIRONMENT=development && go run cmd/main/main.go

test:
	export ENVIRONMENT=testing && go test -v -coverprofile=coverage.out ./...
