// Package component contains individual components that can render themselves as HTML.
package component

import (
	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Section is a section of a resume. For example, "Experience" or "Education".
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
		ns = append(ns, i.Render()...)
	}
	return ns
}

// Item is a single item within a resume section. For example, a particular workplace or school.
type Item struct {
	Title    string
	Subtitle string
	Dates    htmlg.Component
	Lines    []htmlg.Component

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

// Fade is like Text, except the font color is faded out to a light gray shade.
type Fade string

func (t Fade) Render() []*html.Node {
	span := &html.Node{
		Type: html.ElementNode, Data: atom.Span.String(),
		Attr:       []html.Attribute{{Key: atom.Style.String(), Val: `color: #969696;`}},
		FirstChild: htmlg.Text(string(t)),
	}
	return []*html.Node{span}
}
