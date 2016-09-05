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

// ReactableURL is the URL for reactionable items on this resume.
const ReactableURL = "dmitri.shuralyov.com/resume"

// Reactable is a wrapper component for any Content that can be reacted to.
// ID is the reactable ID.
type Reactable struct {
	ID      string
	Content Component
}

// THINK: Should I work really hard (and add verbosity) to eliminate this package-level variable,
//        or is it okay to keep it this way?

// reactableReactionsService is a reactions.Service used by Reactable.Render().
// It's set before the Reactable components are rendered, and not modified during rendering.
var reactableReactionsService reactions.Service

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
	reactions, err := reactableReactionsService.Get(context.TODO(), ReactableURL, r.ID) // TODO: Parallelize this for better performance.
	if err != nil {
		log.Println(err)
		reactions = nil
	}
	for _, reaction := range reactions {
		for _, n := range (Reaction{reaction: reaction}).Render() {
			div.AppendChild(n)
		}
	}
	for _, n := range (NewReaction{ReactableID: r.ID}).Render() {
		div.AppendChild(n)
	}

	return append(r.Content.Render(), div)
}

// Reaction is a component for displaying a single reaction.
type Reaction struct {
	Reaction reactions.Reaction
}

// THINK: Should I work really hard (and add verbosity) to eliminate this package-level variable,
//        or is it okay to keep it this way?

// reactionCurrentUser is the current user used by Reaction.Render().
// It's set before the Reaction components are rendered, and not modified during rendering.
var reactionCurrentUser users.User

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
	if !r.containsCurrentUser(r.Reaction.Users) {
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
			{Key: atom.Title.String(), Val: r.reactionTooltip(r.Reaction)},
			{Key: atom.Onclick.String(), Val: fmt.Sprintf("ToggleReaction(this, event, '%q');", r.Reaction.Reaction)},
		},
	}
	a.AppendChild(div)
	return []*html.Node{a}
}

// THINK.
func (Reaction) containsCurrentUser(users []users.User) bool {
	if reactionCurrentUser.ID == 0 {
		return false
	}
	for _, u := range users {
		if u.ID == reactionCurrentUser.ID {
			return true
		}
	}
	return false
}

func (Reaction) reactionTooltip(reaction reactions.Reaction) string {
	var users string
	for i, u := range reaction.Users {
		if i != 0 {
			if i < len(reaction.Users)-1 {
				users += ", "
			} else {
				users += " and "
			}
		}
		if reactionCurrentUser.ID != 0 && u.ID == reactionCurrentUser.ID {
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
