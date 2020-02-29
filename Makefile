.PHONY: prepare build install

BUILD_OPTION=-ldflags "-X github.com/sachaos/atcoder/cmd.version=${VERSION}"

prepare:
	go get github.com/rakyll/statik
	go generate

build: prepare
	go build $(BUILD_OPTION)

install: prepare
	go install $(BUILD_OPTION)

test: prepare
	go test ./...
