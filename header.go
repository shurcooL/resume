package resume

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Header struct {
	ReturnURL string
}

func (h Header) Render() []*html.Node {
	// TODO: Make this much nicer.
	/*
		{{if .CurrentUser.ID}}{{with .CurrentUser}}
			<div style="text-align: right; margin-bottom: 20px; height: 18px; font-size: 12px;">
				<a class="topbar-avatar" href="{{.HTMLURL}}" target="_blank" tabindex=-1>
					<img class="topbar-avatar" src="{{.AvatarURL}}" title="Signed in as {{.Login}}.">
				</a>
				PostButton{Action: "/logout", Text: "Sign out", ReturnURL: h.ReturnURL}
			</div>
		{{end}}{{end}}
	*/
	if CurrentUser.ID == 0 {
		return nil
	}

	div := &html.Node{
		Type: html.ElementNode, Data: atom.Div.String(),
		Attr: []html.Attribute{
			{Key: atom.Style.String(), Val: "text-align: right; margin-bottom: 20px; height: 18px; font-size: 12px;"},
		},
	}

	a := &html.Node{
		Type: html.ElementNode, Data: atom.A.String(),
		Attr: []html.Attribute{
			{Key: atom.Class.String(), Val: "topbar-avatar"},
			{Key: atom.Href.String(), Val: string(CurrentUser.HTMLURL)},
			{Key: atom.Target.String(), Val: "_blank"},
			{Key: atom.Tabindex.String(), Val: "-1"},
		},
	}
	a.AppendChild(&html.Node{
		Type: html.ElementNode, Data: atom.Img.String(),
		Attr: []html.Attribute{
			{Key: atom.Class.String(), Val: "topbar-avatar"},
			{Key: atom.Src.String(), Val: string(CurrentUser.AvatarURL)},
			{Key: atom.Title.String(), Val: fmt.Sprintf("Signed in as %s.", CurrentUser.Login)},
		},
	})
	div.AppendChild(a)

	logoutButton := PostButton{Action: "/logout", Text: "Sign out", ReturnURL: h.ReturnURL}
	for _, n := range logoutButton.Render() {
		div.AppendChild(n)
	}

	return []*html.Node{div}
}

type PostButton struct {
	Action    string
	Text      string
	ReturnURL string
}

func (b PostButton) Render() []*html.Node {
	// TODO: Make this much nicer.
	/*
		<form method="post" action="{{.Action}}" style="display: inline-block; margin-bottom: 0;">
			<input class="btn" type="submit" value="{{.Text}}">
			<input type="hidden" name="return" value="{{.ReturnURL}}">
		</form>
	*/
	form := &html.Node{
		Type: html.ElementNode, Data: atom.Form.String(),
		Attr: []html.Attribute{
			{Key: atom.Method.String(), Val: "post"},
			{Key: atom.Action.String(), Val: b.Action},
			{Key: atom.Style.String(), Val: `display: inline-block; margin-bottom: 0;`},
		},
	}
	form.AppendChild(&html.Node{
		Type: html.ElementNode, Data: atom.Input.String(),
		Attr: []html.Attribute{
			{Key: atom.Class.String(), Val: "btn"},
			{Key: atom.Type.String(), Val: "submit"},
			{Key: atom.Value.String(), Val: b.Text},
		},
	})
	form.AppendChild(&html.Node{
		Type: html.ElementNode, Data: atom.Input.String(),
		Attr: []html.Attribute{
			{Key: atom.Type.String(), Val: "hidden"},
			{Key: atom.Name.String(), Val: "return"},
			{Key: atom.Value.String(), Val: b.ReturnURL},
		},
	})
	return []*html.Node{form}
}
