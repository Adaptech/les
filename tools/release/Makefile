SHELL := /bin/bash
VERSION := 0.10.4-alpha

.PHONY: dist
dist: dist-les dist-les-node installation-instructions

.PHONY: dist-les
dist-les:
	./go-multi-platform-build.sh github.com/Adaptech/les/cmd/les ../../releases/les/$(VERSION)

.PHONY: dist-les-node
dist-les-node:
	./go-multi-platform-build.sh github.com/Adaptech/les/cmd/les-node ../../releases/les-node/$(VERSION)

.PHONY: installation-instructions
installation-instructions:
	./make-installation-instructions.sh $(VERSION) ../../INSTALL.md

.PHONY: clean
clean:
	rm -fr ../../releases/les/$(VERSION) \
	&& rm -fr ../../releases/les-node/$(VERSION) \
	&& rm ../../INSTALL.md
