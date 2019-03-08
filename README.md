# Codebase for REST API in Go

This codebase help a new Go developer in developing a REST API using Go.
This codebase included with some pre bundle service or helper like authentication, logging, and predefined configuration.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
See deployment for notes on how to deploy the project on a live system.

### Prerequisites

To run this codebase make sure that you already have below package installed:
- Go
- Dep (For Go Dependencies Management Tool)
- Make (For Automated Execution using Makefile)

Optional package:
- GoReleaser (Optional, For Go Automated Binaries Build)
- Docker (Optional, For Application Containerization)

### Installing

Below is the instructions to make this codebase running:
- Create a Go Workspace directory and export it as the extended GOPATH directory
```
cd <your_go_workspace_directory>
export GOPATH=$GOPATH:"`pwd`"
```
- Under the Go Workspace directory create a source directory like following command
```
mkdir -p src/github.com/dimaskiddo/codebase-go-rest
```
- Move to newly created directory and pull the codebase there
```
cd src/github.com/dimaskiddo/codebase-go-rest
git clone -b master https://github.com/dimaskiddo/codebase-go-rest.git .
```
- If you want to use your own repository directory please update any package related to this repository in any Go script with yours then run following command
```
make go-dep-init
```
- Run following command to pull dependencies used by the codebase it self
```
make go-dep-ensure
```
- In this step the codebase is ready to run, you can run it with the following command
```
make go-run
```

## Running The Tests

Currently the test is not ready yet :)

## Deployment

To deploy your application in any environment, since this is a codebase please make sure that there is no more package related with this repository *github.com/dimaskiddo/codebase-go-rest* in your go script, change it to your own repository

## Built With

* [Go](https://golang.org/) - Go Programming Languange
* [Dep](https://github.com/golang/dep) - Go Dependency Management Tool
* [GoReleaser](https://github.com/goreleaser/goreleaser) - Go Automated Binaries Build
* [Make](https://www.gnu.org/software/make/) - GNU Make Automated Execution
* [Docker](https://www.docker.com) - Application Containerization

## Authors

* **Dimas Restu Hidayanto** - *Initial Work* - [DimasKiddo](https://github.com/dimaskiddo)

See also the list of [contributors](https://github.com/dimaskiddo/codebase-go-rest/contributors) who participated in this project.
