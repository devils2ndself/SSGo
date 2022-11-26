# SSGo
A simple Static Site Generator built using Go language. 

Lets you generate HTML files from either a single `.txt` file or multiple `.txt` files in a given directory.

Preview of generated HTML is available at https://devils2ndself.github.io/SSGo.

## Installation

In order to install SSGo, please first [install Go](https://go.dev/dl/).

Use `go install github.com/devils2ndself/SSGo@v1.0.3` in order to install SSGo to your local machine.

## Usage

- `ssgo --input [in] --output [out]` - Generate HTML from .txt or .md file at `in` path (can be a single .txt or .md file, or directory) to `out` path.  
`--output` is optional, the default out is `dist` folder in the current directory

- `ssgo --config [cfg]` - Uses the options from .json configuration file at `cfg` path to specify options. 

- `ssgo --version` - Display installed version of SSGo

- `ssgo --help` - Display detailed help message

Please note, that if built locally, some users will need to use `./ssgo` instead of `ssgo`.

Also, please be aware that shorthand flags, like `-i`, take as the argument everything that follows that character, i.e. `-info` will be understood as `-i nfo`. It's not me, it's the standards :/

## Features

- Generates HTML files for each .txt or .md file in `input`.

- Encloses every paragraph of text separated by a blank line in `<p>` tag.

- If the first line of the .txt or .md file is followed by 2 empty lines, it will be used as a title. The `<title>` will be assigned to it and it will be enclosed into `<h1>` tag instead of regular `<p>` - _optional feature #1_.

- The name of the generated HTML files will be the same as the original .txt or .md files.

- Also provides a cool style to generate HTML - _optional feature #6_.

- Generated files go to `dist` folder unless any other path is specified with `--output` flag. If the new output path does not exist, it will create a new directory to accommodate - _optional feature #2_.

- If the output directory is `path`, each time new text is used for HTML generation, the directory will be wiped completely. Be careful, don't lose valuable files there!

- If `--output` is specified, the directory at path will not be erased like with `dist` folder. Just in case someone specifies `--output C:\...`

- If `--config` is specified, SSG options can be supplied through a .json configuration file instead of through the command-line.

### Markdown Features
These feaures are supported for files with an extension of '.md'

- Lines beginning with "# ", "## " will be wrapped within an h1 and h2 tags respectively: ``# This is heading 1`` becomes ``<h1>This is heading 1</h1> `` 

- Lines that contain only `-`, `*`, or `_` in quantity of *3 or more* will be a horizontal rule: `---` becomes `<hr>`

- Any text that is contained within \` characters will be wrapped within `<code>` tag: `some code` becomes `<code>some code</code>`

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[MIT](https://choosealicense.com/licenses/mit/)
