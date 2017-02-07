package resume

import (
	"time"

	"github.com/shurcooL/component"
	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/reactions"
	resumecomponent "github.com/shurcooL/resume/component"
	"github.com/shurcooL/users"
	"golang.org/x/net/html"
)

// ReactableURL is the URL for reactionable items on this resume.
const ReactableURL = "dmitri.shuralyov.com/resume"

// DmitriShuralyov is a person whose resume is on display.
type DmitriShuralyov struct {
	Reactions   map[string][]reactions.Reaction
	CurrentUser users.User
}

func (DmitriShuralyov) experience() htmlg.Component {
	return resumecomponent.Section{
		Title: "Experience",

		Items: []resumecomponent.Item{
			{
				Title:    "Senior Software Engineer, Full Stack",
				Subtitle: "Sourcegraph",
				Dates: resumecomponent.DateRange{
					From: resumecomponent.Date{Year: 2015, Month: time.April}, To: resumecomponent.Date{Year: 2016, Month: time.November},
				},
				Lines: []htmlg.Component{
					component.Text("Implemented significant non-trivial pieces of core Sourcegraph functionality in Go, including backend language analysis enhancements, and frontend components and visualizations."),
					component.Text("Showed initiative by taking on refactors that led to significant performance improvements."),
					component.Text("Made numerous contributions to open source Go libraries created or used by Sourcegraph."),
					component.Text("Shared knowledge and best practices with new teammates to enable high quality contributions."),
				},
			},
			{
				Title:    "Senior Software Engineer, Backend",
				Subtitle: "Triggit",
				Dates: resumecomponent.DateRange{
					From: resumecomponent.Date{Year: 2013, Month: time.June}, To: resumecomponent.Date{Year: 2015, Month: time.March},
				},
				Lines: []htmlg.Component{
					component.Text("Built distributed low-latency web services and required components for processing hundreds of thousands of ad auction requests per second."),
					component.Text("Automated, improved practices for reproducible builds, continuous testing of complex projects."),
					component.Text("Improved performance and functionality of an ad-serving and bidding platform."),
					component.Text("Created detailed dashboards for monitoring and visualizing logs, statistics, controlling configuration and other relevant metrics."),
				},
			},
			{
				Title: "Toolmaker",
				Dates: resumecomponent.DateRange{
					From: resumecomponent.Date{Year: 2011, Month: time.December}, To: resumecomponent.Date{Year: 2013, Month: time.June},
				},
				Lines: []htmlg.Component{
					component.Text("Researched and implemented experimental software development tools."),
					component.Join("Created Conception, a 1st place winning project of ", component.Link{Text: "LIVE 2013 Programming Contest", URL: "http://liveprogramming.github.io/liveblog/2013/04/live-programming-contest-winners/", NewTab: true}, "."),
				},
			},
			{
				Title:    "Junior Application Programmer",
				Subtitle: "CYSSC/MCYS, Ontario Public Service",
				Dates: resumecomponent.DateRange{
					From: resumecomponent.Date{Year: 2007, Month: time.September}, To: resumecomponent.Date{Year: 2008, Month: time.August},
				},
				Lines: []htmlg.Component{
					component.Text("Designed, created and maintained a complex Java GUI application to aid the development and maintenance of large database applications."),
					component.Text("Wrote PL/SQL procedures to easily enable/disable logging on Oracle DB server on the fly."),
					component.Text("Researched the best approach for new Monitoring Report development; implemented it."),
				},
			},
			{
				Title:    "Game Engine Engineer, Tools",
				Subtitle: "Reverie World Studios",
				Dates: resumecomponent.DateRange{
					From: resumecomponent.Date{Year: 2007, Month: time.January}, To: resumecomponent.Date{Year: 2007, Month: time.August},
				},
				Lines: []htmlg.Component{
					component.Text("Coordinated the development of an upgraded world editor in C# to help streamline content production."),
					component.Text("Engineered a flexible system for reading/writing custom content file formats."),
					component.Text("Improved the performance of the real-time landscape shadowing mechanism."),
				},
			},
		},
	}
}

