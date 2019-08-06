.PHONY: deps build
docker-build:
	docker run --rm -i -t -e TRAVIS_GO_VERSION=$(TRAVIS_GO_VERSION) -e TRAVIS_BUILD_NUMBER=$(TRAVIS_BUILD_NUMBER) -e TRAVIS_TAG=$(TRAVIS_TAG) -v `pwd`:/fs/src/github.com/cad/ovpm -w /fs/src/github.com/cad/ovpm cadthecoder/ovpm-builder:latest
docker-build-shell:
	docker run --rm -i -t -e TRAVIS_GO_VERSION=$(TRAVIS_GO_VERSION) -e TRAVIS_BUILD_NUMBER=$(TRAVIS_BUILD_NUMBER) -e TRAVIS_TAG=$(TRAVIS_TAG) -v `pwd`:/fs/src/github.com/cad/ovpm -w /fs/src/github.com/cad/ovpm cadthecoder/ovpm-builder:latest /bin/bash
docker-build-shell-local:
	$(eval RET = $(shell docker start ovpm-build))
	@if [ "$(RET)" = "ovpm-build" ]; then\
		docker attach ovpm-build; \
	else \
		docker run --name ovpm-build -i -t -v `pwd`:/fs/src/github.com/cad/ovpm -w /fs/src/github.com/cad/ovpm golang:1.12 /bin/bash; \
	fi

deps:
	# grpc related dependencies
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	# webui related dependencies
	pacman -Sy npm
