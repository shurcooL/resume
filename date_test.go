package main_test

import (
	"fmt"

	. "github.com/shurcooL/resume"
)

func ExampleDate() {
	{
		d := Date{Year: 2015, Month: 6}
		fmt.Println(d)
	}
	{
		d := Date{Year: 2008}
		fmt.Println(d)
	}
	{
		d := Present
		fmt.Println(d)
	}

	// Output:
	// 06/2015
	// 2008
	// Present
}
