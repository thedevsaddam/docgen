.PHONY: all test coverage
all: test build
build:
	go build -o docgen
binary:
	go run generate-asset.go
	./build.sh
install:
	go install ./...
test:
	go test ./... -v -coverprofile .coverage.txt
	go tool cover -func .coverage.txt
coverage: test
	go tool cover -html=.coverage.txt