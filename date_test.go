package resume

import (
	"fmt"
	"time"

	"github.com/shurcooL/htmlg"
)

func ExampleDate() {
	{
		d := Date{Year: 2015, Month: time.June}
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
	// 2015/06
	// 2008
	// Present
}

func ExampleDateRange() {
	dr := DateRange{
		From: Date{Year: 2010, Month: time.February}, To: Date{Year: 2012, Month: time.June},
	}
	render(dr)

	// Output: <span title="2 years, 4 months">2010/02–2012/06</span>
}

func render(c Component) { fmt.Println(htmlg.Render(c.Render()...)) }

func ExampleDateRange_yearsMonths() {
	fmt.Println(yearsMonths(Date{Year: 2010, Month: time.February}, Date{Year: 2010, Month: time.February}))
	fmt.Println(yearsMonths(Date{Year: 2010, Month: time.February}, Date{Year: 2010, Month: time.August}))
	fmt.Println(yearsMonths(Date{Year: 2010, Month: time.February}, Date{Year: 2011, Month: time.February}))
	fmt.Println(yearsMonths(Date{Year: 2010, Month: time.February}, Date{Year: 2012, Month: time.June}))

	// Output:
	// 0 months
	// 6 months
	// 1 years
	// 2 years, 4 months
}
