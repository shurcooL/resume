package main

import (
	"fmt"

	"golang.org/x/net/html"
)

// Date represents an imprecise date.
type Date struct {
	Year  int // Year. E.g., 2009.
	Month int // Month is 1 - 12. 0 means unspecified.
}

func (d Date) Render() ([]*html.Node, error) {
	switch d.Month {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		return Text(fmt.Sprintf("%02d/%d", d.Month, d.Year)).Render()
	default:
		return Text(fmt.Sprintf("%d", d.Year)).Render()
	}
}

// Present represents the present date.
var Present = Text("Present")

// DateRange represents a span of time between two dates.
type DateRange struct {
	From, To Component
}

func (dr DateRange) Render() ([]*html.Node, error) {
	return List{dr.From, Text(" - "), dr.To}.Render()
}
