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
