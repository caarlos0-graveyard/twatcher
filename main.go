package main

import (
	"github.com/caarlos0/env"
	"github.com/caarlos0/tvshows/feed"
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
	feed.NewFeed(config.URL, config.Shows).Poll()
}
