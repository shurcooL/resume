package resume

import "golang.org/x/net/html"

type Markup interface {
	Apply(n *html.Node)
}

type Attribute struct {
	Namespace, Key, Val string
}

func (a Attribute) Apply(n *html.Node) {
	if n.Type != html.ElementNode {
		panic("invalid node type")
	}
	n.Attr = append(n.Attr, html.Attribute{Namespace: a.Namespace, Key: a.Key, Val: a.Val})
}
