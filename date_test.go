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

func render(c Component) {
	nodes, err := c.Render()
	if err != nil {
		panic(err)
	}
	fmt.Println(htmlg.Render(nodes...))
}
