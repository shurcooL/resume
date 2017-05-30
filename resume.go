package resume

import (
	"time"

	"github.com/shurcooL/component"
	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/reactions"
	reactionscomponent "github.com/shurcooL/reactions/component"
	resumecomponent "github.com/shurcooL/resume/component"
	"github.com/shurcooL/users"
)

// ReactableURL is the URL for reactionable items on this resume.
const ReactableURL = "dmitri.shuralyov.com/resume"

// DmitriShuralyov returns Dmitri Shuralyov's resume.
// It's annotated with the given reactions from the perspective of currentUser.
func DmitriShuralyov(reactions map[string][]reactions.Reaction, currentUser users.User) htmlg.Component {
	// reactable is a convenience helper that joins reactable content with its ReactionsBar,
	// using id as reactable ID. It populates ReactionsBar's Reactions and CurrentUser fields.
	reactable := func(id string, content htmlg.Component) htmlg.Component {
		reactionsBar := reactionscomponent.ReactionsBar{
			Reactions:    reactions[id],
			ReactableURL: ReactableURL,
			CurrentUser:  currentUser,
			ID:           id,
		}
		return component.List{content, reactionsBar}
	}

	resume := component.Join(
		resumecomponent.Name{
			Name: "Dmitri Shuralyov",
		},

		resumecomponent.ContactInfo{
			GitHub: component.Link{Text: "github.com/shurcooL", URL: "https://github.com/shurcooL", NewTab: true},
			Email:  component.Link{Text: "shurcooL@gmail.com", URL: "mailto:shurcooL@gmail.com", NewTab: true},
		},

		resumecomponent.Section{
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
		},

		resumecomponent.Section{
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
		},

		resumecomponent.Section{
			Title: "Skills",

			Items: []resumecomponent.Item{
				{
					Title: "Languages and APIs",
					Lines: []htmlg.Component{
						reactable("Go", component.Text("Go")),
						reactable("C/C++", resumecomponent.Fade("C/C++")),
						reactable("Java", resumecomponent.Fade("Java")),
						reactable("C#", resumecomponent.Fade("C#")),
						reactable("OpenGL", component.Text("OpenGL")),
						reactable("SQL", resumecomponent.Fade("SQL")),
					},
				},
				{
					Title: "Software",
					Lines: []htmlg.Component{
						reactable("Git", component.Text("Git")),
						reactable("Xcode", component.Text("Xcode")),
						reactable("Visual Studio", component.Text("Visual Studio")),
						reactable("macOS", component.Text("macOS")),
						reactable("Linux", component.Text("Linux")),
						reactable("Windows", component.Text("Windows")),
					},
				},
			},
		},

		resumecomponent.Section{
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
		},
	)

	return resume
}
