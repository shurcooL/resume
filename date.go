package main

import (
	"fmt"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Date represents an imprecise date.
type Date struct {
	Year  int // Year. E.g., 2009.
	Month int // Month is 1 - 12. 0 means unspecified. // TODO: Use time.Month.
}

func (d Date) Date() (year int, month int) { return d.Year, d.Month }

func (d Date) Render() ([]*html.Node, error) {
	switch d.Month {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		return Text(fmt.Sprintf("%02d/%d", d.Month, d.Year)).Render()
	default:
		return Text(fmt.Sprintf("%d", d.Year)).Render()
	}
}

// Present represents the present date.
type Present struct{}

func (Present) Date() (year int, month int) {
	y, m, _ := time.Now().UTC().Date()
	return y, int(m)
}

func (p Present) Render() ([]*html.Node, error) {
	return Text("Present").Render()
}

type Dater interface {
	Date() (year int, month int)
	Component
}

// DateRange represents a span of time between two dates.
type DateRange struct {
	From, To Dater
}

func (dr DateRange) Render() ([]*html.Node, error) {
	nodes, err := List{dr.From, Text(" - "), dr.To}.Render()
	if err != nil {
		return nil, err
	}
	span := span(nodes...)
	span.Attr = append(span.Attr, html.Attribute{Key: atom.Title.String(), Val: yearsMonths(dr.From, dr.To)})
	return []*html.Node{span}, nil
}

// yearsMonths describes the length of a date range in the number of years and months.
func yearsMonths(from, to Dater) string {
	y0, m0 := from.Date()
	y1, m1 := to.Date()
	months := (y1-y0)*12 + m1 - m0
	years, months := months/12, months%12
	switch {
	case years == 0:
		return fmt.Sprintf("%d months", months)
	case years != 0 && months == 0:
		return fmt.Sprintf("%d years", years)
	default:
		return fmt.Sprintf("%d years, %d months", years, months)
	}
}
