.PHONY: all

all: bin/checker bin/runner

bin/checker:
	go build -o bin/checker checker/main.go

bin/runner:
	go build -o bin/runner runner/main.go