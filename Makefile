BINARY := systemintegrity
PREFIX ?= /usr/local

.PHONY: all build install clean

all: build

build:
	CGO_ENABLED=1 go build -o $(BINARY) ./cmds/systemintegrity

install: build
	install -d $(DESTDIR)$(PREFIX)/bin
	install -m 0755 $(BINARY) $(DESTDIR)$(PREFIX)/bin/$(BINARY)

clean:
	rm -f $(BINARY)
