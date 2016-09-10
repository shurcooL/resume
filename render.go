package resume

import (
	"context"
	"io"

	"github.com/shurcooL/home/component"
	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/notifications"
	"github.com/shurcooL/reactions"
	"github.com/shurcooL/users"
)

// RenderBodyInnerHTML renders the inner HTML of the <body> element of the page that displays the resume.
// It's safe for concurrent use.
func RenderBodyInnerHTML(ctx context.Context, w io.Writer, reactions reactions.Service, notifications notifications.Service, authenticatedUser users.User, returnURL string) error {
	// Render the header.
	header := component.Header{
		CurrentUser:   authenticatedUser,
		ReturnURL:     returnURL,
		Notifications: notifications,
	}
	err := htmlg.RenderComponentsContext(ctx, w, header)
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

	return nil
}
