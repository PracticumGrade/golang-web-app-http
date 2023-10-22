APP_VERSION ?= v0.0.1
BUILD_TIME := $(shell date -u +%Y%m%d%H%M%S)

.PHONY: build
build:
	CGO_ENABLED=0 go build	-ldflags "-X main.Version=${APP_VERSION} -X main.BuildTime=${BUILD_TIME}" \
							-o ../../bin/testwebapp \
							./cmd/testwebapp

.PHONY: run
run:
	CGO_ENABLED=0 go run	-ldflags "-X main.Version=${APP_VERSION} -X main.BuildTime=${BUILD_TIME}" \
    						./cmd/testwebapp

.PHONY: tests
tests:
	go test	-timeout 5m \
			-race \
			-short \
    		-tags 'netgo' \
    		testwebapp/internal/handlers/simpleprinter \
			testwebapp/internal/server_test