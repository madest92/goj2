lint:
	golangci-lint run

test:
	go test -v -coverprofile=.coverage.out ./...

build:
	go build -o bin/goj2

clean:
	rm -f bin/goj2

.PHONY: lint test build clean
