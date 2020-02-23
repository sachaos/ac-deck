.PHONY: prepare build install release

ARTIFACTS_DIR=artifacts/${VERSION}
GITHUB_USERNAME=sachaos
BUILD_OPTION=-ldflags "-X github.com/sachaos/atcoder/cmd.version=${VERSION}"

prepare:
	go get github.com/rakyll/statik
	go generate

build: prepare
	go build $(BUILD_OPTION)

install: prepare
	go install $(BUILD_OPTION)

release: prepare
	GOOS=windows GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/atcoder_windows_amd64 $(BUILD_OPTION)
	GOOS=darwin GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/atcoder_darwin_amd64 $(BUILD_OPTION)
	GOOS=linux GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/atcoder_linux_amd64 $(BUILD_OPTION)
	ghr -u $(GITHUB_USERNAME) -t $(shell cat github_token) --replace ${VERSION} $(ARTIFACTS_DIR)
