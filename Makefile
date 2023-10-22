APP_ENTRY ?= simplehttpserver
APP_VERSION ?= rolling

BUILD_TIME := $(shell date -u +%Y%m%d%H%M%S)

.PHONY: dep_update
dep_update:
	go mod download

.PHONY: dep_tidy
dep_tidy: dep_update
	go mod tidy

.PHONY: build
build:
	CGO_ENABLED=0 go build	-a -v \
							-ldflags "-X main.Version=${APP_VERSION} -X main.BuildTime=${BUILD_TIME}" \
							-o ./bin/${APP_ENTRY} \
							./cmd/${APP_ENTRY}

.PHONY: tests
tests:
	go test	-timeout 5m \
			-race \
			-short \
    		-tags 'netgo' \
    		`go list ./...`

.PHONY: clean
clean:
	rm -rf ./bin/*