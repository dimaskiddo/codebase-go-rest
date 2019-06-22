BUILD_OS           := linux
BUILD_OUTPUT       := main
SERVICE_NAME       := codebase-go-rest
SERVICE_PORT       := 3000
IMAGE_NAME         := codebase-go-rest
IMAGE_TAG          := latest
REBASE_URL         := "github.com/dimaskiddo/codebase-go-rest"
COMMIT_MSG         := "update improvement"

.PHONY:

.SILENT:

init:
	make clean
	dep init -v

ensure:
	make clean
	dep ensure -v

compile:
	make ensure
	CGO_ENABLED=0 GOOS=$(BUILD_OS) go build -a -o ./build/$(BUILD_OUTPUT) *.go
	echo "Build complete please check build directory."

run:
	CONFIG_ENV="DEV" CONFIG_FILE_PATH="./build/configs" CONFIG_LOG_LEVEL="DEBUG" CONFIG_LOG_SERVICE="$(SERVICE_NAME)" go run *.go

clean:
	rm -f ./build/$(BUILD_OUTPUT)
	rm -rf ./vendor

commit:
	make ensure
	make clean
	git add .
	git commit -am "$(COMMIT_MSG)"

rebase:
	rm -rf .git
	find . -type f -iname "*.go*" -exec sed -i '' -e "s%github.com/dimaskiddo/codebase-go-rest%$(REBASE_URL)%g" {} \;
	git init
	git remote add origin https://$(REBASE_URL).git

push:
	git push origin master

pull:
	git pull origin master

c-build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) --build-arg SERVICE_NAME=$(SERVICE_NAME) .

c-run:
	docker run -d -p $(SERVICE_PORT):$(SERVICE_PORT) --name $(SERVICE_NAME) --rm $(IMAGE_NAME):$(IMAGE_TAG)
	make docker-logs

c-shell:
	docker exec -it $(SERVICE_NAME) bash

c-stop:
	docker stop $(SERVICE_NAME)

c-logs:
	docker logs $(SERVICE_NAME)

c-push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)

c-clean:
	docker rmi -f $(IMAGE_NAME):$(IMAGE_TAG)
