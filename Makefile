PKG_PREFIX := github.com/vitalvas/cloudrouter
GO_BUILDINFO = -s -w

.PHONY: $(MAKECMDGOALS)

all: cloudrouter-dns cloudrouter-netconfig

cloudrouter-dns:
	GOOS=linux go build -ldflags "$(GO_BUILDINFO)" -o bin/cloudrouter-dns $(PKG_PREFIX)/app/dns

cloudrouter-netconfig:
	GOOS=linux go build -ldflags "$(GO_BUILDINFO)" -o bin/cloudrouter-netconfig $(PKG_PREFIX)/app/netconfig

clean:
	rm -Rf bin/*
