# SSGo
A simple Static Site Generator built using Go language.

## Set-up

In order to build the application, you first need to [get Go CLI](https://go.dev/doc/install).

To build a binary, or executable, run:   
`git clone https://github.com/devils2ndself/SSGo.git`  
`cd SSGO`  
`go build ssgo.go` _or_ `go install` to install globally 

## Usage

- `ssgo --input [in] --output [out[` - Generate HTML from .txt file at `in` path (can be a single .txt file or directory) to `out` path.  
`--output` is optional, default out is `dist` folder in current directory

- `ssgo --version` - Display installed version of SSGo

- `ssgo --help` - Display detailed help message

## Functionality

- Generates HTML files for each .txt file in `input`.

- Encloses every paragraph of text separated by a blank line in `<p>` tag.

- If the first line of the .txt file is followed by 2 empty lines, it will be used as a title. The `<title>` will be assigned to it and it will be enclosed into `<h1>` tag instead of regular `<p>`.

- The name of the generated HTML files will be the same as the original .txt files.

- Generated files go to `dist` folder unless any other path is specified with `--output` flag. If the new output path does not exist, it will create a new directory to accomodate.

- If output directory is `path`, each time new text is used for HTML generation, the directory will be wiped completely. Careful, don't lose valuable files there!

- If `--output` is specified, the directory at path will not be erased like with `dist` folder. Just in case someone specifies `--output C:\`...