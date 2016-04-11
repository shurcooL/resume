package main

import (
	"fmt"

	"github.com/shurcooL/htmlg"
)

func ExampleDate() {
	{
		d := Date{Year: 2015, Month: 6}
		render(d)
	}
	{
		d := Date{Year: 2008}
		render(d)
	}
	{
		d := Present{}
		render(d)
	}

	// Output:
	// 06/2015
	// 2008
	// Present
}

func ExampleDateRange() {
	dr := DateRange{
		From: Date{Year: 2010, Month: 2}, To: Date{Year: 2012, Month: 6},
	}
	render(dr)

	// Output: <span title="2 years, 4 months">02/2010 - 06/2012</span>
}

func render(c Component) { fmt.Println(htmlg.Render(c.Render()...)) }

func ExampleDateRange_yearsMonths() {
	fmt.Println(yearsMonths(Date{Year: 2010, Month: 2}, Date{Year: 2010, Month: 2}))
	fmt.Println(yearsMonths(Date{Year: 2010, Month: 2}, Date{Year: 2010, Month: 8}))
	fmt.Println(yearsMonths(Date{Year: 2010, Month: 2}, Date{Year: 2011, Month: 2}))
	fmt.Println(yearsMonths(Date{Year: 2010, Month: 2}, Date{Year: 2012, Month: 6}))

	// Output:
	// 0 months
	// 6 months
	// 1 years
	// 2 years, 4 months
}
