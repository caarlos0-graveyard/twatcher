package torrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	fmt.Printf("Downloading %s...\n", t.Title)
	response, err := http.Get(t.Href)
	if err != nil {
		fmt.Printf("Failed to download %s\n", t.Title)
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read the contents of %s\n", t.Title)
		return
	}

	if err := ioutil.WriteFile(t.filename(), contents, 0644); err != nil {
		fmt.Printf("Failed to save the contents of %s to the torrent file\n", t.Title)
	}
}

func (t *Torrent) filename() string {
	me, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get the current user!\n")
		return ""
	}
	filename := strings.Replace(t.Title, " ", ".", -1)
	return me.HomeDir + "/Downloads/" + filename + ".torrent"
}
