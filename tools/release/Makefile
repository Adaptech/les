SHELL := /bin/bash
VERSION ?= `date '+%Y-%m-%d-'``git show -s --format='commit-%h'`

.PHONY: all
all: build dist 

.PHONY: dist-les
dist-les:
	./go-multi-platform-build.sh github.com/Adaptech/les/cmd/les ../../releases/les/$(VERSION)

.PHONY: clean
clean:
	rm -fr ../../releases/les/$(VERSION)

