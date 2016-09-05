// +build js

// frontend renders the resume entirely on the frontend.
// It is a Go package meant to be compiled with GopherJS.
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/shurcooL/home/http"
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
	reactionsService := http.Reactions{}
	authenticatedUser, err := http.Users{}.GetAuthenticated(context.TODO())
	if err != nil {
		log.Println(err)
		authenticatedUser = users.User{} // THINK: Should it be a fatal error or not? What about on frontend vs backend?
	}

	if !document.Body().HasChildNodes() {
		var buf bytes.Buffer
		returnURL := dom.GetWindow().Location().Pathname + dom.GetWindow().Location().Search
		err = resume.RenderBodyInnerHTML(context.TODO(), &buf, reactionsService, http.Notifications{}, authenticatedUser, returnURL)
		if err != nil {
			log.Println(err)
			return
		}
		document.Body().SetInnerHTML(buf.String())
	}

	setupReactionsMenu(reactionsService, authenticatedUser)
}
