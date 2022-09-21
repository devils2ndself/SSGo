package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const debug bool = false

type File struct {
	path    string ""
	name	string ""
}

const (
	beforeTitleHTML string = "<!doctype html>\n<html lang=\"en\">\n<head>\n<meta charset=\"utf-8\">\n<link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/water.css@2/out/water.css\">\n<title>"
	afterTitleHTML string = "</title>\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n</head>\n<body>\n"
	closingHTML string = "\n</body>\n</html>"
)

const (
	InputHelpMessage string = "Path to a .txt file OR a folder containing .txt files to be turned into HTML"
	OutputHelpMessage string = "Optional. Additionaly changes the output path of generated HTML"
	HelpHelpMessage string = "Display detailed help message"
	VersionHelpMessage string = "Display installed version of SSGo"
)

func PrintHelp() {

	fmt.Println("Basic usage: ssgo [flag] [value]")
	fmt.Println("Flags:")
	fmt.Println("\t[-i | --input] [path]      \t- " + InputHelpMessage)
	fmt.Println("\t                           \t  For paths with spaces, please enclose them into double quotation marks, e.g. \"some path\"")
	fmt.Println("\t                           \t  By default, places generated HTML into ./dist.")
	fmt.Println("\t[-o | --output] [out path] \t- " + OutputHelpMessage)
	fmt.Println("\n\t[-v | --version]           \t- " + HelpHelpMessage)
	fmt.Println("\t[-h | --help]              \t- " + VersionHelpMessage)
}


