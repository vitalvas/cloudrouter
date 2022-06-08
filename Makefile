PKG_PREFIX := github.com/vitalvas/cloudrouter
GO_BUILDINFO = -s -w

.PHONY: $(MAKECMDGOALS)

all: \
	cloudrouter-dns \
	cloudrouter-dhcp4-server \
	cloudrouter-netconfig \
	cloudrouter-netwatch

cloudrouter-dns:
	GOOS=linux go build -ldflags "$(GO_BUILDINFO)" -o bin/cloudrouter-dns $(PKG_PREFIX)/app/dns

cloudrouter-dhcp4-server:
	GOOS=linux go build -ldflags "$(GO_BUILDINFO)" -o bin/cloudrouter-dhcp4-server $(PKG_PREFIX)/app/dhcp4-server

cloudrouter-netconfig:
	GOOS=linux go build -ldflags "$(GO_BUILDINFO)" -o bin/cloudrouter-netconfig $(PKG_PREFIX)/app/netconfig

cloudrouter-netwatch:
	GOOS=linux go build -ldflags "$(GO_BUILDINFO)" -o bin/cloudrouter-netwatch $(PKG_PREFIX)/app/netwatch

clean:
	rm -Rf bin/*
