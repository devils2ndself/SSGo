package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
)

const (
	beforeTitleHTML string = "<!doctype html>\n<html lang=\"en\">\n<head>\n<meta charset=\"utf-8\">\n<link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/water.css@2/out/water.css\">\n<title>"
	afterTitleHTML  string = "</title>\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n</head>\n<body>\n"
	closingHTML     string = "</body>\n</html>"
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

		content := ParseMarkdown(fileBytes)

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
		fileBytes, err := os.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}
		content, title := ParseText(string(fileBytes))
		if title == "" {
			title = name
		}

		_, writeErr := newFile.WriteString(beforeTitleHTML + title + afterTitleHTML)
		if writeErr != nil {
			log.Fatal("Error writing to file! ", writeErr)
		}
		_, writeErr = newFile.WriteString(content)
		if writeErr != nil {
			log.Fatal("Error writing to file! ", writeErr)
		}
		_, writeErr = newFile.WriteString(closingHTML)
		if writeErr != nil {
			log.Fatal("Error writing to file! ", writeErr)
		}

	}
}

// Reads string slice of lines and returns HTML and title of the
func ParseText(fileString string) (string, string) {
	lines := strings.Split(strings.ReplaceAll(fileString, "\r\n", "\n"), "\n")

	var content string = ""
	var titleExists bool = true
	var title string = ""
	var firstParagraphLine = 0

	firstLine := lines[0]
	if len(firstLine) == 0 || len(lines) < 3 || len(lines[1]) != 0 || len(lines[2]) != 0 {
		titleExists = false
	}
	if titleExists {
		title = firstLine
	}

	// Write title or restart Scanner
	if titleExists {
		if debug {
			fmt.Println("Title:", title)
		}
		content += "<h1>" + title + "</h1>\n\n"
		firstParagraphLine = 3
	}

	paragraphOpen := false
	paragraphDelimiterFound := true

	// Scan line by line and append to html buffer
	for i := firstParagraphLine; i < len(lines); i++ {
		text := strings.TrimSpace(lines[i])
		if text != "" {
			// Write closing </p> tag if delimiter was found and opening <p> was previously written
			if paragraphOpen && paragraphDelimiterFound {
				content += "</p>\n\n"
				paragraphOpen = false
			}

			// If last read line was a paragraph delimiter, write opening <p> first
			if !paragraphOpen && paragraphDelimiterFound {
				text = "<p>" + text
				paragraphOpen = true
				paragraphDelimiterFound = false
			}
			content += text + " "

		} else {
			paragraphDelimiterFound = true
		}
	}

	if paragraphOpen {
		content += "</p>\n"
	}

	return content, title
}

func ParseMarkdown(fileBytes []byte) []byte {
	fileBytes = markdown.NormalizeNewlines(fileBytes)
	return markdown.ToHTML(fileBytes, nil, nil)
}
