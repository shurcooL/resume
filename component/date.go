package component

import (
	"fmt"
	"time"

	"github.com/shurcooL/component"
	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Date represents an imprecise date.
type Date struct {
	Year  int        // Year. E.g., 2009.
	Month time.Month // Month is 1 - 12. 0 means unspecified.
}

func (d Date) Date() (year int, month time.Month) { return d.Year, d.Month }

func (d Date) Render() []*html.Node {
	switch d.Month {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		return component.Text(fmt.Sprintf("%d/%02d", d.Year, d.Month)).Render()
	default:
		return component.Text(fmt.Sprintf("%d", d.Year)).Render()
	}
}

// Present represents the present date.
type Present struct{}

func (Present) Date() (year int, month time.Month) {
	year, month, _ = time.Now().UTC().Date()
	return year, month
}

func (p Present) Render() []*html.Node {
	return component.Text("Present").Render()
}

type DateComponent interface {
	Date() (year int, month time.Month)
	htmlg.Component
}

// DateRange represents a span of time between two dates.
type DateRange struct {
	From, To DateComponent
}

func (dr DateRange) Render() []*html.Node {
	span := htmlg.Span(component.List{dr.From, component.Text("–"), dr.To}.Render()...)

	// Compute the years and months between From and To dates,
	// and set that as the tooltip of DateRange.
	tooltip := yearsMonths(dr.From, dr.To)
	span.Attr = append(span.Attr, html.Attribute{
		Key: atom.Title.String(),
		Val: tooltip,
	})

	return []*html.Node{span}
}

// yearsMonths describes the length of a date range in the number of years and months.
func yearsMonths(from, to DateComponent) string {
	y0, m0 := from.Date()
	y1, m1 := to.Date()
	months := (y1-y0)*12 + int(m1) - int(m0)
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
