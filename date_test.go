package main_test

import (
	"fmt"

	"github.com/shurcooL/htmlg"
	. "github.com/shurcooL/resume"
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
		d := Present
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

	// Output: 02/2010 - 06/2012
}

func render(c Component) {
	nodes, err := c.Render()
	if err != nil {
		panic(err)
	}
	fmt.Println(htmlg.Render(nodes...))
}
