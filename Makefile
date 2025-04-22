all: test

.PHONY: all test

test:
	go test -v cover ./...