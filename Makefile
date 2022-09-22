PKG_PREFIX := github.com/vitalvas/cloudrouter
APPS := $(patsubst app/%,cloudrouter-%,$(wildcard app/*))
GO_BUILDINFO = -s -w

.PHONY: $(MAKECMDGOALS)

all: clean $(APPS)

cloudrouter-%:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/$@ $(PKG_PREFIX)/app/$*

clean:
	rm -Rf build/*
