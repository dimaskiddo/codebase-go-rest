BUILD_CGO_ENABLED  := 0
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
	GO111MODULE=on go mod init

init-dist:
	mkdir -p dist

vendor:
	make clean
	GO111MODULE=on go mod vendor

release:
	make vendor
	make clean-dist
	goreleaser --snapshot --skip-publish --rm-dist
	echo "Release '$(SERVICE_NAME)' complete, please check dist directory."

publish:
	make vendor
	make clean-dist
	GITHUB_TOKEN=$(GITHUB_TOKEN) goreleaser --rm-dist
	echo "Publish '$(SERVICE_NAME)' complete, please check your repository releases."

build:
	make vendor
	CGO_ENABLED=$(BUILD_CGO_ENABLED) go build -ldflags="-s -w" -a -o $(SERVICE_NAME) cmd/main/main.go
	echo "Build '$(SERVICE_NAME)' complete."

run:
	go run cmd/main/*.go

clean-dist:
	rm -rf dist

clean:
	make clean-dist
	rm -rf vendor

commit:
	make vendor
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
	make c-logs

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
