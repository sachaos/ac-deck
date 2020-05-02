.PHONY: prepare build install

BUILD_OPTION=-o acd -ldflags "-X github.com/sachaos/ac-deck/cmd.version=${VERSION}"

prepare:
	go get github.com/rakyll/statik
	go generate

build: prepare
	go build $(BUILD_OPTION)

install: build
	mv acd /usr/local/bin

test: prepare
	go test ./...
