// +build js

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/users"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		go setup()
	})
}

func setup() {
	authenticatedUser, err := getAuthenticatedUser()
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = t().ExecuteTemplate(&buf, "body", DmitriShuralyov{AuthenticatedUser: authenticatedUser})
	if err != nil {
		panic(err)
	}
	document.Body().SetInnerHTML(buf.String())

	setupReactionsMenu(authenticatedUser != nil)
}

func getAuthenticatedUser() (*users.User, error) {
	resp, err := http.Get("http://localhost:8080/user") // TODO: "/user"?
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var u users.User
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

type DmitriShuralyov struct {
	AuthenticatedUser *users.User
}

func (DmitriShuralyov) Experience() Section { return experience }

func (DmitriShuralyov) Projects() Section { return projects }

func (DmitriShuralyov) Skills() Section { return skills }

func (DmitriShuralyov) Education() Section { return education }

type Section struct {
	Title string
	Items []Item
}

type Item struct {
	Title    string
	Subtitle string
	Dates    Component
	Lines    []Component

	WIP bool
}

var experience = Section{
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

var projects = Section{
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

var skills = Section{
	Title: "Skills",

	Items: []Item{
		{
			Title: "Languages and APIs",
			Lines: []Component{
				reactable("Go", Text("Go")),
				reactable("C/C++", fade("C/C++")),
				reactable("Java", fade("Java")),
				reactable("C#", fade("C#")),
				reactable("OpenGL", Text("OpenGL")),
				reactable("SQL", fade("SQL")),
			},
		},
		{
			Title: "Software",
			Lines: []Component{
				reactable("OS X", Text("OS X")),
				reactable("Linux", Text("Linux")),
				reactable("Windows", Text("Windows")),
				reactable("git", Text("git")),
				reactable("Visual Studio", Text("Visual Studio")),
				reactable("Xcode", Text("Xcode")),
			},
		},
	},
}

var education = Section{
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

func reactable(id string, a ...interface{}) Component {
	list := join(a...)
	list = append(list, NewReaction{ReactableID: id})
	return list
}

type NewReaction struct {
	ReactableID string
}

func (nr NewReaction) Render() []*html.Node {
	// TODO: Make this much nicer.
	/*
		<a href="javascript:" title="React" onclick="ShowReactionMenu(this, event, {{.}});">
			<div class="new-reaction">
				<i class="octicon octicon-smiley"><sup>+</sup></i>
			</div>
		</a>
	*/
	sup := &html.Node{
		Type: html.ElementNode, Data: atom.Sup.String(),
	}
	sup.AppendChild(htmlg.Text("+"))
	i := &html.Node{
		Type: html.ElementNode, Data: atom.I.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: "octicon octicon-smiley"}},
	}
	i.AppendChild(sup)
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Div.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: "new-reaction"}},
	}
	div.AppendChild(i)
	a := &html.Node{
		Type: html.ElementNode, Data: atom.A.String(),
		Attr: []html.Attribute{
			{Key: atom.Href.String(), Val: "javascript:"},
			{Key: atom.Title.String(), Val: "React"},
			{Key: atom.Onclick.String(), Val: fmt.Sprintf("ShowReactionMenu(this, event, '%q');", nr.ReactableID)},
		},
	}
	a.AppendChild(div)
	return []*html.Node{a}
}

func t() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"render": func(c Component) template.HTML { return htmlg.Render(c.Render()...) },
	}).Parse(`
{{define "section"}}
	<div class="sectionheader">{{.Title}}</div>
	{{range .Items}}
		{{if not .WIP}}
			{{template "item" .}}
		{{end}}
	{{end}}
{{end}}

{{define "item"}}
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
{{end}}

{{define "body"}}
	{{with .AuthenticatedUser}}
		<div style="text-align: right; margin-bottom: 20px; height: 18px; font-size: 12px;">
			<a class="topbar-avatar" href="{{.HTMLURL}}" target="_blank" tabindex=-1
				><img class="topbar-avatar" src="{{.AvatarURL}}" title="Signed in as {{.Login}}."
			></a>
			<form method="post" action="/logout" style="display: inline-block; margin-bottom: 0;"><input class="btn" type="submit" value="Sign out"><input type="hidden" name="return" value="/"></form> {{/* TODO: Valid value for return. */}}
		</div>
	{{end}}
	<div class="name">Dmitri Shuralyov</div>
	<div class="contactinfo"><a href="https://github.com/shurcooL" target="_blank">github.com/shurcooL</a> &middot; <a href="mailto:shurcooL@gmail.com" target="_blank">shurcooL@gmail.com</a></div>
	<div class="corediv">
		{{template "section" .Experience}}
		{{template "section" .Projects}}
		{{template "section" .Skills}}
		{{template "section" .Education}}
	</div>
{{end}}
`))
}
