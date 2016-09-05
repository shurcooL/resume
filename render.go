package resume

import (
	"context"
	"io"

	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/notifications"
	"github.com/shurcooL/reactions"
	"github.com/shurcooL/users"
)

// RenderBodyInnerHTML renders the inner HTML of the <body> element of the page that displays the resume.
// It's not safe for concurrent use.
func RenderBodyInnerHTML(ctx context.Context, w io.Writer, reactions reactions.Service, notifications notifications.Service, authenticatedUser users.User, returnURL string) error {
	// THINK: Should I work really hard (and add verbosity) to eliminate these package-level variables,
	//        or is it okay to keep them that way?
	reactableReactionsService = reactions
	reactionCurrentUser = authenticatedUser

	// Render the header.
	header := Header{
		notifications: notifications,

		CurrentUser: authenticatedUser,
		ReturnURL:   returnURL,
	}
	err := htmlg.RenderComponentsContext(ctx, w, header)
	if err != nil {
		return err
	}

	// Render the resume contents.
	err = htmlg.RenderComponents(w, DmitriShuralyov{})
	if err != nil {
		return err
	}

	return nil
}
