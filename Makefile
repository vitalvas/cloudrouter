PKG_PREFIX := github.com/vitalvas/cloudrouter
APPS := $(patsubst app/%,cloudrouter-%,$(wildcard app/*))
GO_BUILDINFO = -s -w

.PHONY: $(MAKECMDGOALS)

all: clean test $(APPS)

test:
	golangci-lint run

build: $(APPS)

cloudrouter-%:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/$@-amd64 $(PKG_PREFIX)/app/$*
	GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "$(GO_BUILDINFO)" -o build/$@-armv7 $(PKG_PREFIX)/app/$*

clean:
	rm -Rf build/*

update:
	go clean -modcache
	go get -v -u ./...
	go mod tidy -v
