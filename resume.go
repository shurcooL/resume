// +build js

package main

import (
	"bytes"
	"html/template"

	"github.com/shurcooL/htmlg"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		go setup()
	})
}

func setup() {
	var buf bytes.Buffer
	err := t().ExecuteTemplate(&buf, "body", DmitriShuralyov{})
	if err != nil {
		panic(err)
	}
	document.Body().SetInnerHTML(buf.String())
}

type DmitriShuralyov struct{}

func (DmitriShuralyov) Experience() Section { return experience }

func (DmitriShuralyov) Projects() Section { return projects }

func (DmitriShuralyov) Education() Section { return education }

type Section struct {
	Title string
	Items []Item
}

type Item struct {
	JobTitle    string
	CompanyName string
	Dates       Component
	Lines       []Component

	WIP bool
}

var experience = Section{
	Title: "Experience",

	Items: []Item{
		{
			JobTitle:    "Senior Software Engineer, Full Stack",
			CompanyName: "Sourcegraph",
			Dates: DateRange{
				From: Date{Year: 2015, Month: 4}, To: Present{},
			},
			Lines: []Component{
				Text("Implemented significant non­trivial pieces of core Sourcegraph functionality in Go, including backend language analysis enhancements, and frontend components and visualizations."),
				Text("Showed initiative by taking on refactors that led to significant performance improvements."),
				Text("Made numerous contributions to open source Go libraries created or used by Sourcegraph."),
				Text("Shared knowledge and best practices with new teammates to enable high quality contributions."),
			},
		},
		{
			JobTitle:    "Senior Software Engineer, Backend",
			CompanyName: "Triggit",
			Dates: DateRange{
				From: Date{Year: 2013, Month: 6}, To: Date{Year: 2015, Month: 3},
			},
			Lines: []Component{
				Text("Built distributed low­latency web services and required components for processing hundreds of thousands of ad auction requests per second."),
				Text("Automated, improved practices for reproducible builds, continuous testing of complex projects."),
				Text("Improved performance and functionality of an ad­serving and bidding platform."),
				Text("Created detailed dashboards for monitoring and visualizing logs, statistics, controlling configuration and other relevant metrics."),
			},
		},
		{
			JobTitle: "Toolmaker",
			Dates: DateRange{
				From: Date{Year: 2012}, To: Date{Year: 2013, Month: 6},
			},
			Lines: []Component{
				Text("Researched and implemented experimental software development tools."),
				join("Created Conception, a 1st place winning project of ", Link{"LIVE 2013 Programming Contest", template.URL("http://liveprogramming.github.io/liveblog/2013/04/live-programming-contest-winners/")}, "."),
			},
		},
		{
			JobTitle:    "Junior Application Programmer",
			CompanyName: "CYSSC/MCYS, Ontario Public Service",
			Dates: DateRange{
				From: Date{Year: 2007, Month: 9}, To: Date{Year: 2008, Month: 8},
			},
			Lines: []Component{
				Text("Designed, created and maintained a complex Java GUI application to aid the development and maintenance of large database applications."),
				Text("Wrote PL/SQL procedures to easily enable/disable logging on Oracle DB server on the fly."),
				Text("Researched the best approach for new Monitoring Report development; implemented it."),
			},
		},
		{
			JobTitle:    "Game Engine Engineer, Tools",
			CompanyName: "Reverie World Studios",
			Dates: DateRange{
				From: Date{Year: 2007, Month: 1}, To: Date{Year: 2007, Month: 8},
			},
			Lines: []Component{
				Text("Coordinated the development of an upgraded world editor in C# to help streamline content production."),
				Text("Engineered a flexible system for reading/writing custom content file formats."),
				Text("Improved the performance of the real­time landscape shadowing mechanism."),
			},
		},
	},
}

var projects = Section{
	Title: "Projects",

	Items: []Item{
		// TODO: ProjectItems?
		{
			JobTitle: "Conception",
			Dates: DateRange{
				From: Date{Year: 2012}, To: Date{Year: 2014},
			},
			Lines: []Component{
				Text("Primary creator of a large open­source systems project; implemented in C++ and Go, solved low­level systems challenges to achieve desired behavior."),
				Text("Routinely implemented and iterated upon experimental and novel interface ideas, interaction techniques and design prototypes, some showed great promise."),
				Text("Discovered new techniques that allow for further reduction of information duplication than existing representations."),
				join("1st place winning project of ", Link{"LIVE 2013 Programming Contest", template.URL("http://liveprogramming.github.io/liveblog/2013/04/live-programming-contest-winners/")}, "."),
			},
		},
		{
			JobTitle: "Slide: A User­-Friendly System for Rapid and Precise Object Placement",
			Dates:    Date{Year: 2011},
			Lines: []Component{
				Text("Implemented in C++ with OpenGL, GLSL graphics, employed multiple advanced graphics optimization techniques to get high performance and accurate results in difficult conditions."),
				Text("Had weekly meetings with supervisor to discuss and determine the project direction, iterated based on feedback."),
			},
		},
		{
			JobTitle: "Project eX0",
			Dates:    Date{Year: 2007},
			Lines: []Component{
				Text("Implemented in C++ with OpenGL graphics."),
				Text("Developed own high­-performance and reliable networking protocol over raw TCP/UDP sockets, which uniquely combined beneficial properties of past networking models."),
			},
			WIP: true,
		},
	},
}

var education = Section{
	Title: "Education",

	Items: []Item{
		// TODO: EducationItems?
		{
			JobTitle: "York University",
			Dates: DateRange{
				From: Date{Year: 2009}, To: Date{Year: 2011},
			},
			Lines: []Component{
				Text("Master's Degree, Computer Science"),
			},
		},
		{
			JobTitle: "York University",
			Dates: DateRange{
				From: Date{Year: 2004}, To: Date{Year: 2009},
			},
			Lines: []Component{
				Text("Bachelor's Degree, Specialized Honors Computer Science"),
			},
		},
	},
}

func t() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"render": func(c Component) template.HTML { return htmlg.Render(c.Render()...) },
	}).Parse(`
{{define "section"}}
	<div class="sectionheader">{{.Title}}</div>
	{{range .Items}}
		{{template "item" .}}
	{{end}}
{{end}}

{{define "item"}}
<div class="item{{if .WIP}} wip{{end}}">
	<div class="itemheader">
		<div class="jobtitle">{{.JobTitle}}</div>
		{{with .CompanyName}}<div class="companyname">{{.}}</div>{{end}}
		<div class="dates">{{render .Dates}}</div>
	</div>
	<ul>
		{{range .Lines}}<li>{{render .}}</li>
		{{end}}
	</ul>
</div>
{{end}}

{{define "body"}}
	<div class="name">Dmitri Shuralyov</div>
	<div class="contactinfo"><a href="https://github.com/shurcooL" target="_blank">github.com/shurcooL</a> &middot; <a href="mailto:shurcooL@gmail.com" target="_blank">shurcooL@gmail.com</a></div>
	<div class="corediv">
		{{template "section" .Experience}}
		{{template "section" .Projects}}
		{{template "section" .Education}}
		<div class="sectionheader">Knowledge and Skills Highlights</div>
		<div class="item">
			<b>Languages and APIs</b>: Go<span class="fade">, C/C++, Java, C#, </span>OpenGL<span class="fade">, SQL</span>
		</div>
		<div class="item">
			<b>Software</b>: OS X, Linux, Windows, git, Microsoft Visual Studio, Xcode
		</div>
	</div>
{{end}}
`))
}
