// +build js

// frontend is a Go package to be compiled with GopherJS. It renders the resume entirely
// on the frontend.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/shurcooL/notifications"
	"github.com/shurcooL/reactions"
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
	reactionsService := httpReactions{}
	authenticatedUser, err := httpUsers{}.GetAuthenticated(context.TODO())
	if err != nil {
		log.Println(err)
		authenticatedUser = users.User{} // THINK: Should it be a fatal error or not? What about on frontend vs backend?
	}

	if !document.Body().HasChildNodes() {
		var buf bytes.Buffer
		returnURL := dom.GetWindow().Location().Pathname + dom.GetWindow().Location().Search
		err = resume.RenderBodyInnerHTML(context.TODO(), &buf, reactionsService, httpNotifications{}, authenticatedUser, returnURL)
		if err != nil {
			log.Println(err)
			return
		}
		document.Body().SetInnerHTML(buf.String())
	}

	setupReactionsMenu(reactionsService, authenticatedUser.ID != 0)
}

// httpReactions implements reactions.Service remotely over HTTP.
type httpReactions struct{}

var _ reactions.Service = httpReactions{}

func (httpReactions) Get(_ context.Context, uri string, id string) ([]reactions.Reaction, error) {
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
func (httpReactions) Toggle(_ context.Context, uri string, id string, tr reactions.ToggleRequest) ([]reactions.Reaction, error) {
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

// httpNotifications implements notifications.Service remotely over HTTP.
type httpNotifications struct{}

var _ notifications.Service = httpNotifications{}

func (httpNotifications) Count(_ context.Context, opt interface{}) (uint64, error) {
	resp, err := http.Get("/api/notifications/count")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return 0, fmt.Errorf("did not get acceptable status code: %v body: %q", resp.Status, body)
	}
	var u uint64
	err = json.NewDecoder(resp.Body).Decode(&u)
	return u, err
}
func (httpNotifications) List(_ context.Context, opt notifications.ListOptions) (notifications.Notifications, error) {
	return nil, fmt.Errorf("List: not implemented")
}
func (httpNotifications) MarkAllRead(_ context.Context, repo notifications.RepoSpec) error {
	return fmt.Errorf("MarkAllRead: not implemented")
}
func (httpNotifications) Subscribe(_ context.Context, appID string, repo notifications.RepoSpec, threadID uint64, subscribers []users.UserSpec) error {
	return fmt.Errorf("Subscribe: not implemented")
}
func (httpNotifications) MarkRead(_ context.Context, appID string, repo notifications.RepoSpec, threadID uint64) error {
	return fmt.Errorf("MarkRead: not implemented")
}
func (httpNotifications) Notify(_ context.Context, appID string, repo notifications.RepoSpec, threadID uint64, nr notifications.NotificationRequest) error {
	return fmt.Errorf("Notify: not implemented")
}

// httpUsers implements users.Service remotely over HTTP.
type httpUsers struct{}

var _ users.Service = httpUsers{}

func (httpUsers) GetAuthenticated(_ context.Context) (users.User, error) {
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
func (httpUsers) GetAuthenticatedSpec(_ context.Context) (users.UserSpec, error) {
	resp, err := http.Get("/api/userspec")
	if err != nil {
		return users.UserSpec{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return users.UserSpec{}, fmt.Errorf("did not get acceptable status code: %v body: %q", resp.Status, body)
	}
	var us users.UserSpec
	err = json.NewDecoder(resp.Body).Decode(&us)
	return us, err
}
func (httpUsers) Get(_ context.Context, user users.UserSpec) (users.User, error) {
	return users.User{}, fmt.Errorf("Get: not implemented")
}
func (httpUsers) Edit(_ context.Context, er users.EditRequest) (users.User, error) {
	return users.User{}, fmt.Errorf("Edit: not implemented")
}
