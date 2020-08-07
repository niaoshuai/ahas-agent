# Go parameters
GO           = go
TIMEOUT_UNIT = 5m

.PHONY: all
all: test build

.PHONY: build
build:
	$(GO) build -v ./cmd/...

.PHONY: test
test:
	$(GO) test -timeout $(TIMEOUT_UNIT) -v ./test/...

.PHONY: clean
clean:
	$(GO) clean
	@rm -rf test/tests.* test/coverage.*
	@rm -rf ahas-agent

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -v ./cmd/...
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -v ./cmd/...