package main

import (
	"fmt"
	"html/template"

	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Component interface {
	Render() ([]*html.Node, error)
}

type Text string

func (t Text) Render() ([]*html.Node, error) {
	return []*html.Node{htmlg.Text(string(t))}, nil
}

type Link struct {
	Text string
	Href template.URL
}

func (l Link) Render() ([]*html.Node, error) {
	a := htmlg.A(l.Text, l.Href)
	a.Attr = append(a.Attr, html.Attribute{Key: atom.Target.String(), Val: "_blank"})
	return []*html.Node{a}, nil
}

func join(a ...interface{}) Component {
	var nodes []*html.Node
	for _, v := range a {
		switch v := v.(type) {
		case Component:
			n, err := v.Render()
			if err != nil {
				return composite{err: err}
			}
			nodes = append(nodes, n...)
		case string:
			nodes = append(nodes, htmlg.Text(v))
		default:
			return composite{err: fmt.Errorf("unsupported type: %T", v)}
		}
	}
	return composite{nodes: nodes}
}

type composite struct {
	nodes []*html.Node
	err   error
}

func (c composite) Render() ([]*html.Node, error) { return c.nodes, c.err }
