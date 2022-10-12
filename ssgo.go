package main

import (
	"fmt"
	"os"

	utils "github.com/devils2ndself/SSGo/utils"
	flag "github.com/spf13/pflag"
)

const version string = "0.4"

func main() {
	var (
		input string = ""
		output string = utils.DefaultOutput
		displayHelp bool = false
		displayVersion bool = false
		config string = ""
	)

	// Flag initialization
	flag.StringVarP(&input, "input", "i", "", utils.InputHelpMessage)
	flag.StringVarP(&output, "output", "o", utils.DefaultOutput, utils.OutputHelpMessage)
	flag.BoolVarP(&displayHelp, "help", "h", false, utils.HelpHelpMessage)
	flag.BoolVarP(&displayVersion, "version", "v", false, utils.VersionHelpMessage)
	flag.StringVarP(&config, "config", "c", "", utils.ConfigHelpMessage)

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Usage: ssgo -i [path] -o [out path]\nHelp: ssgo [-h | --help]")
		os.Exit(1)
	} else {
		if config != "" {
			utils.ProcessConfig(config)
		}else if input != "" {
			utils.ProcessInput(input, output)
		} else if displayHelp {
			utils.PrintHelp()
		} else if displayVersion {
			fmt.Println("SSGo version " + version)
		} else {
			fmt.Println("Invalid call. Use 'ssgo [-h | --help]' for available commands.")
		}
	}
}