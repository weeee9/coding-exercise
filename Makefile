ROJECT="CODING EXERCISE"

SOURCE ?= $(shell find . -type f -name '*.go' -not -path '*/generated/*')

all: test

test:
	@echo $(SOURCE)
	go test -v -tags="json1" ./...
	@echo "===\033[32m OK \033[0m==="