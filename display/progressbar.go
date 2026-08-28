package display

import (
	"fmt"
	"io"
	"strings"
)

type ValueStyle int
type TrackStyle int

const (
	HideCursor string = "\033[?25l"
	ShowCursor string = "\033[?25h"

	StyleRatio ValueStyle = iota
	StylePercentage

	StyleDefault TrackStyle = iota
	StyleFilled
	StyleHashtag
)

type TrackerOptions struct {
	UpdateFrequency int
	Styles          Styles
}

type Styles struct {
	ValueStyle ValueStyle
	TrackStyle TrackStyle
}

// track the progress of looping over a slice and executing a function on every element
// maxBarLength Specifies the length of the main progressbar
// optionally an update interval can be provided that specifies that only every nth character of the progressbar is printed
func TrackSlice[T any](slice []T, maxBarLength int, fn func(int, T), optss ...TrackerOptions) {
	prevBarLength := 0

	opts := TrackerOptions{}
	if len(optss) > 0 {
		opts = optss[0]
	}
	specificUpdateInterval := opts.UpdateFrequency

	fmt.Print(HideCursor)
	for i, t := range slice {
		fn(i, t)
		if currentBarLength := ((i + 1) * maxBarLength) / len(slice); currentBarLength > prevBarLength+specificUpdateInterval {
			printTrack(opts.Styles, currentBarLength, maxBarLength, i, len(slice))
			prevBarLength = currentBarLength
		}
	}

	if prevBarLength < maxBarLength {
		printTrack(opts.Styles, maxBarLength, maxBarLength, len(slice), len(slice))
	}
	fmt.Println(ShowCursor)
}

func TrackRange(rangeValue, maxBarLength int, fn func(int), optss ...TrackerOptions) {
	prevBarLength := 0

	opts := TrackerOptions{}
	if len(optss) > 0 {
		opts = optss[0]
	}
	specificUpdateInterval := opts.UpdateFrequency

	fmt.Print(HideCursor)
	for i := range rangeValue {
		fn(i)
		if currentBarLength := ((i + 1) * maxBarLength) / rangeValue; currentBarLength > prevBarLength+specificUpdateInterval {
			printTrack(opts.Styles, currentBarLength, maxBarLength, i, rangeValue)
			prevBarLength = currentBarLength
		}
	}

	if prevBarLength < maxBarLength {
		printTrack(opts.Styles, maxBarLength, maxBarLength, rangeValue, rangeValue)
	}
	fmt.Println(ShowCursor)
}

func printTrack(s Styles, curr, barlength, i, max int) {
	var b strings.Builder
	fPrintTrack(&b, s.TrackStyle, curr, barlength)
	fPrintValues(&b, s.ValueStyle, i, max)
	fmt.Print(b.String())
}

func fPrintValues(w io.Writer, s ValueStyle, curr, max int) {
	switch s {
	case StyleRatio:
		fmt.Fprintf(w, " %v/%v\r", curr+1, max)
	case StylePercentage:
		fmt.Fprintf(w, " %0.2f%%\r", float32(curr+1)/float32(max)*100)
	}
}

func fPrintTrack(w io.Writer, s TrackStyle, curr, max int) {
	switch s {
	case StyleDefault:
		fmt.Fprint(w, "[")
		fmt.Fprintf(w, "%s%s", strings.Repeat("=", curr), strings.Repeat(" ", max-curr))
		fmt.Fprintf(w, "]")
	case StyleFilled:
		fmt.Fprint(w, "🮇")
		fmt.Fprintf(w, "%s%s", strings.Repeat("█", curr), strings.Repeat("🮀", max-curr))
		fmt.Fprintf(w, "▎")
	case StyleHashtag:
		fmt.Fprint(w, "[")
		fmt.Fprintf(w, "%s%s", strings.Repeat("#", curr), strings.Repeat(" ", max-curr))
		fmt.Fprintf(w, "]")
	}
}
