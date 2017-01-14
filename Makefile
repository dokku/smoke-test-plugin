GO_REPO_ROOT := /go/src/github.com/dokku/smoke-test-plugin
BUILD_IMAGE := golang:1.7.1

.PHONY: build-in-docker clean

build-in-docker:
	docker run --rm \
		-v $$PWD:$(GO_REPO_ROOT) \
		-w $(GO_REPO_ROOT) \
		$(BUILD_IMAGE) \
		bash -c "make commands" || exit $$?

commands: **/commands.go
	cd src && \
		go build -a -o ../commands
	ln -sf ./commands ./pre-deploy

clean:
	rm -f commands pre-deploy
