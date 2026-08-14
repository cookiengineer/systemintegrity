BINARY := systemintegrity
PREFIX ?= /usr/local

.PHONY: all build install clean test test-integration

all: build

build:
	CGO_ENABLED=1 go build -o $(BINARY) ./cmds/systemintegrity

install: build
	install -d $(DESTDIR)$(PREFIX)/bin
	install -m 0755 $(BINARY) $(DESTDIR)$(PREFIX)/bin/$(BINARY)

test:
	go test ./types/... ./structs/... ./caches/...

test-integration:
	bash ./test-integration.sh

clean:
	rm -f $(BINARY)
