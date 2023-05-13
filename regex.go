package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"
)

type RegexOptions struct {
	TimeExpr       *regexp.Regexp
	TimeLayout     string
	DurationLayout DurationLayout
}

func RegexHandle(options RegexOptions, stdin *bufio.Reader, stdout *bufio.Writer) error {
	timeExpr := options.TimeExpr
	timeLayout := options.TimeLayout
	durationLayout := options.DurationLayout

	lineNumber := 0
	lineBreak := fmt.Sprintln("")

	var lastValue *time.Time
	var deltaText string

	for {
		line, _, err := stdin.ReadLine()
		if err == io.EOF {
			break
		}

		lineNumber += 1

		matches := timeExpr.FindSubmatch(line)
		if matches == nil || len(matches) < 2 {
			fmt.Fprintf(
				os.Stderr, "line[%d] does not match expression\n",
				lineNumber-1,
			)
			return fmt.Errorf("unable to find match on line %d", lineNumber)
		}

		text := string(matches[1])
		value, err := time.Parse(timeLayout, text)
		if err != nil {
			fmt.Fprintf(
				os.Stderr, "line[%d] has an invalid time format (expected %s): %d\n",
				lineNumber-1,
				text,
				err,
			)
			continue
		}

		if lastValue == nil {
			deltaText = "0"
		} else {
			delta := value.Sub(*lastValue)

			deltaText = FormatDuration(delta, durationLayout)
		}

		lastValue = &value

		_, err = stdout.WriteString(deltaText)
		if err != nil {
			return err
		}

		_, err = stdout.WriteString("\t")
		if err != nil {
			return err
		}

		_, err = stdout.Write(line)
		if err != nil {
			return err
		}

		_, err = stdout.WriteString(lineBreak)
		if err != nil {
			return err
		}
	}

	return nil
}

func RegexMain(stdin *bufio.Reader, stdout *bufio.Writer) error {
	timeExpr, err := regexp.Compile(`\[(.+)\]`)
	if err != nil {
		return err
	}

	// TODO parse options from os.Args
	options := RegexOptions{
		TimeExpr:       timeExpr,
		TimeLayout:     time.RFC3339,
		DurationLayout: Seconds,
	}
	err = RegexHandle(options, stdin, stdout)
	return err
}
