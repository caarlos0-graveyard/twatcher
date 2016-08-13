package feed

import "testing"

var resultSet = []struct {
	url     string
	name    string
	filter  string
	matches bool
}{
	{
		"https://www.blah.com/Pokemon.S19E26.DUBBED.720p.HDTV.x264-W4F.torrent",
		"Pokemon",
		"720p",
		true,
	},
	{
		"https://www.blah.com/Pokemon.S19E26.DUBBED.720p.HDTV.x264-W4F.torrent",
		"Pokemon.s19",
		"720p",
		true,
	},
	{
		"https://www.blah.com/Pokemon.S19E26.DUBBED.720p.HDTV.x264-W4F.torrent",
		"Pokemon s19",
		"720p",
		true,
	},
	{
		"https://www.blah.com/Pokemon.S19E26.DUBBED.720p.HDTV.x264-W4F.torrent",
		"Pokemon S19",
		"720p",
		true,
	},
	{
		"https://www.blah.com/Pokemon.S19E26.DUBBED.720p.HDTV.x264-W4F.torrent",
		"Pokemon",
		"1080p",
		false,
	},
	{
		"https://www.blah.com/Pokemon.S19E26.DUBBED.720p.HDTV.x264-W4F.torrent",
		"Game of thrones",
		"1080p.HDTV",
		false,
	},
}

func TestMatches(t *testing.T) {
	for _, set := range resultSet {
		if set.matches != NewFeed("", set.filter, []string{}).match(set.url, set.name) {
			t.Error("matches", set)
		}
	}
}
