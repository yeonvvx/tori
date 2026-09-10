# Extensions

Extensions add sources to Tori — torrent search providers, online-streaming sources, and manga sources — through an in-app marketplace.

## Installing

Go to **Settings → Extensions → Marketplace**, browse or search, and install. Installed extensions appear under **Settings → Extensions → Installed**, where they can be enabled/disabled or configured.

## Extension types

- **Torrent providers** — power torrent search for Auto Downloader and manual torrent search
- **Online streaming sources** — provide streamable sources for anime
- **Manga sources** — provide chapters for series not in your local manga library

## Third-party nature

Extensions are community-made and unaffiliated with the Tori project itself. They may be removed from the marketplace if found to violate copyright law — installing and using any extension is at your own responsibility.

## Writing your own

Tori's extension system is scriptable (JavaScript, via an embedded Goja VM). See the [DEVELOPMENT_AND_BUILD.md](https://github.com/yeonvvx/tori/blob/main/DEVELOPMENT_AND_BUILD.md) guide in the repo for building/testing extensions locally.

## Related

- [Streaming](Streaming.md)
- [Auto Downloader](Auto-Downloader.md)
- [Manga](Manga.md)
