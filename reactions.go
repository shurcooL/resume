package resume

import (
	"context"
	"fmt"
	"log"

	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/reactions"
	"github.com/shurcooL/users"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const ReactableURL = "dmitri.shuralyov.com/resume"

type Reactable struct {
	ID      string
	Content Component
}

func (r Reactable) Render() []*html.Node {
	// TODO: Make this much nicer.
	/*
		<div class="reactable-container" data-reactableID="{{.ReactableID}}">
			{{template "reactions" .Reactions}}{{template "new-reaction" .ID}}
		</div>
	*/
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Div.String(),
		Attr: []html.Attribute{
			{Key: atom.Class.String(), Val: "reactable-container"},
			{Key: "data-reactableID", Val: r.ID},
		},
	}
	reactions, err := Reactions.Get(context.TODO(), ReactableURL, r.ID) // TODO: Parallelize this for better performance.
	if err != nil {
		log.Println(err)
		reactions = nil
	}
	for _, r := range reactions {
		for _, n := range (Reaction{r}).Render() {
			div.AppendChild(n)
		}
	}
	for _, n := range (NewReaction{ReactableID: r.ID}).Render() {
		div.AppendChild(n)
	}

	return append(r.Content.Render(), div)
}

type Reaction struct {
	reactions.Reaction
}

func (r Reaction) Render() []*html.Node {
	// TODO: Make this much nicer.
	/*
		<a class="reaction" href="javascript:" title="{{reactionTooltip .}}" onclick="ToggleReaction(this, event, {{.Reaction | json}});">
			<div class="reaction {{if (not (containsCurrentUser .Users))}}others{{end}}">
				<span class="emoji-outer emoji-sizer">
					<span class="emoji-inner" style="background-position: {{reactionPosition .Reaction}};">
					</span>
				</span>
				<b>{{len .Users}}</b>
			</div>
		</a>
	*/
	innerSpan := &html.Node{
		Type: html.ElementNode, Data: atom.Span.String(),
		Attr: []html.Attribute{
			{Key: atom.Class.String(), Val: "emoji-inner"},
			{Key: atom.Style.String(), Val: fmt.Sprintf("background-position: %s;", reactions.Position(":"+string(r.Reaction.Reaction)+":"))},
		},
	}
	outerSpan := &html.Node{
		Type: html.ElementNode, Data: atom.Span.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: "emoji-outer emoji-sizer"}},
	}
	outerSpan.AppendChild(innerSpan)
	b := htmlg.Strong(fmt.Sprint(len(r.Reaction.Users)))
	divClass := "reaction"
	if !containsCurrentUser(r.Reaction.Users) {
		divClass += " others"
	}
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Div.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: divClass}},
	}
	div.AppendChild(outerSpan)
	div.AppendChild(b)
	a := &html.Node{
		Type: html.ElementNode, Data: atom.A.String(),
		Attr: []html.Attribute{
			{Key: atom.Class.String(), Val: "reaction"},
			{Key: atom.Href.String(), Val: "javascript:"},
			{Key: atom.Title.String(), Val: reactionTooltip(r.Reaction)},
			{Key: atom.Onclick.String(), Val: fmt.Sprintf("ToggleReaction(this, event, '%q');", r.Reaction.Reaction)},
		},
	}
	a.AppendChild(div)
	return []*html.Node{a}
}

type NewReaction struct {
	ReactableID string
}

func (nr NewReaction) Render() []*html.Node {
	// TODO: Make this much nicer.
	/*
		<a href="javascript:" title="React" onclick="ShowReactionMenu(this, event, {{.}});">
			<div class="new-reaction">
				<i class="octicon octicon-smiley"><sup>+</sup></i>
			</div>
		</a>
	*/
	sup := &html.Node{
		Type: html.ElementNode, Data: atom.Sup.String(),
	}
	sup.AppendChild(htmlg.Text("+"))
	i := &html.Node{
		Type: html.ElementNode, Data: atom.I.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: "octicon octicon-smiley"}},
	}
	i.AppendChild(sup)
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Div.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: "new-reaction"}},
	}
	div.AppendChild(i)
	a := &html.Node{
		Type: html.ElementNode, Data: atom.A.String(),
		Attr: []html.Attribute{
			{Key: atom.Class.String(), Val: "new-reaction"},
			{Key: atom.Href.String(), Val: "javascript:"},
			{Key: atom.Title.String(), Val: "React"},
			{Key: atom.Onclick.String(), Val: fmt.Sprintf("ShowReactionMenu(this, event, '%q');", nr.ReactableID)},
		},
	}
	a.AppendChild(div)
	return []*html.Node{a}
}

var CurrentUser users.User // TODO, THINK, HACK.
var Reactions reactions.Service

// THINK.
func containsCurrentUser(users []users.User) bool {
	if CurrentUser.ID == 0 {
		return false
	}
	for _, u := range users {
		if u.ID == CurrentUser.ID {
			return true
		}
	}
	return false
}

func reactionTooltip(reaction reactions.Reaction) string {
	var users string
	for i, u := range reaction.Users {
		if i != 0 {
			if i < len(reaction.Users)-1 {
				users += ", "
			} else {
				users += " and "
			}
		}
		if CurrentUser.ID != 0 && u.ID == CurrentUser.ID {
			if i == 0 {
				users += "You"
			} else {
				users += "you"
			}
		} else {
			users += u.Login
		}
	}
	return fmt.Sprintf("%v reacted with :%v:.", users, reaction.Reaction)
}
