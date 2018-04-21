SHELL := /bin/bash

.PHONY: install
install: build-les build-les-node

.PHONY: build-les
build-les:
	cd cmd/les \
	&& go get \
	&& go install 


.PHONY: build-les-node
build-les-node:
	cd cmd/les-node \
	&& go get \
	&& go install

