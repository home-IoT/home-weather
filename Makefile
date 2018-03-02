include ./MANIFEST
include ./scripts/release.mk

DATE := $(shell date | sed 's/\ /_/g')

GOPATH ?= $$PWD/../../../..
GOOS ?= linux

SWAGGER_FILE=api/jupiter.yml

PKGS := $(shell go list ./... | grep -vF /vendor/)
CLI_PACKAGE=github.com/home-IoT/home-weather/internal/cli
MAIN_FILE=cmd/home-weather/main.go

# --- Repo 

initialize: clean swagger-generate
	dep init
	$(MAKE) dep

clean:
	mkdir -p bin
	rm -rf ./bin/*
	rm -rf ./pkg/*

# --- Tools

get-tools:
	go get -u github.com/golang/lint/golint
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# --- Swagger

get-swagger:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger-serve:
	swagger serve $(SWAGGER_FILE)

swagger-validate:
	swagger validate $(SWAGGER_FILE)

swagger-gen: clean
	swagger generate client -f $(SWAGGER_FILE) -t gen

# --- Common Go

go-dep:
	dep ensure

go-dep-status:
	dep status

go-dep-clean: clean
	mkdir -p ./bin
	rm -rf ./vendor/*
	dep ensure

go-fmt:
	@go fmt $(PKGS)

go-validate:
	@echo go vet
	@go vet $(PKGS)
	@echo golint
	@golint -set_exit_status $(PKGS)

# --- Home Weather CLI

go-build-linux:
	@echo "build linux binary"
	$(MAKE) go-build GOOS=linux GOARCH=amd64 TARGET=$(PROJECT)-linux-amd64

go-build-pi:
	@echo "build linux binary for raspberry pi"
	$(MAKE) go-build GOOS=linux GOARCH=arm GOARM=7 TARGET=$(PROJECT)-linux-arm7

go-build-windows:
	@echo "build windows binary"
	$(MAKE) go-build GOOS=windows GOARCH=386 TARGET=$(PROJECT)-windows-386.exe

go-build-mac:
	@echo "build Mac binary"
	$(MAKE) go-build GOOS=darwin GOARCH=amd64 TARGET=$(PROJECT)-darwin-amd64

TARGET ?= $(EXECUTABLE)

go-build: 
	go build -ldflags="-X $(CLI_PACKAGE).GitRevision=$(shell git rev-parse HEAD) -X $(CLI_PACKAGE).BuildVersion=$(VERSION) -X $(CLI_PACKAGE).BuildTime=$(DATE)" -i -o ./bin/$(TARGET) $(MAIN_FILE)

go-build-all: go-build-pi go-build-linux go-build-windows go-build-mac

# --- Release

go-release-all: clean 
	$(MAKE) go-build-all
	mkdir -p ./release
	rm -rf ./release/*
	chmod +x bin/*
	cp ./bin/* ./release
	for bf in ./release/*; do shasum -a 256 "$$bf" > "$$bf".sha256; done


