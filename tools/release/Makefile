SHELL := /bin/bash
VERSION ?= `date '+%Y-%m-%d-'``git show -s --format='commit-%h'`

.PHONY: dist
dist: dist-les dist-les-node

.PHONY: dist-les
dist-les:
	./go-multi-platform-build.sh github.com/Adaptech/les/cmd/les ../../releases/les/$(VERSION)

.PHONY: dist-les-node
dist-les-node:
	./go-multi-platform-build.sh github.com/Adaptech/les/cmd/les-node ../../releases/les-node/$(VERSION)

.PHONY: clean
clean:
	rm -fr ../../releases/les/$(VERSION) \
	&& rm -fr ../../releases/les-node/$(VERSION)

