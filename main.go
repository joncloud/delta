package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	var command string
	if len(args) > 0 {
		command = args[0]
	} else {
		command = "text"
	}

	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	var err error
	switch command {
	case "csv":
		err = CsvMain(stdin, stdout)
	case "jsonl":
		err = JsonlMain(stdin, stdout)
	case "regex":
		err = RegexMain(stdin, stdout)
	// TODO help
	default:
		err = fmt.Errorf("invalid command %s", command)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%d\n", err)
		os.Exit(1)
	}
}
