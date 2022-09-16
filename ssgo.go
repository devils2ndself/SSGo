package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

func main() {

	var (
		input string = ""
		output string = "dist"
		help bool = false
		version bool = false
	)

	flag.StringVarP(&input, "input", "i", "", "path to .txt file or folder to be turned into HTML")
	flag.StringVarP(&output, "output", "o", output, "path to end result")
	flag.BoolVarP(&help, "help", "h", false, "display detailed help")
	flag.BoolVarP(&version, "version", "v", false, "display current version")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Usage: ssgo -i [path] -o [out path]\nHelp: ssgo [-h | --help]")
		os.Exit(1)
	} else {
		if input != "" {
			processInput(input, output)
		} else if help {
			fmt.Println("Basic usage: ssgo [flag] [value]")
			fmt.Println("Flags:")
			fmt.Println("\t[-i | --input] [path]      \t- Path to a .txt file OR a folder containing .txt files to be turned into HTML")
			fmt.Println("\t                           \t  For paths with spaces, please enclose them into double quotation marks, e.g. \"some path\"")
			fmt.Println("\t                           \t  By default, places generated HTML into ./dist.")
			fmt.Println("\t[-o | --output] [out path] \t- Optional. Additionaly changes the output path of generated HTML")
			fmt.Println("\n\t[-v | --version]           \t- Display installed version of SSGo")
			fmt.Println("\t[-h | --help]              \t- Display detailed help message")

		} else if version {
			fmt.Println("SSGo version 0.1")
		} else {
			fmt.Println("Invalid call. Use 'ssgo [-h | --help]' for available commands.")
		}
	}

}

func processInput(input string, output string) {
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
	} else {
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

		fmt.Println("Generation from directory not implemented yet...")
	
	} else {

		var (
			basename string = fileInfo.Name()
			ext string = filepath.Ext(basename)
			name string = strings.TrimSuffix(basename, ext)
		)

		if ext != ".txt" {
			fmt.Println("Error! Input file is not a .txt")
			os.Exit(1)
		}

		fmt.Println("Title: " + name)

		newFile, err := os.Create(output + "/" + name + ".html")
		
		if err != nil {
			log.Fatal(err)
		}
		log.Println(newFile)
		newFile.Close()
	}
}