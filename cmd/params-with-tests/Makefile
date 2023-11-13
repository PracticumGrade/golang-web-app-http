APP_VERSION ?= v0.0.1
BUILD_TIME := $(shell date -u +%Y%m%d%H%M%S)

.PHONY: build
build:
	CGO_ENABLED=0 go build	-ldflags "-X main.Version=${APP_VERSION} -X main.BuildTime=${BUILD_TIME}" \
							-o ../../bin/params-with-tests \
							./cmd/params-with-tests

.PHONY: run
run:
	CGO_ENABLED=0 go run	-ldflags "-X main.Version=${APP_VERSION} -X main.BuildTime=${BUILD_TIME}" \
    						./cmd/params-with-tests

.PHONY: tests
tests:
	go test	-timeout 5m \
			-race \
			-short \
    		-tags 'netgo' \
    		params-with-tests/internal/bookstore