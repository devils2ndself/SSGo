# Contributing

Pull requests are welcome. 

For major changes, please open an issue first to discuss what you would like to change. 

## Installation

In order to build the application, you first need to [get Go CLI](https://go.dev/doc/install). 

The project was built using Go version `go1.19.1 windows/amd64`. If you are on a different version and experiencing any issues, please submit an issue regarding that.

To build a binary or executable, run:

```
git clone https://github.com/devils2ndself/SSGo.git
cd SSGO
go build ssgo.go 
```
Or `go install` to install globally.

## Formatting

Anytime you perform changes that you want to contribute, please make sure the code is formatted appropriately.

> Use `gofmt -s -w .` in order to format all code according to Go standards.

## Linting

We use [GoLangCI-Lint](https://golangci-lint.run/) in order to lint SSGo. You can see installation guide [here](https://golangci-lint.run/usage/install/).

Linting is a part of a successful PR, so please make sure your code passes it.

> Use `golangci-lint run` in order to lint-check all code in the directory.

## IDE Integrations

### VS Code

It is recommended to install `golang.go` extension. This extension allows for automatic linting and on-save formatting.