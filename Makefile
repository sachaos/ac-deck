.PHONY: prepare build install

BUILD_OPTION=-o acd -ldflags "-X github.com/sachaos/ac-deck/cmd.version=${VERSION}"

build:
	go build $(BUILD_OPTION)

install: build
	mv acd /usr/local/bin

test:
	go test ./...
