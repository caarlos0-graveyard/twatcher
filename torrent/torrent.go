package torrent

import (
	"io/ioutil"
	"net/http"
	"os/user"
	"strings"

	log "github.com/Sirupsen/logrus"
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
	log.WithField("title", t.Title).Println("Downloading...")
	response, err := http.Get(t.Href)
	if err != nil {
		log.WithError(err).WithField("title", t.Title).
			Println("Failed to download!")
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithError(err).WithField("title", t.Title).
			Println("Failed to open torrent!")
		return
	}

	if err := ioutil.WriteFile(t.filename(), contents, 0644); err != nil {
		log.WithError(err).WithField("title", t.Title).
			Println("Failed to write torrent to folder!")
		return
	}
}

func (t *Torrent) filename() string {
	filename := strings.Replace(t.Title, " ", ".", -1)
	me, err := user.Current()
	if err != nil {
		log.WithError(err).WithField("title", t.Title).
			Println("Failed to get current user! Will save files in /tmp.")
		return "/tmp/" + filename + ".torrent"
	}
	return me.HomeDir + "/Downloads/" + filename + ".torrent"
}
