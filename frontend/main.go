// +build js

// frontend is a Go package to be compiled with GopherJS. It renders the resume entirely
// on the frontend.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
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
	authenticatedUser, err := getAuthenticatedUser()
	if err != nil {
		log.Println(err)
	}

	resume.CurrentUser = authenticatedUser // THINK.
	resume.Reactions = httpReactions{}     // THINK.

	if !document.Body().HasChildNodes() {
		var buf bytes.Buffer
		err = t.Execute(&buf, Header{AuthenticatedUser: authenticatedUser})
		if err != nil {
			panic(err)
		}
		_, err = io.WriteString(&buf, string(htmlg.Render(resume.DmitriShuralyov{}.Render()...)))
		if err != nil {
			panic(err)
		}
		document.Body().SetInnerHTML(buf.String())
	}

	setupReactionsMenu(authenticatedUser.ID != 0)
}

// TODO: Should this be a method of an HTTP users.Service implementation (similar to httpReactions)?
func getAuthenticatedUser() (users.User, error) {
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
	if err != nil {
		return users.User{}, err
	}
	return u, nil
}

type Header struct {
	AuthenticatedUser users.User
}

func (Header) Return() template.URL {
	return template.URL(dom.GetWindow().Location().Pathname)
}

var t = template.Must(template.New("").Parse(`
{{- if .AuthenticatedUser.ID}}{{with .AuthenticatedUser}}
	<div style="text-align: right; margin-bottom: 20px; height: 18px; font-size: 12px;">
		<a class="topbar-avatar" href="{{.HTMLURL}}" target="_blank" tabindex=-1
			><img class="topbar-avatar" src="{{.AvatarURL}}" title="Signed in as {{.Login}}."
		></a>
		<form method="post" action="/logout" style="display: inline-block; margin-bottom: 0;"><input class="btn" type="submit" value="Sign out"><input type="hidden" name="return" value="{{$.Return}}"></form>
	</div>
{{end}}{{end -}}
`))

// httpReactions implements reactions.Service remotely over HTTP.
type httpReactions struct {
	BaseURI string
}

// Get reactions for id at uri.
// uri is clean '/'-separated URI. E.g., "example.com/page".
func (hr httpReactions) Get(ctx context.Context, uri string, id string) ([]reactions.Reaction, error) {
	u := url.URL{Path: hr.BaseURI + "/react", RawQuery: url.Values{"reactableURL": {uri}, "reactableID": {id}}.Encode()}
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
func (hr httpReactions) Toggle(ctx context.Context, uri string, id string, tr reactions.ToggleRequest) ([]reactions.Reaction, error) {
	resp, err := http.PostForm(hr.BaseURI+"/react", url.Values{"reactableURL": {uri}, "reactableID": {id}, "reaction": {string(tr.Reaction)}})
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
