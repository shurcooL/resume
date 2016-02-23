package main

import (
	"bytes"
	"html/template"

	"honnef.co/go/js/dom"
)

var document dom.HTMLDocument

func main() {
	document = dom.GetWindow().Document().(dom.HTMLDocument)
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		go setup()
	})
}

func setup() {
	var buf bytes.Buffer
	err := t().ExecuteTemplate(&buf, "body", sourcegraph)
	if err != nil {
		panic(err)
	}
	document.Body().SetInnerHTML(buf.String())
}

type Item struct {
	JobTitle    string
	CompanyName string
	Dates       DateRange
	Lines       []string
}

var sourcegraph = Item{
	JobTitle:    "Senior Software Engineer, Full Stack",
	CompanyName: "Sourcegraph",
	Dates: DateRange{
		From: Date{Year: 2015, Month: 4},
		To:   Present,
	},
	Lines: []string{
		"Implemented significant non­trivial pieces of core Sourcegraph functionality in Go, including backend language analysis enhancements, and frontend components and visualizations.",
		"Showed initiative by taking on refactors that led to significant performance improvements.",
		"Made numerous contributions to open source Go libraries created or used by Sourcegraph.",
		"Shared knowledge and best practices with new teammates to enable high quality contributions.",
	},
}

func t() *template.Template {
	return template.Must(template.New("").Parse(`
{{define "item"}}
<div class="item">
	<div class="itemheader">
		<div class="jobtitle">{{.JobTitle}}</div>
		<div class="companyname">{{.CompanyName}}</div>
		<div class="dates">{{.Dates}}</div>
	</div>
	<ul>
		{{range .Lines}}<li>{{.}}</li>
		{{end}}
	</ul>
</div>
{{end}}

{{define "body"}}
	<div class="name">Dmitri Shuralyov</div>
	<div class="contactinfo"><a href="https://github.com/shurcooL" target="_blank">github.com/shurcooL</a> &middot; <a href="mailto:shurcooL@gmail.com" target="_blank">shurcooL@gmail.com</a></div>
	<div class="corediv">
		<div class="sectionheader">Experience</div>
		{{template "item" .}}
		<div class="item">
			<div class="itemheader">
				<div class="jobtitle">Senior Software Engineer, Backend</div>
				<div class="companyname">Triggit</div>
				<div class="dates">06/2013 - 03/2015</div>
			</div>
			<ul>
				<li>Built distributed low­latency web services and required components for processing hundreds of thousands of ad auction requests per second.</li>
				<li>Automated, improved practices for reproducible builds, continuous testing of complex projects.</li>
				<li>Improved performance and functionality of an ad­serving and bidding platform.</li>
				<li>Created detailed dashboards for monitoring and visualizing logs, statistics, controlling configuration and other relevant metrics.</li>
			</ul>
		</div>
		<div class="item">
			<div class="itemheader">
				<div class="jobtitle">Toolmaker</div>
				<div class="dates">2012 - 06/2013</div>
			</div>
			<ul>
				<li>Researched and implemented experimental software development tools.</li>
				<li>Created Conception, a 1st place winning project of <a href="http://liveprogramming.github.io/liveblog/2013/04/live-programming-contest-winners/" target="_blank">LIVE 2013 Programming Contest</a>.</li>
			</ul>
		</div>
		<div class="item">
			<div class="itemheader">
				<div class="jobtitle">Junior Application Programmer</div>
				<div class="companyname">CYSSC/MCYS, Ontario Public Service</div>
				<div class="dates">09/2007 - 08/2008</div>
			</div>
			<ul>
				<li>Designed, created and maintained a complex Java GUI application to aid the development and maintenance of large database applications.</li>
				<li>Wrote PL/SQL procedures to easily enable/disable logging on Oracle DB server on the fly.</li>
				<li>Researched the best approach for new Monitoring Report development; implemented it.</li>
			</ul>
		</div>
		<div class="item">
			<div class="itemheader">
				<div class="jobtitle">Game Engine Engineer, Tools</div>
				<div class="companyname">Reverie World Studios</div>
				<div title="8 months" class="dates">01/2007 - 08/2007</div><!-- TODO: Think about the title thing, do it? -->
			</div>
			<ul>
				<li>Coordinated the development of an upgraded world editor in C# to help streamline content production.</li>
				<li>Engineered a flexible system for reading/writing custom content file formats.</li>
				<li>Improved the performance of the real­time landscape shadowing mechanism.</li>
			</ul>
		</div>
		<div class="sectionheader">Projects</div>
		<div class="item">
			<div class="itemheader">
				<div class="projectname">Conception</div>
				<div class="dates">2012 - 2014</div>
			</div>
			<ul>
				<li>Primary creator of a large open­source systems project; implemented in C++ and Go, solved low­level systems challenges to achieve desired behaviour.</li>
				<li>Routinely implemented and iterated upon experimental and novel interface ideas, interaction techniques and design prototypes, some showed great promise.</li>
				<li>Discovered new techniques that allow for further reduction of information duplication than existing representations.</li>
				<li>1st place winning project of <a href="http://liveprogramming.github.io/liveblog/2013/04/live-programming-contest-winners/" target="_blank">LIVE 2013 Programming Contest</a>.</li>
			</ul>
		</div>
		<div class="item">
			<div class="itemheader">
				<div class="projectname">Slide: A User­-Friendly System for Rapid and Precise Object Placement</div>
				<div class="dates">2011</div>
			</div>
			<ul>
				<li>Implemented in C++ with OpenGL, GLSL graphics, employed multiple advanced graphics optimization techniques to get high performance and accurate results in difficult conditions.</li>
				<li>Had weekly meetings with supervisor to discuss and determine the project direction, iterated based on feedback.</li>
			</ul>
		</div>
		<div class="item wip">
			<div class="itemheader">
				<div class="projectname">Project eX0 </div>
				<div class="dates">2007</div>
			</div>
			<ul>
				<li>Implemented in C++ with OpenGL graphics.</li>
				<li>Developed own high­-performance and reliable networking protocol over raw TCP/UDP sockets, which uniquely combined beneficial properties of past networking models.</li>
			</ul>
		</div>
		<div class="sectionheader">Education</div>
		<div class="item">
			<div class="itemheader">
				<div class="schoolname">York University</div>
				<div class="dates">2009 - 2011</div>
			</div>
			<ul>
				<li>Master's Degree, Computer Science.</li>
			</ul>
		</div>
		<div class="item">
			<div class="itemheader">
				<div class="schoolname">York University</div>
				<div class="dates">2004 - 2009</div>
			</div>
			<ul>
				<li>Bachelor's Degree, Specialized Honors Computer Science.</li>
			</ul>
		</div>
		<div class="sectionheader">Knowledge and Skills Highlights</div>
		<div class="item">
			<b>Languages and APIs</b>: Go<span class="fade">, C/C++, Java, C#, </span>OpenGL<span class="fade">, SQL.</span>
		</div>
		<div class="item">
			<b>Software</b>: OS X, Linux, Windows, git, Microsoft Visual Studio, Xcode.
		</div>
	</div>
{{end}}
`))
}
