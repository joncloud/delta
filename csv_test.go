package main

import (
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
	// TODO
}
