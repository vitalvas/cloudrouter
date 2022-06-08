PKG_PREFIX := github.com/vitalvas/cloudrouter
GO_BUILDINFO = -s -w

.PHONY: $(MAKECMDGOALS)

all: \
	clean \
	cloudrouter-dns \
	cloudrouter-dhcp4-server \
	cloudrouter-netconfig \
	cloudrouter-netwatch

cloudrouter-%:
	GOOS=linux go build -ldflags "$(GO_BUILDINFO)" -o bin/$@ $(PKG_PREFIX)/app/$*

clean:
	rm -Rf bin/*
