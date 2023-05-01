build:
	@go build -o bin/gobudget

run: build
	@./bin/gobudget

test:
	@go test -v ./...