.PHONY: prepare build install

BUILD_OPTION=-o acd -ldflags "-X github.com/sachaos/ac-deck/cmd.version=${VERSION}"

prepare:
	go install github.com/rakyll/statik@latest
	go generate

build: prepare
	go build $(BUILD_OPTION)

install: build
	mv acd /usr/local/bin

test: prepare
	go test ./...
