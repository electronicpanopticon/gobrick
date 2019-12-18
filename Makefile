# Basic go commands
GOCMD=go
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test

.PHONY: lint release test tidy

all: test

lint:
	golint ./...

release: tidy
		$(GOCMD) fmt

test:
		$(GOTEST) -v ./... -cover

tidy:
		$(GOMOD) tidy
		$(GOCMD) list -m all
