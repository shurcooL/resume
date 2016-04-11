package resume

import (
	"fmt"
	"html/template"

	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Component interface {
	Render() []*html.Node
}

type Text string

func (t Text) Render() []*html.Node {
	return []*html.Node{htmlg.Text(string(t))}
}

type fade string

func (t fade) Render() []*html.Node {
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

// join components and strings into a single Component.
func join(a ...interface{}) List {
	var list List
	for _, v := range a {
		switch v := v.(type) {
		case Component:
			list = append(list, v)
		case string:
			list = append(list, Text(v))
		default:
			panic(fmt.Errorf("join: unsupported type: %T", v))
		}
	}
	return list
}

type List []Component

func (l List) Render() []*html.Node {
	var nodes []*html.Node
	for _, c := range l {
		nodes = append(nodes, c.Render()...)
	}
	return nodes
}
