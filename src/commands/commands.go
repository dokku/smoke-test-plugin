package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	columnize "github.com/ryanuber/columnize"
)

const (
	helpHeader = `Usage: dokku repo[:COMMAND]

Runs commands that interact with the app's repo

Additional commands:`

	helpContent = `
    smoke-test-plugin:test, prints test message
`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "smoke-test-plugin:help":
		usage()
	case "help":
		fmt.Print(helpContent)
	case "smoke-test-plugin:test":
		fmt.Println("triggered smoke-test-plugin from: commands")
	default:
		dokkuNotImplementExitCode, err := strconv.Atoi(os.Getenv("DOKKU_NOT_IMPLEMENTED_EXIT"))
		if err != nil {
			fmt.Println("failed to retrieve DOKKU_NOT_IMPLEMENTED_EXIT environment variable")
			dokkuNotImplementExitCode = 10
		}
		os.Exit(dokkuNotImplementExitCode)
	}
}

func usage() {
	config := columnize.DefaultConfig()
	config.Delim = ","
	config.Prefix = "\t"
	config.Empty = ""
	content := strings.Split(helpContent, "\n")[1:]
	fmt.Println(helpHeader)
	fmt.Println(columnize.Format(content, config))
}
