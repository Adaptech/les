SHELL := /bin/bash

.PHONY: test
test:
	go run main.go

.PHONY: setup
setup: build-les build-les-node build-api

.PHONY: teardown
teardown: 
	pushd api \
	&& docker-compose down \
	&& popd \
	&& rm -fr api

.PHONY: build-les
build-les:
	cd ../les \
	&& go build

.PHONY: build-les-node
build-les-node:
	cd ../les-node \
	&& go build

.PHONY: build-api
build-api:
	../les/les convert \
	&& ../les-node/les-node -b \
	&& cd api && npm install && docker-compose up -d \
	&& docker-compose restart api # workaround for race condition (eventstore not yet available when API starts)

