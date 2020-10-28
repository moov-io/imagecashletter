PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+(-[a-zA-Z0-9]*)?)' version.go)

.PHONY: build build-server docker release check

build: check build-server build-webui

build-server:
	CGO_ENABLED=1 go build -o ./bin/server github.com/moov-io/imagecashletter/cmd/server

build-webui:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./cmd/webui/assets/wasm_exec.js
	GOOS=js GOARCH=wasm go build -o ./cmd/webui/assets/imagecashletter.wasm github.com/moov-io/imagecashletter/cmd/webui/icl/
	CGO_ENABLED=0 go build -o ./bin/webui ./cmd/webui

check:
	go fmt ./...
	@mkdir -p ./bin/

.PHONY: client
client:
ifeq ($(OS),Windows_NT)
	@echo "Please generate ./client/ on macOS or Linux, currently unsupported on windows."
else
# Versions from https://github.com/OpenAPITools/openapi-generator/releases
	@chmod +x ./openapi-generator
	@rm -rf ./client
	OPENAPI_GENERATOR_VERSION=4.2.0 ./openapi-generator generate -i openapi.yaml -g go -o ./client
	rm -f client/go.mod client/go.sum
	go fmt ./...
	go build github.com/moov-io/imagecashletter/client
	go test ./client
endif

.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	@echo "Skipping cleanup on Windows, currently unsupported."
else
	@rm -rf ./bin/ openapi-generator-cli-*.jar
endif

dist: clean client build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/imagecashletter.exe github.com/moov-io/imagecashletter/cmd/server
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/imagecashletter-$(PLATFORM)-amd64 github.com/moov-io/imagecashletter/cmd/server
endif

docker: clean docker-hub docker-fuzz docker-webui

docker-hub:
	docker build --pull -t moov/imagecashletter:$(VERSION) -f Dockerfile .
	docker tag moov/imagecashletter:$(VERSION) moov/imagecashletter:latest

docker-fuzz:
	docker build --pull -t moov/imagecashletterfuzz:$(VERSION) . -f Dockerfile.fuzz
	docker tag moov/imagecashletterfuzz:$(VERSION) moov/imagecashletterfuzz:latest

docker-openshift:
	docker build --pull -t quay.io/moov/imagecashletter:$(VERSION) -f Dockerfile.openshift --build-arg VERSION=$(VERSION) .
	docker tag quay.io/moov/imagecashletter:$(VERSION) quay.io/moov/imagecashletter:latest

docker-webui:
	docker build --pull -t moov/imagecashletter-webui:$(VERSION) -f Dockerfile.webui .
	docker tag moov/imagecashletter-webui:$(VERSION) moov/imagecashletter-webui:latest

release: docker AUTHORS
	go vet ./...
	go test -coverprofile=cover-$(VERSION).out ./...
	git tag -f $(VERSION)

release-push:
	docker push moov/imagecashletter:$(VERSION)
	docker push moov/imagecashletter:latest
	docker push moov/imagecashletter-webui:$(VERSION)
	docker push moov/imagecashletter-webui:latest
	docker push moov/imagecashletterfuzz:$(VERSION)

quay-push:
	docker push quay.io/moov/imagecashletter:$(VERSION)
	docker push quay.io/moov/imagecashletter:latest

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@
