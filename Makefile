.PHONY: all clean test

GOPATH=$(shell go env GOPATH)

GDFLAGS ?= $(GDFLAGS:)
PACKAGES = $(shell glide novendor)
MAINFILES = main.go routes.go

all: test

deps:
	@echo "===> Vendoring dependencies"
	@glide install

test: lint
	@echo "===> Testing"
	@echo "mode: count" > coverage-all.out
	@ $(foreach pkg, $(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)

coverage: test
	@echo "===> Creating coverage report"
	go tool cover -html=coverage-all.out

clean:
	@echo "===> Cleaning $$GOPATH"
	@go clean $(PACKAGES)
	@rm *.out

fmt:
	@echo "===> Formatting"
	@go fmt $(PACKAGES)

lint:
	@echo "===> Linting with vet"
	@go vet $(PACKAGES)

build: fmt lint
	@echo "===> Building"
	go build -o app $(MAINFILES)

run: lint
	@echo "===> Running Server"
	go run $(MAINFILES)