package component

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

// ReactionsBar is a component next to anything that can be reacted to, with reactable ID.
// It displays all reactions for that reactable ID, and a NewReaction component for adding new reactions.
type ReactionsBar struct {
	Reactions    reactions.Service
	ReactableURL string
	CurrentUser  users.User
	ID           string // ID is the reactable ID.
}

func (r ReactionsBar) RenderContext(ctx context.Context) []*html.Node {
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
	reactions, err := r.Reactions.Get(ctx, r.ReactableURL, r.ID) // TODO: Parallelize this for better performance.
	if err != nil {
		log.Println(err)
		reactions = nil
	}
	spacingAfter := prioritizeThumbsUpDown(reactions)
	for i, reaction := range reactions {
		for _, n := range (Reaction{Reaction: reaction, CurrentUser: r.CurrentUser}).Render() {
			div.AppendChild(n)
		}
		if i == spacingAfter {
			div.AppendChild(&html.Node{
				Type: html.ElementNode, Data: atom.Span.String(),
				Attr: []html.Attribute{{Key: atom.Style.String(), Val: "margin-right: 4px;"}},
			})
		}
	}
	for _, n := range (NewReaction{ReactableID: r.ID}).Render() {
		div.AppendChild(n)
	}
	return []*html.Node{div}
}
func (r ReactionsBar) Render() []*html.Node {
	return r.RenderContext(context.TODO())
}

// prioritizeThumbsUpDown bubbles +1 and -1 reactions to the front. It returns
// an index after which spacing should be inserted to visually separate
// +1 and -1 reactions from the rest, or -1 if no need.
func prioritizeThumbsUpDown(reactions []reactions.Reaction) (spacingAfter int) {
	spacingAfter = -1
	for i, reaction := range reactions {
		if reaction.Reaction == "+1" {
			for ; i > 0; i-- {
				reactions[i-1], reactions[i] = reactions[i], reactions[i-1]
			}
			spacingAfter++
			break
		}
	}
	for i, reaction := range reactions {
		if reaction.Reaction == "-1" {
			for ; i > 1; i-- {
				reactions[i-1], reactions[i] = reactions[i], reactions[i-1]
			}
			spacingAfter++
			break
		}
	}
	if spacingAfter == len(reactions)-1 {
		spacingAfter = -1 // No need for spacing if there are no other reactions after +1 and -1.
	}
	return spacingAfter
}

// Reaction is a component for displaying a single Reaction, as seen by CurrentUser.
type Reaction struct {
	Reaction    reactions.Reaction
	CurrentUser users.User
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
				<strong>{{len .Users}}</strong>
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
	strong := htmlg.Strong(fmt.Sprint(len(r.Reaction.Users)))
	divClass := "reaction"
	if !r.containsCurrentUser(r.Reaction.Users) {
		divClass += " others"
	}
	div := &html.Node{
		Type: html.ElementNode, Data: atom.Div.String(),
		Attr: []html.Attribute{{Key: atom.Class.String(), Val: divClass}},
	}
	div.AppendChild(outerSpan)
	div.AppendChild(strong)
	a := &html.Node{
		Type: html.ElementNode, Data: atom.A.String(),
		Attr: []html.Attribute{
			{Key: atom.Class.String(), Val: "reaction"},
			{Key: atom.Href.String(), Val: "javascript:"},
			{Key: atom.Title.String(), Val: r.reactionTooltip(r.Reaction)},
			{Key: atom.Onclick.String(), Val: fmt.Sprintf("ToggleReaction(this, event, '%q');", r.Reaction.Reaction)},
		},
	}
	a.AppendChild(div)
	return []*html.Node{a}
}

func (r Reaction) containsCurrentUser(users []users.User) bool {
	if r.CurrentUser.ID == 0 {
		return false
	}
	for _, u := range users {
		if u.ID == r.CurrentUser.ID {
			return true
		}
	}
	return false
}

func (r Reaction) reactionTooltip(reaction reactions.Reaction) string {
	var users string
	for i, u := range reaction.Users {
		if i != 0 {
			if i < len(reaction.Users)-1 {
				users += ", "
			} else {
				users += " and "
			}
		}
		if r.CurrentUser.ID != 0 && u.ID == r.CurrentUser.ID {
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

// NewReaction is a component for adding new reactions to a Reactable with ReactableID id.
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
