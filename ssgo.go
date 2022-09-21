package main

import (
	"fmt"
	"os"

	"github.com/devils2ndself/SSGo/utils"
	flag "github.com/spf13/pflag"
)

const version string = "0.1"

func main() {

	var (
		input string = ""
		output string = "dist"
		displayHelp bool = false
		displayVersion bool = false
	)

	flag.StringVarP(&input, "input", "i", "", "Path to a .txt file OR a folder containing .txt files to be turned into HTML")
	flag.StringVarP(&output, "output", "o", output, "Optional. Additionaly changes the output path of generated HTML")
	flag.BoolVarP(&displayHelp, "help", "h", false, "Display detailed help message")
	flag.BoolVarP(&displayVersion, "version", "v", false, "Display installed version of SSGo")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Usage: ssgo -i [path] -o [out path]\nHelp: ssgo [-h | --help]")
		os.Exit(1)
	} else {
		if input != "" {
			utils.ProcessInput(input, output)
		} else if displayHelp {
			fmt.Println("Basic usage: ssgo [flag] [value]")
			fmt.Println("Flags:")
			fmt.Println("\t[-i | --input] [path]      \t- Path to a .txt file OR a folder containing .txt files to be turned into HTML")
			fmt.Println("\t                           \t  For paths with spaces, please enclose them into double quotation marks, e.g. \"some path\"")
			fmt.Println("\t                           \t  By default, places generated HTML into ./dist.")
			fmt.Println("\t[-o | --output] [out path] \t- Optional. Additionaly changes the output path of generated HTML")
			fmt.Println("\n\t[-v | --version]           \t- Display installed version of SSGo")
			fmt.Println("\t[-h | --help]              \t- Display detailed help message")

		} else if displayVersion {
			fmt.Println("SSGo version " + version)
		} else {
			fmt.Println("Invalid call. Use 'ssgo [-h | --help]' for available commands.")
		}
	}

}