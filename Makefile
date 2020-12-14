include .env

BINARY := cronparse
VERSION := $(shell git describe --always --dirty --tags 2>/dev/null || echo "undefined")
ECHO := echo

.NOTPARALLEL:

.PHONY: clean
clean:
	rm -f $(BINARY)

.PHONY: build
build: clean $(BINARY)

.PHONY: test
test:
	@ $(ECHO) "\033[36mRunning test suite in Ginkgo\033[0m"
	$(GINKGO) -v -p -race -randomizeAllSpecs ./pkg/... ./cmd/...
	@ $(ECHO)

.PHONY: fmt
fmt:
	$(GO) fmt ./pkg/... ./cmd/...

.PHONY: vet
vet:
	$(GO) vet -tags dev -composites=false ./pkg/... ./cmd/...

.PHONY: lint
lint:
	@ $(ECHO) "\033[36mLinting code\033[0m"
	$(LINTER) run --disable-all --build-tags dev \
                --exclude-use-default=false \
                --enable=govet \
                --enable=ineffassign \
                --enable=deadcode \
                --enable=golint \
                --enable=goconst \
                --enable=gofmt \
                --enable=goimports \
                --skip-dirs=pkg/client/ \
                --deadline=120s \
                --tests ./...
	@ $(ECHO)

.PHONY: check
check: fmt vet lint test

$(BINARY):
	GO111MODULE=on CGO_ENABLED=0 $(GO) build -o $(BINARY) -ldflags="-X main.VERSION=${VERSION}" github.com/gargath/cronparse/cmd
