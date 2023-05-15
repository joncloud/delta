package main

import (
	"fmt"
	"time"
)

type DurationLayout string

const (
	Hours        DurationLayout = "H"
	Minutes      DurationLayout = "M"
	Seconds      DurationLayout = "s"
	Milliseconds DurationLayout = "m"
	Microseconds DurationLayout = "u"
	Nanoseconds  DurationLayout = "n"
)

func ParseDuration(text string) (DurationLayout, error) {
	switch text {
	case "H":
		return Hours, nil
	case "M":
		return Minutes, nil
	case "s":
		return Seconds, nil
	case "m":
		return Milliseconds, nil
	case "u":
		return Microseconds, nil
	case "n":
		return Nanoseconds, nil
	}
	return Seconds, fmt.Errorf("%s is not a valid duration layout", text)
}

func FormatDuration(duration time.Duration, layout DurationLayout) string {
	// TODO have a better layout parser
	switch layout {
	case Hours:
		return fmt.Sprintf("%f", duration.Hours())
	case Minutes:
		return fmt.Sprintf("%f", duration.Minutes())
	case Seconds:
		return fmt.Sprintf("%f", duration.Seconds())
	case Milliseconds:
		return fmt.Sprintf("%d", duration.Milliseconds())
	case Microseconds:
		return fmt.Sprintf("%d", duration.Microseconds())
	case Nanoseconds:
		return fmt.Sprintf("%d", duration.Nanoseconds())
	}
	return ""
}
