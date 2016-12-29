# TorrentWatcher [![Build Status](https://travis-ci.org/caarlos0/twatcher.svg?branch=master)](https://travis-ci.org/caarlos0/twatcher)

Agent to automagically download torrent files to `~/Downloads` :bomb:

## Install

```console
brew tap caarlos0/formulae
brew install twatcher
```

## Usage

```console
$ twatcher \
  --feed YOUR_TORRENT_FEED_URL \
  --name Vikings \
  --name "Game of Thrones" \
  --filter "1080p.HDTV"
```

### crontab

You can make it run every X minutes by putting it on your crontab.

```crontab
0/10 * * * * /usr/local/bin/twatcher --feed YOUR_TORRENT_FEED_URL --filter "1080p.HDTV" --name Vikings --name "Game of Thrones" > ~/Library/Logs/twatcher.log 2>&1
```
