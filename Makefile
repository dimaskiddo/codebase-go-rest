GO_OUTPUT ?= main
GO_EXPOSE_PORT ?= 3000
DOCKER_IMAGE_NAME ?= go-frame
DOCKER_IMAGE_VERSION ?= latest

git-push:
	make clean-go
	make clean-vendor
	git add .
	git commit -am "$(COMMIT_MSG)"
	git push origin master

go-dep:
	dep ensure -v
	dep check

go-build:
	make go-dep
	make clean-go
	CGO_ENABLED=0 GOOS=linux go build -a -o ./build/$(GO_OUTPUT) *.go

go-run:
	make go-dep
	CONFIG_PATH="./build/config" go run *.go

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) .

docker-run:
	docker run -d -p $(GO_EXPOSE_PORT):$(GO_EXPOSE_PORT) --name $(GO_OUTPUT) --rm $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

docker-stop:
	docker stop $(GO_OUTPUT)

docker-logs:
	docker logs $(GO_OUTPUT)

clean-go:
	rm -f ./build/$(GO_OUTPUT)

clean-vendor:
	rm -rf ./vendor/*
	touch ./vendor/.gitkeep

clean-docker:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

clean-all:
	make clean-go
	make clean-vendor
	make clean-docker
