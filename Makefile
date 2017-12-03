ifndef GOPATH
$(warning You need to set up a GOPATH. Run "go help gopath".)
endif

version:=0.0.1
pkg_name:=go2
git:=$(strip $(shell which git 2> /dev/null))
go:=$(strip $(shell which go 2> /dev/null))

ifneq ($(go),)
		packages:=$(shell $(go) list ./... | grep -v /vendor/)
endif

ifeq ($(GOOS),)
		GOOS:=linux
endif

ifeq ($(GOARCH),)
		GOARCH:=amd64
endif


repo_branch:=$(subst refs/heads/,,$(shell $(git) symbolic-ref HEAD 2> /dev/null))
repo_branch:=$(subst /,-,$(repo_branch))

.PHONY: build clean format run cov test vendor

default: build

build:
		CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(go) build -v -ldflags "-X main.version=$(version) -X main.repoBranch=$(repo_branch)" -o ./bin/$(GOOS)_$(GOARCH)/$(pkg_name)$(BIN_EXTENSION)

clean:
		rm -rf bin

format:
		$(go) fmt $(packages)

run: build
		./bin/linux_amd64/$(pkg_name) -c etc/config-default.yaml

test: format
		$(go) vet $(packages)
		$(go) test -race -cover $(packages)

cov:
		echo "mode: set" > coverage-all.out
		$(foreach pkg,$(packages),\
				go test -coverprofile=coverage.out $(pkg);\
				tail -n +2 coverage.out >> coverage-all.out;)
		go tool cover -func=coverage-all.out

vendor:
		rm -rf vendor
		dep ensure
