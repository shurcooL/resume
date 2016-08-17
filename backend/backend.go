// Package backend helps with rendering the resume on the backend.
package backend

import (
	"html/template"

	"github.com/shurcooL/users"
)

type Header struct {
	AuthenticatedUser users.User
	ReturnURL         string
}

func (h Header) Return() template.URL {
	return template.URL(h.ReturnURL)
}

var T = template.Must(template.New("").Parse(`
{{- if .AuthenticatedUser.ID}}{{with .AuthenticatedUser}}
	<div style="text-align: right; margin-bottom: 20px; height: 18px; font-size: 12px;">
		<a class="topbar-avatar" href="{{.HTMLURL}}" target="_blank" tabindex=-1
			><img class="topbar-avatar" src="{{.AvatarURL}}" title="Signed in as {{.Login}}."
		></a>
		<form method="post" action="/logout" style="display: inline-block; margin-bottom: 0;"><input class="btn" type="submit" value="Sign out"><input type="hidden" name="return" value="{{$.Return}}"></form>
	</div>
{{end}}{{end -}}
`))