func (DmitriShuralyov) projects() htmlg.Component {
	return resumecomponent.Section{
		Title: "Projects",

		Items: []resumecomponent.Item{
			{
				Title: "Conception",
				Dates: resumecomponent.DateRange{
					From: resumecomponent.Date{Year: 2012}, To: resumecomponent.Date{Year: 2014},
				},
				Lines: []htmlg.Component{
					component.Text("Primary creator of a large open-source systems project; implemented in C++ and Go, solved low-level systems challenges to achieve desired behavior."),
					component.Text("Routinely implemented and iterated upon experimental and novel interface ideas, interaction techniques and design prototypes, some showed great promise."),
					component.Text("Discovered new techniques that allow for further reduction of information duplication than existing representations."),
					component.Join("1st place winning project of ", component.Link{Text: "LIVE 2013 Programming Contest", URL: "http://liveprogramming.github.io/liveblog/2013/04/live-programming-contest-winners/", NewTab: true}, "."),
				},
			},
			{
				Title: "Slide: A User-Friendly System for Rapid and Precise Object Placement",
				Dates: resumecomponent.Date{Year: 2011},
				Lines: []htmlg.Component{
					component.Text("Implemented in C++ with OpenGL, GLSL graphics, employed multiple advanced graphics optimization techniques to get high performance and accurate results in difficult conditions."),
					component.Text("Had weekly meetings with supervisor to discuss and determine the project direction, iterated based on feedback."),
				},
			},
			{
				Title: "Project eX0",
				Dates: resumecomponent.Date{Year: 2007},
				Lines: []htmlg.Component{
					component.Text("Implemented in C++ with OpenGL graphics."),
					component.Text("Developed own high-performance and reliable networking protocol over raw TCP/UDP sockets, which uniquely combined beneficial properties of past networking models."),
				},
				WIP: true,
			},
		},
	}
}

func (ds DmitriShuralyov) skills() htmlg.Component {
	return resumecomponent.Section{
		Title: "Skills",

		Items: []resumecomponent.Item{
			{
				Title: "Languages and APIs",
				Lines: []htmlg.Component{
					ds.reactable("Go", component.Text("Go")),
					ds.reactable("C/C++", resumecomponent.Fade("C/C++")),
					ds.reactable("Java", resumecomponent.Fade("Java")),
					ds.reactable("C#", resumecomponent.Fade("C#")),
					ds.reactable("OpenGL", component.Text("OpenGL")),
					ds.reactable("SQL", resumecomponent.Fade("SQL")),
				},
			},
			{
				Title: "Software",
				Lines: []htmlg.Component{
					ds.reactable("Git", component.Text("Git")),
					ds.reactable("Xcode", component.Text("Xcode")),
					ds.reactable("Visual Studio", component.Text("Visual Studio")),
					ds.reactable("OS X", component.Text("OS X")),
					ds.reactable("Linux", component.Text("Linux")),
					ds.reactable("Windows", component.Text("Windows")),
				},
			},
		},
	}
}

func (DmitriShuralyov) education() htmlg.Component {
	return resumecomponent.Section{
		Title: "Education",

		Items: []resumecomponent.Item{
			{
				Title: "York University",
				Dates: resumecomponent.DateRange{
					From: resumecomponent.Date{Year: 2009}, To: resumecomponent.Date{Year: 2011},
				},
				Lines: []htmlg.Component{
					component.Text("Master's Degree, Computer Science"),
				},
			},
			{
				Title: "York University",
				Dates: resumecomponent.DateRange{
					From: resumecomponent.Date{Year: 2004}, To: resumecomponent.Date{Year: 2009},
				},
				Lines: []htmlg.Component{
					component.Text("Bachelor's Degree, Specialized Honors Computer Science"),
				},
			},
		},
	}
}

// reactable is a convenience helper that joins reactable content with its ReactionsBar,
// using id as reactable ID. It populates ReactionsBar's Reactions, CurrentUser fields from ds.
func (ds DmitriShuralyov) reactable(id string, content htmlg.Component) htmlg.Component {
	reactionsBar := resumecomponent.ReactionsBar{
		Reactions:    ds.Reactions[id],
		ReactableURL: ReactableURL,
		CurrentUser:  ds.CurrentUser,
		ID:           id,
	}
	return component.List{content, reactionsBar}
}

func (ds DmitriShuralyov) Render() []*html.Node {
	// TODO: Make this much nicer. Less verbose, more readable, more flexible. Maybe like JSX.
	/*
		<div class="name">Dmitri Shuralyov</div>
		<div class="contactinfo"><a href="https://github.com/shurcooL">github.com/shurcooL</a> &middot; <a href="mailto:shurcooL@gmail.com">shurcooL@gmail.com</a></div>
		<div class="core">
			{{render .Experience}}
			{{render .Projects}}
			{{render .Skills}}
			{{render .Education}}
		</div>
	*/
	var ns []*html.Node
	ns = append(ns, htmlg.DivClass("name", htmlg.Text("Dmitri Shuralyov")))
	contactInfo := htmlg.DivClass("contactinfo", component.Join(
		component.Link{Text: "github.com/shurcooL", URL: "https://github.com/shurcooL", NewTab: true},
		" · ",
		component.Link{Text: "shurcooL@gmail.com", URL: "mailto:shurcooL@gmail.com", NewTab: true},
	).Render()...)
	ns = append(ns, contactInfo)
	core := htmlg.DivClass("core")
	for _, n := range ds.experience().Render() {
		core.AppendChild(n)
	}
	for _, n := range ds.projects().Render() {
		core.AppendChild(n)
	}
	for _, n := range ds.skills().Render() {
		core.AppendChild(n)
	}
	for _, n := range ds.education().Render() {
		core.AppendChild(n)
	}
	ns = append(ns, core)
	return ns
}
