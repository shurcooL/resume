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
	var list List
	for _, v := range a {
		switch v := v.(type) {
		case Component:
			list = append(list, v)
		case string:
			list = append(list, Text(v))
		default:
			return Error{fmt.Errorf("join: unsupported type: %T", v)}
		}
	}
	return list
}

type List []Component

func (l List) Render() ([]*html.Node, error) {
	var nodes []*html.Node
	for _, c := range l {
		n, err := c.Render()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, n...)
	}
	return nodes, nil
}

type Error struct{ err error }

func (e Error) Render() ([]*html.Node, error) { return nil, e.err }
