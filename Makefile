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

.PHONY: test-all
test-all: unit-test compliance-test-emd compliance-test-eml

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: compliance-test-eml
compliance-test-eml:
	cd cmd/compliance-test/eml \
	&& sleep 1 \
	&& make setup && sleep 1 && make test \
	&& make teardown

.PHONY: compliance-test-emd
compliance-test-emd:
	cd cmd/compliance-test/emd \
	&& sleep 1 \
	&& make setup && sleep 1 && make test \
	&& make teardown
