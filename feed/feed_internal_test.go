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
		"https://www.blah.com/Marvels.Agents.of.S.H.I.E.L.D.S01.E02.720p.HDTV.torrent",
		"Marvels Agents of SHIELD",
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
		if set.matches != NewFeed("", set.filter, []string{}).matches(set.url, set.name) {
			t.Error("matches", set)
		}
	}
}
