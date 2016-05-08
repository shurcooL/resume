// +build js

// frontend is a Go package to be compiled with GopherJS. It renders the resume entirely
// on the frontend.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shurcooL/resume"
	"github.com/shurcooL/users"
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
	resume.CurrentUser = authenticatedUser // THINK, HACK.

	var buf bytes.Buffer
	err = t.Execute(&buf, Header{AuthenticatedUser: authenticatedUser})
	if err != nil {
		panic(err)
	}
	err = resume.T.ExecuteTemplate(&buf, "body", resume.DmitriShuralyov{})
	if err != nil {
		panic(err)
	}
	document.Body().SetInnerHTML(buf.String())

	setupReactionsMenu(authenticatedUser != nil)
}

func getAuthenticatedUser() (*users.User, error) {
	resp, err := http.Get("/api/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("did not get acceptable status code: %v body: %q", resp.Status, body)
	}
	var u = new(users.User)
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

type Header struct {
	AuthenticatedUser *users.User
}

func (Header) Return() template.URL {
	return template.URL(dom.GetWindow().Location().Pathname)
}

var t = template.Must(template.New("").Parse(`
{{with .AuthenticatedUser}}
	<div style="text-align: right; margin-bottom: 20px; height: 18px; font-size: 12px;">
		<a class="topbar-avatar" href="{{.HTMLURL}}" target="_blank" tabindex=-1
			><img class="topbar-avatar" src="{{.AvatarURL}}" title="Signed in as {{.Login}}."
		></a>
		<form method="post" action="/logout" style="display: inline-block; margin-bottom: 0;"><input class="btn" type="submit" value="Sign out"><input type="hidden" name="return" value="{{$.Return}}"></form>
	</div>
{{end}}
`))
