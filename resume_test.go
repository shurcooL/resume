package resume_test

import (
	"bytes"
	"context"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/shurcooL/notifications"
	"github.com/shurcooL/reactions"
	"github.com/shurcooL/resume"
	"github.com/shurcooL/users"
)

var updateFlag = flag.Bool("update", false, "Update golden files.")

var (
	alice = users.User{UserSpec: users.UserSpec{ID: 1}, Login: "Alice"}
	bob   = users.User{UserSpec: users.UserSpec{ID: 2}, Login: "Bob"}
)

// TestBodyInnerHTML validates that resume.RenderBodyInnerHTML renders the body inner HTML as expected.
func TestBodyInnerHTML(t *testing.T) {
	want, err := ioutil.ReadFile(filepath.Join("testdata", "body-inner.html"))
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	resume.RenderBodyInnerHTML(context.TODO(), &buf, mockReactions{}, mockNotifications{}, alice, "/")
	got := buf.Bytes()
	if *updateFlag {
		err := ioutil.WriteFile(filepath.Join("testdata", "body-inner.html"), got, 0644)
		if err != nil {
			t.Fatal(err)
		}
		return
	}

	if !bytes.Equal(got, want) {
		t.Error("resume.RenderBodyInnerHTML produced output that doesn't match 'testdata/body-inner.html'")
	}
}

type mockReactions struct{ reactions.Service }

func (mockReactions) Get(_ context.Context, uri string, id string) ([]reactions.Reaction, error) {
	return []reactions.Reaction{{
		Reaction: "smile",
		Users:    []users.User{alice, bob},
	}, {
		Reaction: "balloon",
		Users:    []users.User{bob},
	}}, nil
}

type mockNotifications struct{ notifications.Service }

func (mockNotifications) Count(_ context.Context, opt interface{}) (uint64, error) { return 0, nil }
