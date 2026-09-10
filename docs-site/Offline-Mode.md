# Offline Mode

Offline Mode lets you browse and watch/read your library without an internet connection, using cached metadata and locally-downloaded assets.

## Enabling

Toggle **Offline Mode** from Settings, or set `offline = true` under `[server]` in `config.toml` to start the server in offline mode directly.

## What works offline

- Browsing your previously-synced library and metadata
- Playing local files
- Reading previously-downloaded manga chapters (see [Manga](Manga.md))

## What doesn't work offline

- AniList sync (list updates queue and sync once you're back online)
- Torrent/debrid/online streaming — these all require a live connection
- Extension updates

## Manual offline toggle (mobile)

Tori Tenji has its own manual offline switch, separate from the server's offline setting, useful when the app can reach a cached state but the server itself is unreachable (e.g. no signal, server temporarily down).

## Related

- [Manga](Manga.md)
- [Configuration](Configuration.md)
