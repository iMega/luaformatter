REPO = github.com/imega/luaformatter
IMG = imega/luaformatter
TAG = latest
CWD = /go/src/$(REPO)
GO_IMG = golang:1.14.6-alpine3.12

unit:
	@docker run --rm \
		-w $(CWD) \
		-v $(CURDIR):$(CWD) \
		$(GO_IMG) sh -c "go list ./... | grep -v 'vendor\|tests' | xargs go test"

lint:
	@-docker run --rm -t -v $(CURDIR):$(CWD) -w $(CWD) -e GOFLAGS=-mod=vendor \
		golangci/golangci-lint golangci-lint run -v

test: lint unit
