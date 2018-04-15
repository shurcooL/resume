// Package component contains individual components that can render themselves as HTML.
package component

import (
	"github.com/shurcooL/component"
	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Name is a component for displaying the person's name on top of a resume.
type Name struct {
	Name string
}

func (n Name) Render() []*html.Node {
	div := htmlg.DivClass("name",
		htmlg.Text(n.Name),
	)
	return []*html.Node{div}
}

// ContactInfo is a component for displaying the persons's contact information on top of a resume.
type ContactInfo struct {
	GitHub htmlg.Component
	Email  htmlg.Component
}

func (c ContactInfo) Render() []*html.Node {
	div := htmlg.DivClass("contactinfo", component.Join(
		c.GitHub, " · ", c.Email,
	).Render()...)
	return []*html.Node{div}
}

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

	// WIP controls whether the item is considered a work-in-progress,
	// and is therefore omitted during production use.
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

	ul := htmlg.UL()
	for _, l := range i.Lines {
		li := htmlg.LI(l.Render()...)
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
