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
	execPath := strings.Split(os.Args[0], "/")
	executable := execPath[len(execPath)-1]
	switch executable {
	case "commands":
		commands()
	case "pre-deploy":
		printMsg("pre-deploy")
	}
}

func commands() {
	flag.Usage = usage
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "smoke-test-plugin:help":
		usage()
	case "help":
		fmt.Print(helpContent)
	case "smoke-test-plugin:test":
		printMsg("commands")
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

func printMsg(msg string) {
	fmt.Printf("triggered smoke-test-plugin from: %s\n", msg)
}
