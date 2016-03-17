package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env"
	"github.com/caarlos0/tvshows/torrent"
	rss "github.com/jteeuwen/go-pkg-rss"
)

// Config for the poller
type Config struct {
	URL   string   `env:"FEED_URL"`
	Shows []string `env:"SHOWS" envSeparator:"|"`
}

var config Config

func main() {
	config = Config{}
	env.Parse(&config)
	PollFeed(config.URL, 5)
}

// PollFeed updates the feed and download it to ~/Downloads
func PollFeed(uri string, timeout int) {
	feed := rss.New(timeout, true, chanHandler, itemHandler)
	for {
		if err := feed.Fetch(uri, nil); err != nil {
			fmt.Fprintf(os.Stderr, "[e] %s: %s\n", uri, err)
			return
		}
		<-time.After(time.Duration(feed.SecondsTillUpdate() * 1e9))
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	for _, item := range newitems {
		for _, link := range item.Links {
			for _, show := range config.Shows {
				if strings.Contains(link.Href, show) {
					go torrent.NewTorrent(item.Title, link.Href).Download()
				}
			}
		}
	}
}
