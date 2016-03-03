package main

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// span returns a span element <span>{{range .nodes}}{{.}}{{end}}</span>.
func span(nodes ...*html.Node) *html.Node {
	span := &html.Node{
		Type: html.ElementNode, Data: atom.Span.String(),
	}
	for _, n := range nodes {
		span.AppendChild(n)
	}
	return span
}
