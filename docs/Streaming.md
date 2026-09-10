# Streaming

Reference for configuring Tori's streaming providers — torrent streaming, debrid services, and online-streaming extensions.

## Torrent streaming (built-in)

No account needed. Streams directly from BitTorrent peers to the player. Performance depends on seeders and your connection.

## Debrid services

Supported: **Torbox, Real-Debrid, AllDebrid, Premiumize**.

1. Get an API key/token from your provider's dashboard.
2. Enter it under **Settings → Streaming → [Provider]**.
3. Debrid-sourced torrents will show as an available streaming source alongside built-in torrent streaming.

Debrid generally streams faster and more reliably than raw BitTorrent since the provider handles the download server-side.

## Online streaming extensions

Installed via the [Extension Marketplace](Extensions.md). Once installed, matching shows display an online-streaming source option. Quality/availability depends entirely on the extension and its upstream source.

## Choosing a source

When multiple sources are available for an episode (local file, torrent stream, debrid, online extension), Tori lets you pick which to play from the episode's source selector.

## Related

- [Desktop Streaming Setup](Desktop-Streaming-Setup.md)
- [Transcoding & Direct Play](Transcoding-and-Direct-Play.md)
