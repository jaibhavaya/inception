package main

import (
	"fmt"
	"os"
	"strings"
)

var supportedCommands = [4]string{
	"list-projects",
	"read-project",
	"render-pipeline",
	"self-test",
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you dun f!@#ed up")
		usage()
		os.Exit(1)
	}
	command := os.Args[1]
	if !isSupportedCommand(command) {
		fmt.Println(command, "is an unsupported command, please use one of the following:", strings.Join(supportedCommands[:], ", "))
		os.Exit(2)
	}
	fmt.Println("this command is goooood: ", command)
}

func isSupportedCommand(command string) bool {
	found := false
	for _, c := range supportedCommands {
		if c == command {
			found = true
			break
		}
	}
	return found
}

func usage() {
	fmt.Println("usage: \nscript command")
}
