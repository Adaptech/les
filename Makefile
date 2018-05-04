SHELL := /bin/bash

.PHONY: install
install: build-les build-les-node build-les-viz

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

build-les-viz:
	cd cmd/les-viz \
	&& go get \
	&& go install

.PHONY: test-all
test-all: unit-test test-samples-are-valid test-emd-compliance test-eml-compliance
	echo "COMPLIANCE TESTS PASS."

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: test-eml-compliance
test-eml-compliance:
	cd cmd/compliance-test/eml \
	&& sleep 2 \
	&& make setup && sleep 2 && make test \
	&& make teardown

.PHONY: test-emd-compliance
test-emd-compliance:
	cd cmd/compliance-test/emd \
	&& sleep 2 \
	&& make setup && sleep 2 && make test \
	&& make teardown

.PHONY: test-samples-are-valid
test-samples-are-valid:
	cd samples/consentaur \
	&& les validate \
	&& echo "samples/email" \
	&& cd ../email \
	&& les validate \
	&& echo "samples/helloworld" \
	&& cd ../helloworld \
	&& les validate \
	&& echo "samples/inventory" \
	&& cd ../inventory \
	&& les validate \
	&& echo "samples/subscriptions" \
	&& cd ../subscriptions \
	&& les validate \
	&& echo "samples/timesheets" \
	&& cd ../timesheets \
	&& les validate \
	&& echo "samples/todolist" \
	&& cd ../todolist \
	&& les validate \
	&& echo "samples//users" \
	&& cd ../users \
	&& les validate Eventsourcing.eml.yaml \
	&& echo "samples/veggiesgalore" \
	&& cd ../veggiesgalore \
	&& les validate \
	&& echo "SAMPLES VALIDATION TESTS PASS."

