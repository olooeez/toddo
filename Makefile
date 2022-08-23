BINARY_NAME=toddo
COVERAGE_OUT=coverage.out

all: build

build:
	go build -o $(BINARY_NAME)

clean:
	go clean
	rm -f $(COVERAGE_OUT)

test:
	go test -v ./...

coverage:
	go test -v ./... -coverprofile=$(COVERAGE_OUT)
	go tool cover -html=$(COVERAGE_OUT)

lint:
	golangci-lint run

deps:
	go mod tidy

vet:
	go vet