// frontend renders the resume entirely on the frontend.
// It is a Go package meant to be compiled with GOARCH=js
// and executed in a browser, where the DOM is available.
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/shurcooL/frontend/reactionsmenu"
	homehttp "github.com/shurcooL/home/http"
	"github.com/shurcooL/notificationsapp/httpclient"
	"github.com/shurcooL/resume"
	"github.com/shurcooL/users"
	"golang.org/x/oauth2"
	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	switch readyState := document.ReadyState(); readyState {
	case "loading":
		document.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			go setup()
		})
	case "interactive", "complete":
		setup()
	default:
		panic(fmt.Errorf("internal error: unexpected document.ReadyState value: %v", readyState))
	}
}

func setup() {
	httpClient := httpClient()

	reactionsService := homehttp.Reactions{}
	notificationsService := httpclient.NewNotifications(httpClient, "", "")
	authenticatedUser, err := homehttp.Users{}.GetAuthenticated(context.TODO())
	if err != nil {
		log.Println(err)
		authenticatedUser = users.User{} // THINK: Should it be a fatal error or not? What about on frontend vs backend?
	}

	query, _ := url.ParseQuery(strings.TrimPrefix(dom.GetWindow().Location().Search, "?"))
	prerender, _ := strconv.ParseBool(query.Get("prerender"))
	if !prerender {
		var buf bytes.Buffer
		returnURL := dom.GetWindow().Location().Pathname + dom.GetWindow().Location().Search
		err = resume.RenderBodyInnerHTML(context.TODO(), &buf, reactionsService, notificationsService, authenticatedUser, returnURL)
		if err != nil {
			log.Println(err)
			return
		}
		document.Body().SetInnerHTML(buf.String())
	}

	reactionsmenu.Setup(resume.ReactableURL, reactionsService, authenticatedUser)
}

// httpClient gives an *http.Client for making API requests.
func httpClient() *http.Client {
	document := &http.Request{Header: http.Header{"Cookie": {document.Cookie()}}}
	if accessToken, err := document.Cookie("accessToken"); err == nil {
		// Authenticated client.
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: accessToken.Value},
		)
		return oauth2.NewClient(context.Background(), src)
	}
	// Not authenticated client.
	return http.DefaultClient
}
