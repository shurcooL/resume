package main

import "fmt"

// Date represents an imprecise date.
type Date struct {
	Year  int // Year. E.g., 2009.
	Month int // Month is 1 - 12. 0 means unspecified.
}

func (d Date) String() string {
	switch d.Month {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		return fmt.Sprintf("%02d/%d", d.Month, d.Year)
	default:
		return fmt.Sprintf("%d", d.Year)
	}
}

// Present represents the present date.
const Present = "Present"

// DateRange represents a span of time between two dates.
type DateRange struct {
	From, To interface{}
}

func (dr DateRange) String() string {
	return fmt.Sprintf("%v - %v", dr.From, dr.To)
}
