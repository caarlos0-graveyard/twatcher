# TorrentWatcher

Agent to automagically download torrent files to ~/Downloads :bomb:

## Usage

```console
$ go get github.com/caarlos0/twatcher
$ twatcher \
  --feed YOUR_TORRENT_FEED_URL \
  --name Vikings \
  --name "Game.of.Thrones" \
  --filter "1080p"
```

### As an OSX daemon

If you want it to run all the time - which makes sense, you can
use add it as an OSX Daemon:

```console
$ cp twatcher.plist{.example,}

# Make your own changes:
$ $EDITOR twatcher.plist

# Make it run!
$ ./update

# Make sure it is running:
$ ps aux | grep twatcher
```
