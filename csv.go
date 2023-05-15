package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type CsvOptions struct {
	TimeColumn     int
	TimeLayout     string
	DurationLayout DurationLayout
}

func CsvOptionsParse(args []string) (CsvOptions, error) {
	options := CsvOptions{
		TimeColumn:     0,
		TimeLayout:     time.RFC3339,
		DurationLayout: Seconds,
	}
	for i := 0; i < len(args); i += 2 {
		if i+1 == len(args) {
			return options, errors.New("expected a positive integer, ex: --time-column 3")
		}
		switch args[i] {
		case "--time-column":
			column, err := strconv.Atoi(args[i+1])
			if err != nil {
				return options, err
			}
			if column < 0 {
				return options, errors.New("expected a positive integer, ex: --time-column 3")
			}
			options.TimeColumn = column
		case "--time-layout":
			options.TimeLayout = args[i+1]
		case "--duration-layout":
			layout, err := ParseDuration(args[i+1])
			if err != nil {
				return options, err
			}
			options.DurationLayout = layout
		default:
			return options, fmt.Errorf("invalid arg %s", args[i])
		}
	}
	return options, nil
}

func CsvHandle(options CsvOptions, stdin *bufio.Reader, stdout *bufio.Writer) error {
	csvReader := csv.NewReader(stdin)
	csvWriter := csv.NewWriter(stdout)
	defer csvWriter.Flush()

	rowNumber := 0

	index := options.TimeColumn
	timeLayout := options.TimeLayout
	durationLayout := options.DurationLayout

	var lastValue *time.Time

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		rowNumber += 1

		if index > len(record) {
			fmt.Fprintf(os.Stderr, "row[%d] missing columns (%d)\n", rowNumber-1, len(record))
			continue
		}

		value, err := time.Parse(timeLayout, record[index])
		if err != nil {
			fmt.Fprintf(
				os.Stderr, "row[%d][%d] has an invalid time format (expected %s): %d\n",
				rowNumber-1,
				index,
				record[index],
				err,
			)
			continue
		}

		record = append(record[:index+1], record[index:]...)

		if lastValue == nil {
			record[index] = "0"
		} else {
			delta := value.Sub(*lastValue)
			record[index] = FormatDuration(delta, durationLayout)
		}

		lastValue = &value

		err = csvWriter.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
