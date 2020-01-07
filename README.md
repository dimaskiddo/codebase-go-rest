# Codebase for REST API in Go

This codebase help a new Go developer in developing a REST API using Go.
This codebase included with some pre bundle service or helper like authentication, logging, and predefined configuration.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Prequisites packages:
* Go (Go Programming Language)
* Dep (Go Dependencies Management Tool)
* Make (Automated Execution using Makefile)

Optional packages:
* GoReleaser (Go Automated Binaries Build)
* Docker (Application Containerization)

### Installing

Below is the instructions to make this codebase running:
* Create a Go Workspace directory and export it as the extended GOPATH directory
```
cd <your_go_workspace_directory>
export GOPATH=$GOPATH:"`pwd`"
```
* Under the Go Workspace directory create a source directory
```
mkdir -p src/<your_repository_domain>/<your_username>/<your_repository>
```
* Move to the created directory and pull codebase
```
cd src/<your_repository_domain>/<your_username>/<your_repository>
git clone -b master https://github.com/dimaskiddo/codebase-go-rest.git .
```
* Run following command to change the codebase repository URL to your own
```
make rebase REBASE_URL="<your_repository_domain>/<your_username>/<your_repository>"
```
* Run following command to renew and pull dependecies package
```
make init
make vendor
```
* Until this step you already can run this code by using this command
```
make run
```

## Running The Tests

Currently the test is not ready yet :)

## Deployment

**Make sure your your go script does not relate to github.com/dimaskiddo/codebase-go-rest anymore**.
To build this codebase to binaries for distribution purposes you can run following command:
```
make release
```
The build result will shown in build directory

## Built With

* [Go](https://golang.org/) - Go Programming Languange
* [Dep](https://github.com/golang/dep) - Go Dependency Management Tool
* [GoReleaser](https://github.com/goreleaser/goreleaser) - Go Automated Binaries Build
* [Make](https://www.gnu.org/software/make/) - GNU Make Automated Execution
* [Docker](https://www.docker.com/) - Application Containerization

## Authors

* **Dimas Restu Hidayanto** - *Initial Work* - [DimasKiddo](https://github.com/dimaskiddo)

See also the list of [contributors](https://github.com/dimaskiddo/codebase-go-rest/contributors) who participated in this project

## Annotation

You can seek more information for the make command parameters in the [Makefile](https://raw.githubusercontent.com/dimaskiddo/codebase-go-rest/master/Makefile)
