package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
)

const (
	beforeTitleHTML string = "<!doctype html>\n<html lang=\"en\">\n<head>\n<meta charset=\"utf-8\">\n<link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/water.css@2/out/water.css\">\n<title>"
	afterTitleHTML string = "</title>\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n</head>\n<body>\n"
	closingHTML string = "\n</body>\n</html>"
)

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

	if filepath.Ext(input) == ".md" {
		// If markdown file, read file into bytes and use outside module
		fileBytes, fileReadErr := os.ReadFile(input)
		if fileReadErr != nil {
			log.Fatal("Error reading the file! ", fileReadErr)
		}

		fileBytes = markdown.NormalizeNewlines(fileBytes)
		content := markdown.ToHTML(fileBytes, nil, nil)

		_, writeErr := newFile.WriteString(beforeTitleHTML + name + afterTitleHTML)
		if writeErr != nil {
			log.Fatal("Error writing to file! ", writeErr)
		}
		_, writeErr = newFile.Write(content)
		if writeErr != nil {
			log.Fatal("Error writing to file! ", writeErr)
		}
		_, writeErr = newFile.WriteString(closingHTML)
		if writeErr != nil {
			log.Fatal("Error writing to file! ", writeErr)
		}

	} else {
		// 
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

		// Write title or restart Scanner
		if titleExists {
			if debug {
				fmt.Println("Title:", title)
			}
			firstLine = "<h1>" + title + "</h1>\n"
			_, werr = writer.WriteString(firstLine)
			if werr != nil {
				log.Fatal("Error writing to buffer!")
			}
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
		}

		paragraphOpen := false
		paragraphDelimiterFound := true

		// Scan line by line and append to html buffer
		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())
			if text != "" {
				// Write closing </p> tag if delimiter was found and opening <p> was previously written
				if paragraphOpen && paragraphDelimiterFound {
					_, werr = writer.WriteString("</p>\n")
					if werr != nil {
						log.Fatal("Error writing to new file!")
					}
					paragraphOpen = false
				}
				
				// If last read line was a paragraph delimiter, write opening <p> first
				if !paragraphOpen && paragraphDelimiterFound {
					text = "<p>" + text
					paragraphOpen = true
					paragraphDelimiterFound = false
				}
				_, werr = writer.WriteString(text)
				if werr != nil {
					log.Fatal("Error writing to new file!")
				}
				
			} else {
				paragraphDelimiterFound = true
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
}