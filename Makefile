GO_REPO_ROOT := /go/src/github.com/dokku/smoke-test-plugin
BUILD_IMAGE := golang:1.7.1

.PHONY: build-in-docker build clean

build-in-docker:
	docker run --rm \
		-v $$PWD:$(GO_REPO_ROOT) \
		-w $(GO_REPO_ROOT) \
		$(BUILD_IMAGE) \
		bash -c "make build" || exit $$?

build: commands triggers
triggers: pre-deploy
commands: **/**/commands.go
	go build -a -o commands ./src/commands/commands.go

pre-deploy: **/**/pre-deploy.go
	go build -a -o pre-deploy ./src/triggers/pre-deploy.go

clean:
	rm -f commands pre-deploy
