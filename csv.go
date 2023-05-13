package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

type CsvOptions struct {
	TimeColumn     int
	TimeLayout     string
	DurationLayout DurationLayout
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

func CsvMain(stdin *bufio.Reader, stdout *bufio.Writer) error {
	// TODO parse options from os.Args
	options := CsvOptions{
		TimeColumn:     0,
		TimeLayout:     time.RFC3339,
		DurationLayout: Seconds,
	}
	err := CsvHandle(options, stdin, stdout)
	return err
}
