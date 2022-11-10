package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	path string ""
	name string ""
}

// Takes path to .json file, reads it, and calls ProcessInput using contained options
func ProcessConfig(config string) {

	// Read config .json file
	configFile, err := os.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}

	// Assign json values to options struct
	options := struct {
		Input  string `json:"input"`
		Output string `json:"output"`
	}{
		Output: "dist",
	}

	jsonErr := json.Unmarshal(configFile, &options)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// Call ProcessInput using config file options
	ProcessInput(options.Input, options.Output)
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
		fmt.Println("Looking for .txt / .md files in the directory...")

		// Walks through all elements in the directory
		walkErr := filepath.Walk(input, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			// If file
			if !info.IsDir() {
				name, ext := GetNameAndExt(info.Name())

				// If .txt, add file to files slice
				if _, exists := AcceptedInputFileTypes[ext]; exists {
					fmt.Printf("\tFile: %s\n", path)
					var f File = File{path: path, name: name}
					files = append(files, f)
				}
			}
			return nil
		})
		if walkErr != nil {
			log.Fatal(walkErr)
		}

		// If there is at least 1 .txt file
		if len(files) != 0 {
			// Prepare output directory
			prepareOutput(output)
			// Generate .html for each .txt file in the input directory
			for i := 0; i < len(files); i++ {
				GenerateHTML(files[i].path, output, files[i].name)
			}
		} else {
			log.Fatal("No .txt / .md files in the directory!")
		}

	} else {
		//If input is a file
		name, ext := GetNameAndExt(fileInfo.Name())

		if _, exists := AcceptedInputFileTypes[ext]; !exists {
			log.Fatal("Error! Input file is not a .txt or .md file")
		}

		// Prepare output directory
		prepareOutput(output)

		GenerateHTML(input, output, name)
	}
	fmt.Println("Done! Check '" + output + "' directory to see generated HTML.")
}

// Returns split name and extension of a filename
func GetNameAndExt(basename string) (string, string) {
	var (
		ext  string = filepath.Ext(basename)
		name string = strings.TrimSuffix(basename, ext)
	)
	return name, ext
}

// Checks output string and make/clear directory
func prepareOutput(output string) {

	output = strings.TrimSpace(output)
	if output == "" {
		log.Fatal("Error! Output directory string is empty!")
	}

	if debug {
		fmt.Println("Output folder: " + output)
	}

	out, outerr := os.Stat(output)

	if os.IsNotExist(outerr) {
		// Output path doesn't exist, create a new directory
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
