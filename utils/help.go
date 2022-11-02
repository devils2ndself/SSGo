package utils

import "fmt"

const (
	InputHelpMessage   string = "Path to a .txt / .md file OR a folder containing .txt / .md files to be turned into HTML"
	OutputHelpMessage  string = "Optional. Additionaly changes the output path of generated HTML"
	HelpHelpMessage    string = "Display detailed help message"
	VersionHelpMessage string = "Display installed version of SSGo"
	ConfigHelpMessage  string = "Path to a .json file containing SSGo configuration options"
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
	fmt.Println("\t[-c | --config]              \t- " + ConfigHelpMessage)
}
