// Package component contains individual components that can render themselves as HTML.
package component

import (
	"github.com/shurcooL/htmlg"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

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
