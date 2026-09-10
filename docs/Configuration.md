# Configuration

Tori stores its settings in a `config.toml` file inside its data directory.

## File location

- **macOS**: `~/Library/Application Support/Tori/config.toml`
- **Windows**: `%APPDATA%\Tori\config.toml`
- **Linux**: `~/.local/share/Tori/config.toml` (or under `$SEANIME_DATA_DIR` if set)

Most settings can also be changed from the in-app Settings UI, which writes back to this file — editing it directly is mainly useful for headless/server setups or fixing something the UI won't let you (like the bind host).

## `[server]`

| Key | Description |
|---|---|
| `host` | Interface to bind to. `127.0.0.1` = local machine only. `0.0.0.0` = reachable from other devices on your network (required for mobile apps, VPS setups, etc). |
| `port` | Port the server listens on. Default `43211`. |
| `offline` | Start in offline mode. |
| `usebinarypath` | Whether to resolve external tool paths (mpv, ffmpeg, etc.) relative to the binary's location. |

## `[cache]`

Controls where Tori caches images, metadata, and transcoded segments. See [Transcoding & Direct Play](Transcoding-and-Direct-Play.md) for transcode-specific cache behavior.

## Environment variables

| Variable | Description |
|---|---|
| `SEANIME_DATA_DIR` | Overrides the data directory location. |
| `SEANIME_WORKING_DIR` | Overrides the working directory Tori runs from. |

> These env var names are internal identifiers carried over from the underlying engine and don't affect branding — they're not shown to you anywhere in the app.

## Applying changes

Config file edits require a full restart of the server (fully quit Tori Denshi, not just close the window) to take effect.

## Related

- [Local Anime Library](Local-Anime-Library.md)
- [Streaming](Streaming.md)
- [Remote Access](Remote-Access.md)
