PKG_PREFIX := github.com/vitalvas/cloudrouter
GO_BUILDINFO = -s -w

.PHONY: $(MAKECMDGOALS)

all: cloudrouter-dns

cloudrouter-dns:
	GOOS=linux go build -ldflags "$(GO_BUILDINFO)" -o bin/cloudrouter-dns $(PKG_PREFIX)/app/dns

clean:
	rm -Rf bin/*
