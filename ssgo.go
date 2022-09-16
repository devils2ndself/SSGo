package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {

	var (
		input string = ""
		output string = "./dist"
		help bool = false
		version bool = false
	)

	flag.StringVarP(&input, "input", "i", "", "path to .txt file or folder to be turned into HTML")
	flag.StringVarP(&output, "output", "o", output, "path to end result")
	flag.BoolVarP(&help, "help", "h", false, "display detailed help")
	flag.BoolVarP(&version, "version", "v", false, "display current version")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Usage: ssgo -i [path]\nHelp: ssgo [-h | --help]")
		os.Exit(1)
	} else {
		if input != "" {
			processInput(input, output)
		} else if help {
			fmt.Println("Basic usage: ssgo [flag]")
			fmt.Println("Flags:")
			fmt.Println("\t[-i | --input] [path] - ")
			fmt.Println("\t[-v | --version] - Display installed version of SSGo")
			fmt.Println("\t[-h | --help] - Display detailed help message")

		} else if version {
			fmt.Println("SSGo version 0.1")
		} else {
			fmt.Println("Unrecognized flag. Use 'ssgo [-h | --help]' for available commands.")
		}
	}

}

func processInput(input string, output string) {
	fmt.Println("Input: " + input)
	fmt.Println("Input: " + output)
}