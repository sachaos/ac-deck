.PHONY: prepare build install release

ARTIFACTS_DIR=artifacts/${VERSION}
GITHUB_USERNAME=sachaos

prepare:
	go get github.com/rakyll/statik
	go generate

build: prepare
	go build

install: prepare
	go install

release: prepare
	GOOS=windows GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/atcoder_windows_amd64
	GOOS=darwin GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/atcoder_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/atcoder_linux_amd64
	ghr -u $(GITHUB_USERNAME) -t $(shell cat github_token) --replace ${VERSION} $(ARTIFACTS_DIR)
