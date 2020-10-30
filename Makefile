REPO = github.com/imega/luaformatter
IMG = imega/luaformatter
TAG = latest
CWD = /go/src/$(REPO)
GO_IMG = golang:1.15-alpine3.12

unit:
	@docker run --rm -w $(CWD) -v $(CURDIR):$(CWD) \
		$(GO_IMG) sh -c "go list ./... | xargs go test -vet=off"

lint:
	@-docker run --rm -t -v $(CURDIR):$(CWD) -w $(CWD) \
		golangci/golangci-lint golangci-lint run

test:
	@docker run --rm -w $(CWD) -v $(CURDIR):$(CWD) \
		$(GO_IMG) sh -c "apk add --upd bash && go build -o /bin/app && tests/test.sh"

release: test
