package resume

import (
	"html/template"

	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// DmitriShuralyov is a person whose resume is on display.
type DmitriShuralyov struct{}

func (DmitriShuralyov) Experience() Section {
	return Section{
		Title: "Experience",

		Items: []Item{
			{
				Title:    "Senior Software Engineer, Full Stack",
				Subtitle: "Sourcegraph",
				Dates: DateRange{
					From: Date{Year: 2015, Month: 4}, To: Present{},
				},
				Lines: []Component{
					Text("Implemented significant non-trivial pieces of core Sourcegraph functionality in Go, including backend language analysis enhancements, and frontend components and visualizations."),
					Text("Showed initiative by taking on refactors that led to significant performance improvements."),
					Text("Made numerous contributions to open source Go libraries created or used by Sourcegraph."),
					Text("Shared knowledge and best practices with new teammates to enable high quality contributions."),
				},
			},
			{
				Title:    "Senior Software Engineer, Backend",
				Subtitle: "Triggit",
				Dates: DateRange{
					From: Date{Year: 2013, Month: 6}, To: Date{Year: 2015, Month: 3},
				},
				Lines: []Component{
					Text("Built distributed low-latency web services and required components for processing hundreds of thousands of ad auction requests per second."),
					Text("Automated, improved practices for reproducible builds, continuous testing of complex projects."),
					Text("Improved performance and functionality of an ad-serving and bidding platform."),
					Text("Created detailed dashboards for monitoring and visualizing logs, statistics, controlling configuration and other relevant metrics."),
				},
			},
			{
				Title: "Toolmaker",
				Dates: DateRange{
					From: Date{Year: 2012}, To: Date{Year: 2013, Month: 6},
				},
				Lines: []Component{
					Text("Researched and implemented experimental software development tools."),
					join("Created Conception, a 1st place winning project of ", Link{"LIVE 2013 Programming Contest", template.URL("http://liveprogramming.github.io/liveblog/2013/04/live-programming-contest-winners/")}, "."),
				},
			},
			{
				Title:    "Junior Application Programmer",
				Subtitle: "CYSSC/MCYS, Ontario Public Service",
				Dates: DateRange{
					From: Date{Year: 2007, Month: 9}, To: Date{Year: 2008, Month: 8},
				},
				Lines: []Component{
					Text("Designed, created and maintained a complex Java GUI application to aid the development and maintenance of large database applications."),
					Text("Wrote PL/SQL procedures to easily enable/disable logging on Oracle DB server on the fly."),
					Text("Researched the best approach for new Monitoring Report development; implemented it."),
				},
			},
			{
				Title:    "Game Engine Engineer, Tools",
				Subtitle: "Reverie World Studios",
				Dates: DateRange{
					From: Date{Year: 2007, Month: 1}, To: Date{Year: 2007, Month: 8},
				},
				Lines: []Component{
					Text("Coordinated the development of an upgraded world editor in C# to help streamline content production."),
					Text("Engineered a flexible system for reading/writing custom content file formats."),
					Text("Improved the performance of the real-time landscape shadowing mechanism."),
				},
			},
		},
	}
}

func (DmitriShuralyov) Projects() Section {
	return Section{
		Title: "Projects",

		Items: []Item{
			{
				Title: "Conception",
				Dates: DateRange{
					From: Date{Year: 2012}, To: Date{Year: 2014},
				},
				Lines: []Component{
					Text("Primary creator of a large open-source systems project; implemented in C++ and Go, solved low-level systems challenges to achieve desired behavior."),
					Text("Routinely implemented and iterated upon experimental and novel interface ideas, interaction techniques and design prototypes, some showed great promise."),
					Text("Discovered new techniques that allow for further reduction of information duplication than existing representations."),
					join("1st place winning project of ", Link{"LIVE 2013 Programming Contest", template.URL("http://liveprogramming.github.io/liveblog/2013/04/live-programming-contest-winners/")}, "."),
				},
			},
			{
				Title: "Slide: A User-Friendly System for Rapid and Precise Object Placement",
				Dates: Date{Year: 2011},
				Lines: []Component{
					Text("Implemented in C++ with OpenGL, GLSL graphics, employed multiple advanced graphics optimization techniques to get high performance and accurate results in difficult conditions."),
					Text("Had weekly meetings with supervisor to discuss and determine the project direction, iterated based on feedback."),
				},
			},
			{
				Title: "Project eX0",
				Dates: Date{Year: 2007},
				Lines: []Component{
					Text("Implemented in C++ with OpenGL graphics."),
					Text("Developed own high-performance and reliable networking protocol over raw TCP/UDP sockets, which uniquely combined beneficial properties of past networking models."),
				},
				WIP: true,
			},
		},
	}
}

func (DmitriShuralyov) Skills() Section {
	return Section{
		Title: "Skills",

		Items: []Item{
			{
				Title: "Languages and APIs",
				Lines: []Component{
					Reactable{ID: "Go", Content: Text("Go")},
					Reactable{ID: "C/C++", Content: fade("C/C++")},
					Reactable{ID: "Java", Content: fade("Java")},
					Reactable{ID: "C#", Content: fade("C#")},
					Reactable{ID: "OpenGL", Content: Text("OpenGL")},
					Reactable{ID: "SQL", Content: fade("SQL")},
				},
			},
			{
				Title: "Software",
				Lines: []Component{
					Reactable{ID: "Git", Content: Text("Git")},
					Reactable{ID: "Xcode", Content: Text("Xcode")},
					Reactable{ID: "Visual Studio", Content: Text("Visual Studio")},
					Reactable{ID: "OS X", Content: Text("OS X")},
					Reactable{ID: "Linux", Content: Text("Linux")},
					Reactable{ID: "Windows", Content: Text("Windows")},
				},
			},
		},
	}
}

func (DmitriShuralyov) Education() Section {
	return Section{
		Title: "Education",

		Items: []Item{
			{
				Title: "York University",
				Dates: DateRange{
					From: Date{Year: 2009}, To: Date{Year: 2011},
				},
				Lines: []Component{
					Text("Master's Degree, Computer Science"),
				},
			},
			{
				Title: "York University",
				Dates: DateRange{
					From: Date{Year: 2004}, To: Date{Year: 2009},
				},
				Lines: []Component{
					Text("Bachelor's Degree, Specialized Honors Computer Science"),
				},
			},
		},
	}
}

func (ds DmitriShuralyov) Render() []*html.Node {
	// TODO: Make this much nicer. Less verbose, more readable, more flexible. Maybe like JSX.
	/*
		<div class="name">Dmitri Shuralyov</div>
		<div class="contactinfo"><a href="https://github.com/shurcooL" target="_blank">github.com/shurcooL</a> &middot; <a href="mailto:shurcooL@gmail.com" target="_blank">shurcooL@gmail.com</a></div>
		<div class="core">
			{{render .Experience}}
			{{render .Projects}}
			{{render .Skills}}
			{{render .Education}}
		</div>
	*/
	var ns []*html.Node
	ns = append(ns, htmlg.DivClass("name", htmlg.Text("Dmitri Shuralyov")))
	contactInfo := htmlg.DivClass("contactinfo",
		htmlg.A("github.com/shurcooL", template.URL("https://github.com/shurcooL")),
		htmlg.Text(" · "),
		htmlg.A("shurcooL@gmail.com", template.URL("mailto:shurcooL@gmail.com")),
	)
	contactInfo.FirstChild.Attr = append(contactInfo.FirstChild.Attr, html.Attribute{Key: atom.Target.String(), Val: "_blank"})
	contactInfo.LastChild.Attr = append(contactInfo.LastChild.Attr, html.Attribute{Key: atom.Target.String(), Val: "_blank"})
	ns = append(ns, contactInfo)
	core := htmlg.DivClass("core")
	for _, n := range ds.Experience().Render() {
		core.AppendChild(n)
	}
	for _, n := range ds.Projects().Render() {
		core.AppendChild(n)
	}
	for _, n := range ds.Skills().Render() {
		core.AppendChild(n)
	}
	for _, n := range ds.Education().Render() {
		core.AppendChild(n)
	}
	ns = append(ns, core)
	return ns
}

type Section struct {
	Title string
	Items []Item
}

func (s Section) Render() []*html.Node {
	// TODO: Make this much nicer.
	/*
		<div class="sectionheader">{{.Title}}</div>
		{{range .Items}}
			{{if not .WIP}}
				{{render .}}
			{{end}}
		{{end}}
	*/
	var ns []*html.Node
	ns = append(ns, htmlg.DivClass("sectionheader", htmlg.Text(s.Title)))
	for _, i := range s.Items {
		if i.WIP {
			continue
		}
		for _, n := range i.Render() {
			ns = append(ns, n)
		}
	}
	return ns
}

type Item struct {
	Title    string
	Subtitle string
	Dates    Component
	Lines    []Component

	WIP bool
}

func (i Item) Render() []*html.Node {
	// TODO: Make this much nicer.
	/*
		<div class="item{{if .WIP}} wip{{end}}">
			<div class="itemheader">
				<div class="title">{{.Title}}</div>
				{{with .Subtitle}}<div class="subtitle">{{.}}</div>{{end}}
				{{with .Dates}}<div class="dates">{{render .}}</div>{{end}}
			</div>
			<ul>
				{{range .Lines}}<li>{{render .}}</li>
				{{end}}
			</ul>
		</div>
	*/
	itemClass := "item"
	if i.WIP {
		itemClass += " wip"
	}
	item := htmlg.DivClass(itemClass)

	itemHeader := htmlg.DivClass("itemheader")
	itemHeader.AppendChild(htmlg.DivClass("title", htmlg.Text(i.Title)))
	if i.Subtitle != "" {
		itemHeader.AppendChild(htmlg.DivClass("subtitle", htmlg.Text(i.Subtitle)))
	}
	if i.Dates != nil {
		itemHeader.AppendChild(htmlg.DivClass("dates", i.Dates.Render()...))
	}
	item.AppendChild(itemHeader)

	ul := &html.Node{Type: html.ElementNode, Data: atom.Ul.String()}
	for _, l := range i.Lines {
		li := &html.Node{Type: html.ElementNode, Data: atom.Li.String()}
		for _, n := range l.Render() {
			li.AppendChild(n)
		}
		ul.AppendChild(li)
	}
	item.AppendChild(ul)

	return []*html.Node{item}
}
