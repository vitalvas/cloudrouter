PKG_PREFIX := github.com/vitalvas/cloudrouter
APPS := $(patsubst app/%,cloudrouter-%,$(wildcard app/*))
PLUGINS := $(patsubst plugin/%,plugin-%,$(wildcard plugin/*))
GO_BUILDINFO = -s -w

.PHONY: $(MAKECMDGOALS)

all: clean $(APPS) $(PLUGINS)

cloudrouter-%:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/$@ $(PKG_PREFIX)/app/$*

plugin-%:
	GOOS=linux GOARCH=amd64 go build -buildmode=plugin -o build/$@.so $(PKG_PREFIX)/plugin/$*

clean:
	rm -Rf build/*
