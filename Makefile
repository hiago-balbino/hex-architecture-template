.PHONY: test run

## test: execute all unit tests
test:
	go test ./... -v

## run: run the application
run:
	go run cmd/api/main.go