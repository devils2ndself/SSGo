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

## Testing

Anytime you contribute changes, please make sure all of the tests pass, and make sure to update tests when necessary.

>In order to run all tests, run the following in CLI:  
>`go test -v github.com/devils2ndself/SSGo/utils`

>If you would like to run any specific test, you can run (change `Test_ParseText_NoHeading` to desired test)  
>`go test -run ^Test_ParseText_NoHeading$ -v github.com/devils2ndself/SSGo/utils`  

### Coverage 

To see test coverage, first generate coverage file, then use the cover tool to generate HTML from the coverage file.  

```
go test -coverprofile cover.out github.com/devils2ndself/SSGo/utils
go tool cover -html=cover.out
```

Then you can open coverage in browser as HTML.

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