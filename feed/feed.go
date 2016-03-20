package feed

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/tvshows/torrent"
	rss "github.com/jteeuwen/go-pkg-rss"
)

// Feed type
type Feed struct {
	URL   string
	Shows []string
}

// NewFeed instance
func NewFeed(uri string, shows []string) *Feed {
	return &Feed{
		URL:   uri,
		Shows: shows,
	}
}

// Poll updates the feed and download it to ~/Downloads
func (f *Feed) Poll() {
	feed := rss.New(10, true, f.chanHandler, f.itemHandler)
	for {
		fmt.Println("Looking for new episodes...")
		if err := feed.Fetch(f.URL, nil); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to fetch feed: %s: %s\n", f.URL, err)
		}
		<-time.After(time.Duration(5 * time.Minute))
	}
}

func (f *Feed) itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	for _, item := range newitems {
		for _, link := range item.Links {
			for _, show := range f.Shows {
				if strings.Contains(link.Href, show) {
					go torrent.NewTorrent(item.Title, link.Href).Download()
				}
			}
		}
	}
}

func (f *Feed) chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	// does nothing
}
