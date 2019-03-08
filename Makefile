GO_BUILD_OS						:= linux
GO_BUILD_OUTPUT 			:= main
DOCKER_IMAGE_NAME			:= codebase-go-rest
DOCKER_IMAGE_TAG			:= latest
DOCKER_SERVICE_NAME		:= codebase-go-rest
DOCKER_SERVICE_PORT		:= 3000
GIT_REBASE_URL        := "github.com/dimaskiddo/codebase-go-rest"
GIT_COMMIT_MSG        := "update improvement"

.PHONY:

.SILENT:

git-push:
	make go-dep-init
	make go-dep-clean
	make go-clean
	git add .
	git commit -am "$(GIT_COMMIT_MSG)"
	git push origin master

git-pull:
	git pull origin master

git-rebase:
	rm -rf .git
	sed -i -e "s%github.com/dimaskiddo/codebase-go-rest%$(GIT_REBASE_URL)%g" *.go
	sed -i -e "s%github.com/dimaskiddo/codebase-go-rest%$(GIT_REBASE_URL)%g" controller/*.go
	sed -i -e "s%github.com/dimaskiddo/codebase-go-rest%$(GIT_REBASE_URL)%g" model/*.go
	sed -i -e "s%github.com/dimaskiddo/codebase-go-rest%$(GIT_REBASE_URL)%g" service/*.go

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
	CGO_ENABLED=0 GOOS=$(GO_BUILD_OS) go build -a -o ./build/$(GO_BUILD_OUTPUT) *.go

go-run:
	CONFIG_ENV="DEV" CONFIG_FILE_PATH="./build/configs" CONFIG_LOG_LEVEL="DEBUG" CONFIG_LOG_SERVICE="$(DOCKER_SERVICE_NAME)" go run *.go

go-clean:
	rm -f ./build/$(GO_BUILD_OUTPUT)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG) --build-arg SERVICE_NAME=$(DOCKER_SERVICE_NAME) .

docker-run:
	docker run -d -p $(DOCKER_SERVICE_PORT):$(DOCKER_SERVICE_PORT) --name $(DOCKER_SERVICE_NAME) --rm $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)
	make docker-logs

docker-exec:
	docker exec -it $(DOCKER_SERVICE_NAME) bash

docker-stop:
	docker stop $(DOCKER_SERVICE_NAME)

docker-logs:
	docker logs $(DOCKER_SERVICE_NAME)

docker-push:
	docker push $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

docker-clean:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

clean:
	make go-clean
	make go-dep-clean
	make docker-clean
