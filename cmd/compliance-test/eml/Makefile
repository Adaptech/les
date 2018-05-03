SHELL := /bin/bash
LES_DIRECTORY := ../../les
LES_NODE_DIRECTORY := ../../les-node

.PHONY: test
test:
	go run main.go

.PHONY: setup
setup: build-les build-les-node build-api

.PHONY: teardown
teardown: 
	cd api \
	&& docker-compose down \
	&& rm -fr config \
	&& rm -fr src \
	&& rm -fr test \
	&& rm -fr web \
	&& rm Dockerfile \
	&& rm *.*

.PHONY: build-les
build-les:
	cd $(LES_DIRECTORY) \
	&& go build

.PHONY: build-les-node
build-les-node:
	cd $(LES_NODE_DIRECTORY) \
	&& go build

.PHONY: build-api
build-api:
	$(LES_DIRECTORY)/les convert \
	&& $(LES_NODE_DIRECTORY)/les-node -b Eventsourcing.eml.yaml \
	&& cd api && npm install && docker-compose up -d --force-recreate \
	&& docker-compose restart api # workaround for race condition (eventstore not yet available when API starts)

