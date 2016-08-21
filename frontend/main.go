// +build js

// frontend is a Go package to be compiled with GopherJS. It renders the resume entirely
// on the frontend.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/reactions"
	"github.com/shurcooL/resume"
	"github.com/shurcooL/users"
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	switch readyState := document.ReadyState(); readyState {
	case "loading":
		document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
			go setup()
		})
	case "interactive", "complete":
		setup()
	default:
		panic(fmt.Errorf("internal error: unexpected document.ReadyState value: %v", readyState))
	}
}

func setup() {
	authenticatedUser, err := httpUsers{}.GetAuthenticated(context.TODO())
	if err != nil {
		log.Println(err)
	}

	resume.CurrentUser = authenticatedUser // THINK.
	resume.Reactions = httpReactions{}     // THINK.

	if !document.Body().HasChildNodes() {
		var buf bytes.Buffer
		returnURL := dom.GetWindow().Location().Pathname
		_, err = io.WriteString(&buf, string(htmlg.Render(resume.Header{ReturnURL: returnURL}.Render()...)))
		if err != nil {
			panic(err) // TODO.
		}
		_, err = io.WriteString(&buf, string(htmlg.Render(resume.DmitriShuralyov{}.Render()...)))
		if err != nil {
			panic(err)
		}
		document.Body().SetInnerHTML(buf.String())
	}

	setupReactionsMenu(authenticatedUser.ID != 0)
}

// httpReactions implements reactions.Service remotely over HTTP.
type httpReactions struct{}

// Get reactions for id at uri.
// uri is clean '/'-separated URI. E.g., "example.com/page".
func (httpReactions) Get(ctx context.Context, uri string, id string) ([]reactions.Reaction, error) {
	u := url.URL{Path: "/api/react", RawQuery: url.Values{"reactableURL": {uri}, "reactableID": {id}}.Encode()}
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("did not get acceptable status code: %v body: %q", resp.Status, body)
	}
	var rs []reactions.Reaction
	err = json.NewDecoder(resp.Body).Decode(&rs)
	return rs, err
}

// Toggle a reaction for id at uri.
func (httpReactions) Toggle(ctx context.Context, uri string, id string, tr reactions.ToggleRequest) ([]reactions.Reaction, error) {
	resp, err := http.PostForm("/api/react", url.Values{"reactableURL": {uri}, "reactableID": {id}, "reaction": {string(tr.Reaction)}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("did not get acceptable status code: %v body: %q", resp.Status, body)
	}
	var rs []reactions.Reaction
	err = json.NewDecoder(resp.Body).Decode(&rs)
	return rs, err
}

// httpUsers implements users.Service remotely over HTTP.
type httpUsers struct{}

// GetAuthenticated fetches the currently authenticated user,
// or User{UserSpec: UserSpec{ID: 0}} if there is no authenticated user.
func (httpUsers) GetAuthenticated(ctx context.Context) (users.User, error) {
	resp, err := http.Get("/api/user")
	if err != nil {
		return users.User{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return users.User{}, fmt.Errorf("did not get acceptable status code: %v body: %q", resp.Status, body)
	}
	var u users.User
	err = json.NewDecoder(resp.Body).Decode(&u)
	return u, err
}

// GetAuthenticatedSpec fetches the currently authenticated user specification,
// or UserSpec{ID: 0} if there is no authenticated user.
func (httpUsers) GetAuthenticatedSpec(ctx context.Context) (users.UserSpec, error) {
	return users.UserSpec{}, fmt.Errorf("GetAuthenticatedSpec: not implemented")
}

// Get fetches the specified user.
func (httpUsers) Get(ctx context.Context, user users.UserSpec) (users.User, error) {
	return users.User{}, fmt.Errorf("Get: not implemented")
}

// Edit the authenticated user.
func (httpUsers) Edit(ctx context.Context, er users.EditRequest) (users.User, error) {
	return users.User{}, fmt.Errorf("Edit: not implemented")
}
