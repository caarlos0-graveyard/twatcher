# tvshows

Agent to automagically download tv shows torrent files to ~/Downloads :bomb:

## Usage

```console
$ go get github.com/caarlos0/tvshows
$ tvshows --feed YOUR_TORRENT_FEED_URL \
  --show Vikings \
  --show "Marvels.Agents.of.S.H.I.E.L.D" \
  --show Gotham \
  --show "Game.of.Thrones"
```

### As an OSX daemon

If you want it to run all the time - which makes sense, by the way, you can
use the OS X Daemon feature:

```console
$ cp tvshows.plist{.example,}

# Make your own changes:
$ $EDITOR tvshows.plist

# Make it run!
$ ./update

# Make sure it is running:
$ ps aux | grep tvshow
```
