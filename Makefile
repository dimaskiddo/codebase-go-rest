GO_OUTPUT ?= main
GO_EXPOSE_PORT ?= 3000
DOCKER_IMAGE_NAME ?= frame-go
DOCKER_IMAGE_VERSION ?= latest

git-push:
	make go-dep-init
	make go-dep-clean
	make go-clean
	git add .
	git commit -am "$(COMMIT_MSG)"
	git push origin master

git-pull:
	git pull origin master

go-dep-init:
	make go-dep-clean
	rm -f Gopkg.toml Gopkg.lock
	dep init -v

go-dep-ensure:
	make go-dep-clean
	dep ensure -v

go-dep-clean:
	rm -rf ./vendor

go-build:
	make go-clean
	make go-dep-ensure
	CGO_ENABLED=0 GOOS=linux go build -a -o ./build/$(GO_OUTPUT) *.go

go-run:
	CONFIG_FILE="dev" CONFIG_PATH="./build/configs" go run *.go

go-clean:
	rm -f ./build/$(GO_OUTPUT)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) .

docker-run:
	docker run -d -p $(GO_EXPOSE_PORT):$(GO_EXPOSE_PORT) --name $(GO_OUTPUT) --rm $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

docker-stop:
	docker stop $(GO_OUTPUT)

docker-logs:
	docker logs $(GO_OUTPUT)

docker-clean:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

clean:
	make go-clean
	make go-dep-clean
	make docker-clean
