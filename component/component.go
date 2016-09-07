// Package component contains individual components that can render themselves as HTML.
package component

import (
	"fmt"
	"html/template"

	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Text string

func (t Text) Render() []*html.Node {
	return []*html.Node{htmlg.Text(string(t))}
}

// Fade is like Text, except the font color is faded out to a light gray shade.
type Fade string

func (t Fade) Render() []*html.Node {
	return []*html.Node{htmlg.SpanClass("fade", htmlg.Text(string(t)))}
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

// Join components and strings into a single component.
func Join(a ...interface{}) List {
	var list List
	for _, v := range a {
		switch v := v.(type) {
		case htmlg.Component:
			list = append(list, v)
		case string:
			list = append(list, Text(v))
		default:
			panic(fmt.Errorf("Join: unsupported type: %T", v))
		}
	}
	return list
}

type List []htmlg.Component

func (l List) Render() []*html.Node {
	var nodes []*html.Node
	for _, c := range l {
		nodes = append(nodes, c.Render()...)
	}
	return nodes
}
