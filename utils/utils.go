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

func ProcessInput(input string, output string) {

	if debug {
		fmt.Println("Input: " + input)
	}

	fileInfo, err := os.Stat(input)

	if os.IsNotExist(err) {
		log.Fatal("Error! No such file or directory!")
	}

	if fileInfo.IsDir() {

		var files []File
				
		fmt.Println("Looking for .txt files in the directory...")

		filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			if !info.IsDir() {
				var (
					basename string = info.Name()
					ext      string = filepath.Ext(basename)
					name     string = strings.TrimSuffix(basename, ext)
				)

				if ext == ".txt" {
					fmt.Printf("\tFile: %s\n", path)
					var f File = File{path: path, name: name}
					files = append(files, f)
				}
			}
			return nil
		})

		if len(files) != 0 {
			prepareOutput(output)
			for i := 0; i < len(files); i++ {
				GenerateHTML(files[i].path, output, files[i].name)
			}
		} else {
			log.Fatal("No .txt files in the directory!")
		}


	} else {
		var (
			basename string = fileInfo.Name()
			ext      string = filepath.Ext(basename)
			name     string = strings.TrimSuffix(basename, ext)
		)

		if ext != ".txt" {
			log.Fatal("Error! Input file is not a .txt")
		}

		prepareOutput(output)

		GenerateHTML(input, output, name)
	}
	fmt.Println("Done! Check '" + output + "' directory to see generated HTML.")
}

func GenerateHTML(input string, output string, name string) {

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

	writer := bufio.NewWriter(newFile)
	_, werr := writer.WriteString(beforeTitleHTML + title + afterTitleHTML)
	if werr != nil {
		log.Fatal("Error writing to buffer!")
	}

	if titleExists {
		if debug {
			fmt.Println("Title:", title)
		}
		firstLine = "<h1>" + title + "</h1>\n<p>"
	} else {
		if debug {
			fmt.Println("No title found")
		}
		reTxtFile, err := os.Open(input)
		if err != nil {
			log.Fatal(err)
		}
		defer reTxtFile.Close()

		scanner = bufio.NewScanner(reTxtFile)
		firstLine = "<p>"
	}

	_, werr = writer.WriteString(firstLine)
	if werr != nil {
		log.Fatal("Error writing to buffer!")
	}

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			_, werr = writer.WriteString(text)
			if werr != nil {
				log.Fatal("Error writing to new file!")
			}
		} else {
			_, werr = writer.WriteString("</p>\n<p>")
			if werr != nil {
				log.Fatal("Error writing to new file!")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	_, w2err := writer.WriteString("</p>" + closingHTML)
	if w2err != nil {
		log.Fatal("Error writing to new file!")
	}

	if debug {
		log.Println("Writing buffer to file...")
	}
	writer.Flush()
}