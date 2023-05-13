package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type JsonlOptions struct {
	TimeName       string
	TimeLayout     string
	DurationName   string
	DurationLayout DurationLayout
}

func JsonlHandle(options JsonlOptions, stdin *bufio.Reader, stdout *bufio.Writer) error {
	jsonReader := json.NewDecoder(stdin)
	jsonWriter := json.NewEncoder(stdout)

	lineNumber := 0

	timeName := options.TimeName
	timeLayout := options.TimeLayout
	durationName := options.DurationName
	durationLayout := options.DurationLayout

	var lastValue *time.Time

	for {
		o := map[string]interface{}{}
		err := jsonReader.Decode(&o)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		lineNumber += 1

		val, ok := o[timeName]
		if !ok {
			fmt.Fprintf(
				os.Stderr,
				"line[%d][%s] is undefined\n",
				lineNumber-1,
				timeName,
			)
			continue
		}

		textValue := fmt.Sprintf("%v", val)
		value, err := time.Parse(timeLayout, textValue)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"line[%d][%s] has an invalid time format (expected %s): %d\n",
				lineNumber-1,
				timeName,
				value,
				err,
			)
			continue
		}

		if lastValue == nil {
			o[durationName] = "0"
		} else {
			delta := value.Sub(*lastValue)
			o[durationName] = FormatDuration(delta, durationLayout)
		}

		lastValue = &value

		err = jsonWriter.Encode(o)
		if err != nil {
			return err
		}
	}

	return nil
}

func JsonlMain(stdin *bufio.Reader, stdout *bufio.Writer) error {
	// TODO parse options from os.Args
	options := JsonlOptions{
		TimeName:       "time",
		TimeLayout:     time.RFC3339,
		DurationLayout: Seconds,
		DurationName:   "duration",
	}
	err := JsonlHandle(options, stdin, stdout)
	return err
}
