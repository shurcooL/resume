// Package component contains individual components that can render themselves as HTML.
package component

import (
	"html/template"

	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Fade is like Text, except the font color is faded out to a light gray shade.
type Fade string

func (t Fade) Render() []*html.Node {
	span := &html.Node{
		Type: html.ElementNode, Data: atom.Span.String(),
		Attr: []html.Attribute{{Key: atom.Style.String(), Val: `color: #969696;`}},
	}
	span.AppendChild(htmlg.Text(string(t)))
	return []*html.Node{span}
}

type Link struct {
	Text string
	Href template.URL
}

func (l Link) Render() []*html.Node {
	a := htmlg.A(l.Text, l.Href)
	a.Attr = append(a.Attr, html.Attribute{Key: atom.Target.String(), Val: "_blank"})
	return []*html.Node{a}
}
