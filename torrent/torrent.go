package torrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

// Torrent is a Title/Href combo
type Torrent struct {
	Title string
	Href  string
}

// NewTorrent instance
func NewTorrent(title, href string) *Torrent {
	return &Torrent{
		Title: title,
		Href:  href,
	}
}

// Download the torrent to ~/Downloads
func (t *Torrent) Download() {
	fmt.Fprintf(os.Stderr, "Downloading %s...\n", t.Title)
	response, err := http.Get(t.Href)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to download %s\n", t.Title)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to read the contents of %s\n",
			t.Title,
		)
	}

	if err := ioutil.WriteFile(t.filename(), contents, 0644); err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to save the contents of %s to the torrent file\n",
			t.Title,
		)
	}
}

func (t *Torrent) filename() string {
	filename := strings.Replace(t.Title, " ", ".", -1)
	me, err := user.Current()
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Failed to get the current user! Will save files in /tmp.\n",
		)
		return "/tmp/" + filename + ".torrent"
	}
	return me.HomeDir + "/Downloads/" + filename + ".torrent"
}
