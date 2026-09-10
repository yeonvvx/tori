# Desktop Streaming Setup

This guide covers using Tori without keeping a local library — streaming torrents directly, using a debrid service, or watching via online-streaming extensions.

## Torrent streaming

Tori can stream a torrent directly to the player without waiting for the full download to finish.

1. Go to **Settings → Torrent Streaming** and enable it.
2. Choose a backend:
   - **Built-in (BitTorrent)** — streams directly, no third-party account needed.
   - **TorBox / Real-Debrid / AllDebrid / Premiumize** — enter your API key/token for the provider under its section. Debrid services generally give faster, more consistent streams than raw BitTorrent.
3. Install a torrent-search extension from the [Extension Marketplace](Extensions.md) so Tori has sources to search.

## Downloading via debrid/torrent clients

If you'd rather download than direct-stream, connect a downloader under **Settings → Torrent Client**:

- qBittorrent / Transmission (self-hosted)
- Torbox, Real-Debrid, AllDebrid, or Premiumize (cloud-based)

## Online streaming extensions

Under **Settings → Extensions**, browse the marketplace for online-streaming source extensions. Once installed, matching episodes will show a streaming source option alongside (or instead of) local/torrent options.

## Notes

- Streaming performance depends heavily on your chosen provider/source and your own connection.
- Tori does not host or provide any of this content — extensions and providers are third-party and unaffiliated with the project.

## Next steps

- [Streaming](Streaming.md) — full provider configuration reference
- [Auto Downloader](Auto-Downloader.md)
- [Remote Access](Remote-Access.md) — if you want to reach this setup from outside your home network
