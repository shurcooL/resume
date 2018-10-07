package resume_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/shurcooL/htmlg"
	"github.com/shurcooL/reactions"
	"github.com/shurcooL/resume"
	"github.com/shurcooL/users"
)

var updateFlag = flag.Bool("update", false, "Update golden files.")

// TestDmitriShuralyov validates that rendering resume.DmitriShuralyov produces expected HTML.
func TestDmitriShuralyov(t *testing.T) {
	var (
		shurcool = users.User{
			UserSpec: users.UserSpec{ID: 1924134, Domain: "github.com"},
			Name:     "Dmitri Shuralyov",
			Email:    "dmitri@shuralyov.com",
		}

		mockTime = time.Date(2018, time.August, 26, 9, 41, 0, 0, time.UTC)

		alice = users.User{UserSpec: users.UserSpec{ID: 1, Domain: "example.org"}, Login: "Alice"}
		bob   = users.User{UserSpec: users.UserSpec{ID: 2, Domain: "example.org"}, Login: "Bob"}

		mockReactions = map[string][]reactions.Reaction{
			"Go": {{
				Reaction: "smile",
				Users:    []users.User{alice, bob},
			}, {
				Reaction: "balloon",
				Users:    []users.User{bob},
			}},
		}
	)

	var buf bytes.Buffer
	err := htmlg.RenderComponents(&buf, resume.DmitriShuralyov(shurcool, mockTime, mockReactions, alice))
	if err != nil {
		t.Fatal(err)
	}
	got := buf.Bytes()
	if *updateFlag {
		err := ioutil.WriteFile(filepath.Join("testdata", "resume.html"), got, 0644)
		if err != nil {
			t.Fatal(err)
		}
		return
	}

	want, err := ioutil.ReadFile(filepath.Join("testdata", "resume.html"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Error("rendering resume.DmitriShuralyov produced HTML that doesn't match 'testdata/resume.html'")
	}
}
