build:
	@go build -o bin/jinks-backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/jinks-backend