// Checks output string and make/clear directory
func prepareOutput(output string) {

	output = strings.TrimSpace(output)
	if debug {
		fmt.Println("Output folder: " + output)
	}

	if output == "" {
		log.Fatal("Error! Output directory string is empty!")
	}

	out, outerr := os.Stat(output)

	if os.IsNotExist(outerr) {
		// Output path doesn't exist
		if err := os.Mkdir(output, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	} else if !out.IsDir() {
		// If path exists but is not a directory
		log.Fatal("Error! Output path exists, but is not a directory!")
	} else if output == "dist" {
		// Limiting to only dist folder in order to prevent unwwanted deletions
		// If directory exists, delete it and re-create to get rid of the previous contents
		fmt.Println("Removing previous compilations from the output folder...")
		err := os.RemoveAll(output)
		if err != nil {
			log.Fatal(err)
		}
		if err := os.Mkdir(output, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}


// Takes input path, validates single .txt file OR folder and checks all files in the folder
func ProcessInput(input string, output string) {

	if debug {
		fmt.Println("Input: " + input)
	}

	// Reads stats about input path
	fileInfo, err := os.Stat(input)

	if os.IsNotExist(err) {
		log.Fatal("Error! No such file or directory!")
	}

	// If input is directory
	if fileInfo.IsDir() {

		var files []File

		fmt.Println("Looking for .txt files in the directory...")

		// Walks through all elements in the directory
		filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			// If file
			if !info.IsDir() {
				var (
					basename string = info.Name()
					ext      string = filepath.Ext(basename)
					name     string = strings.TrimSuffix(basename, ext)
				)

				// If .txt, add file to files slice
				if ext == ".txt" {
					fmt.Printf("\tFile: %s\n", path)
					var f File = File{path: path, name: name}
					files = append(files, f)
				}
			}
			return nil
		})

		// If there is at least 1 .txt file
		if len(files) != 0 {
			// Prepare output directory
			prepareOutput(output)
			// Generate .html for each .txt file in the input directory
			for i := 0; i < len(files); i++ {
				GenerateHTML(files[i].path, output, files[i].name)
			}
		} else {
			log.Fatal("No .txt files in the directory!")
		}

	} else {
		//If input is a file
		var (
			basename string = fileInfo.Name()
			ext      string = filepath.Ext(basename)
			name     string = strings.TrimSuffix(basename, ext)
		)

		if ext != ".txt" && ext != ".md" {
			log.Fatal("Error! Input file is not a .txt or .md file")
		}

		// Prepare output directory
		prepareOutput(output)

		GenerateHTML(input, output, name)
	}
	fmt.Println("Done! Check '" + output + "' directory to see generated HTML.")
}


// Takes path to .txt file as an input, reads it, and creates name.html in output folder
func GenerateHTML(input string, output string, name string) {

	// Create new empty .html file
	newFile, err := os.Create(output + "/" + name + ".html")
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		log.Println("File created at " + output + "/" + name + ".html")
	}
	defer newFile.Close()

	txtFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer txtFile.Close()

	scanner := bufio.NewScanner(txtFile)

	var titleExists bool = true
	var title string = name

	// Check title
	scanner.Scan()
	firstLine := scanner.Text()
	if len(firstLine) != 0 {
		for i := 0; i < 2; i++ {
			if !scanner.Scan() {
				break
			}
			if len(scanner.Text()) != 0 {
				titleExists = false
			}
		}
		if titleExists {
			title = firstLine
		}
	}

	// Writing HTML head to buffer
	writer := bufio.NewWriter(newFile)
	_, werr := writer.WriteString(beforeTitleHTML + title + afterTitleHTML)
	if werr != nil {
		log.Fatal("Error writing to buffer!")
	}

	// Determine first line to be written in the body - title or regular <p>
	if titleExists {
		if debug {
			fmt.Println("Title:", title)
		}
		firstLine = "<h1>" + title + "</h1>\n<p>"
	} else {
		if debug {
			fmt.Println("No title found")
		}
		// If title does not exist, restart Scanner by opening another instance of the same file
		// This way, it will read again from the first line
		reTxtFile, err := os.Open(input)
		if err != nil {
			log.Fatal(err)
		}
		defer reTxtFile.Close()

		scanner = bufio.NewScanner(reTxtFile)
		firstLine = "<p>"
	}

	// Write first line - title in <h1> or empty opening <p>
	_, werr = writer.WriteString(firstLine)
	if werr != nil {
		log.Fatal("Error writing to buffer!")
	}

	// Scan line by line and append to html buffer
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			// Line with content = append to the current <p>

			// Parse markdown file
			markdownValid := false

			if filepath.Ext(input) == ".md" {
				if prefix, validPrefix := CheckMarkdownPrefix(text); validPrefix {
					_, werr = writer.WriteString(GenerateMarkdownHtml(prefix, text))
					if werr != nil {
						log.Fatal("Error writing to new file!")
					} else {
						markdownValid = true
					}
				}
			}

			// Write the text as is if markdown format was invalid
			if !markdownValid {
				_, werr = writer.WriteString(text)
				if werr != nil {
					log.Fatal("Error writing to new file!")
				}
			}
		} else {
			// Empty line = close </p> and open next one
			_, werr = writer.WriteString("</p>\n<p>")
			if werr != nil {
				log.Fatal("Error writing to new file!")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Writing closing HTML tags to buffer
	_, w2err := writer.WriteString("</p>" + closingHTML)
	if w2err != nil {
		log.Fatal("Error writing to new file!")
	}

	// Writing buffer to the file
	if debug {
		log.Println("Writing buffer to file...")
	}
	writer.Flush()
}

func CheckMarkdownPrefix(text string) (string, bool) {
	acceptedPrefixes := [2]string{"# ", "## "}

	for _, prefix := range acceptedPrefixes {
		if strings.HasPrefix(text, prefix) {
			return prefix, true
		}
	}

	return "", false
}

func GenerateMarkdownHtml(prefix string, text string) string {
	prefixesHtmlFormatString := map[string]string{
		"# ":  "<h1>%s</h1>",
		"## ": "<h2>%s</h2>",
	}

	if formatString, found := prefixesHtmlFormatString[prefix]; found {
		return fmt.Sprintf(formatString, text)
	}

	return text
}
