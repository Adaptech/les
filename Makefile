SHELL := /bin/bash

.PHONY: build
build: build-les build-les-node

.PHONY: build-les
build-les:
	cd cmd/les \
	&& go install

.PHONY: build-les-node
build-les-node:
	cd cmd/les-node \
	&& go install
