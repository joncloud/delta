package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
	"time"
)

func TestCsvOptionsParseEmptyArgs(t *testing.T) {
	args := []string{}

	options, err := CsvOptionsParse(args)
	if options.TimeColumn != 0 ||
		options.TimeLayout != time.RFC3339 ||
		options.DurationLayout != Seconds ||
		err != nil {
		t.Fatalf("CsvOptionsParse(%v) = %v, %v", args, options, err)
	}
}

func TestCsvOptionsParseOddArgsLength(t *testing.T) {
	args := []string{"--anything"}

	options, err := CsvOptionsParse(args)
	if err == nil {
		t.Fatalf("CsvOptionsParse(%v) = %v, %v", args, options, err)
	}
}

func TestCsvOptionsParseNonNumericTimeColumn(t *testing.T) {
	args := []string{"--time-column", "abc"}

	options, err := CsvOptionsParse(args)
	if err == nil {
		t.Fatalf("CsvOptionsParse(%v) = %v, %v", args, options, err)
	}
}

func TestCsvOptionsParseNegativeTimeColumn(t *testing.T) {
	args := []string{"--time-column", "-1"}

	options, err := CsvOptionsParse(args)
	if err == nil {
		t.Fatalf("CsvOptionsParse(%v) = %v, %v", args, options, err)
	}
}

func TestCsvOptionsParsePositiveTimeColumn(t *testing.T) {
	args := []string{"--time-column", "123"}

	options, err := CsvOptionsParse(args)
	if options.TimeColumn != 123 ||
		err != nil {
		t.Fatalf("CsvOptionsParse(%v) = %v, %v", args, options, err)
	}
}

func TestCsvOptionsParseTimeLayout(t *testing.T) {
	args := []string{"--time-layout", "abc"}

	options, err := CsvOptionsParse(args)
	if options.TimeLayout != "abc" ||
		err != nil {
		t.Fatalf("CsvOptionsParse(%v) = %v, %v", args, options, err)
	}
}

func TestCsvOptionsParseInvalidDurationLayout(t *testing.T) {
	args := []string{"--duration-layout", "foo"}

	options, err := CsvOptionsParse(args)
	if err == nil {
		t.Fatalf("CsvOptionsParse(%v) = %v, %v", args, options, err)
	}
}

func TestCsvOptionsParseValidDurationLayout(t *testing.T) {
	args := []string{"--duration-layout", string(Hours)}

	options, err := CsvOptionsParse(args)
	if options.DurationLayout != Hours ||
		err != nil {
		t.Fatalf("CsvOptionsParse(%v) = %v, %v", args, options, err)
	}
}

func TestCsvHandle(t *testing.T) {
	input, err := os.Open("test.csv")
	if err != nil {
		t.Fatal("Unable to open test.csv")
	}

	defer input.Close()

	var b bytes.Buffer

	var args []string
	options, err := CsvOptionsParse(args)
	if err != nil {
		t.Fatalf("Unable to read csv options: %v", err)
	}
	stdout := bufio.NewWriter(&b)
	err = CsvHandle(options, bufio.NewReader(input), stdout)

	if err != nil {
		t.Fatalf("Unable to handle csv: %v", err)
	}

	stdout.Flush()

	file, err := os.ReadFile("test.expected.csv")
	if err != nil {
		t.Fatalf("Unable to read expected csv file: %v", err)
	}

	actual := b.String()
	expected := string(file)

	if actual != expected {
		t.Fatalf("CSV contents differed, expected (%v), actual (%v)", expected, actual)
	}
}
