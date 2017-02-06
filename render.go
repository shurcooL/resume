package resume

import (
	"context"
	"io"

	homecomponent "github.com/shurcooL/home/component"
	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/notifications"
	"github.com/shurcooL/reactions"
	"github.com/shurcooL/users"
)

// RenderBodyInnerHTML renders the inner HTML of the <body> element of the page that displays the resume.
// It's safe for concurrent use.
func RenderBodyInnerHTML(ctx context.Context, w io.Writer, reactions reactions.Service, notifications notifications.Service, authenticatedUser users.User, returnURL string) error {
	nc, err := notifications.Count(ctx, nil)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, `<div style="max-width: 800px; margin: 0 auto 100px auto;">`)
	if err != nil {
		return err
	}

	// Render the header.
	header := homecomponent.Header{
		CurrentUser:       authenticatedUser,
		NotificationCount: nc,
		ReturnURL:         returnURL,
	}
	err = htmlg.RenderComponents(w, header)
	if err != nil {
		return err
	}

	// Render the resume contents.
	resume := DmitriShuralyov{
		Reactions:   reactions,
		CurrentUser: authenticatedUser,
	}
	err = htmlg.RenderComponents(w, resume)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, `</div>`)
	return err
}
