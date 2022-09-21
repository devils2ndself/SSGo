package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	beforeTitleHTML string = "<!doctype html>\n<html lang=\"en\">\n<head>\n<meta charset=\"utf-8\">\n<link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/water.css@2/out/water.css\">\n<title>"
	afterTitleHTML string = "</title>\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n</head>\n<body>\n"
	closingHTML string = "\n</body>\n</html>"
)

func ProcessInput(input string, output string) {

	fmt.Println("Input: " + input)

	fileInfo, err := os.Stat(input)

	if os.IsNotExist(err) {
		fmt.Println("Error! No such file or directory!")
		os.Exit(1)
	}

	output = strings.TrimSpace(output)
	fmt.Println("Output: " + output)

	if output == "" {
		fmt.Println("Error! Output directory string is empty!")
		os.Exit(1)
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

	if fileInfo.IsDir() {

		fmt.Println("Iterating through directory...")

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
					fmt.Printf("File: %s\n", path)
					GenerateHTML(path, output, name)
				}
			}
			return nil
		})

	} else {
		var (
			basename string = fileInfo.Name()
			ext      string = filepath.Ext(basename)
			name     string = strings.TrimSuffix(basename, ext)
		)

		if ext != ".txt" {
			fmt.Println("Error! Input file is not a .txt")
			os.Exit(1)
		}

		GenerateHTML(input, output, name)
	}
}

func GenerateHTML(input string, output string, name string) {

	newFile, err := os.Create(output + "/" + name + ".html")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File created at " + output + "/" + name + ".html")
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
		fmt.Println("Title:", title)
		firstLine = "<h1>" + title + "</h1>\n<p>"
	} else {
		fmt.Println("No title found")
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

	log.Println("Writing buffer to file...")
	writer.Flush()
}