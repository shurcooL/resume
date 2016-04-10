package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/shurcooL/users"
	"honnef.co/go/js/dom"
)

func getAuthenticatedUser() (*users.User, error) {
	resp, err := http.Get("/user")
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

type Auth struct {
	AuthenticatedUser *users.User
}

func (Auth) Return() template.URL {
	return template.URL(dom.GetWindow().Location().Pathname)
}
