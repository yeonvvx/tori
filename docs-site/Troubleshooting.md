# Troubleshooting & Bug Reporting

## Finding your logs

Tori writes structured logs on startup and during operation. Locations:

- **macOS**: `~/Library/Application Support/Tori/logs`
- **Windows**: `%APPDATA%\Tori\logs`
- **Linux**: `~/.local/share/Tori/logs` (or wherever your `SEANIME_DATA_DIR`/data directory points)

Log lines are leveled (`TRC`, `DBG`, `INF`, `WRN`, `ERR`, `FTL`) — when reporting a bug, the `ERR`/`FTL` lines right before the failure are usually the most useful.

## Common issues

**"Address already in use" on startup**
Another Tori process (or Tori Denshi's embedded server) is already bound to the port. Only run one server instance at a time. Check with:
```
lsof -i :43211   # macOS/Linux
```

**Mobile app can't connect / times out**
- Confirm `host` in `config.toml` is `0.0.0.0`, not `127.0.0.1` (see [Connect Tenji to Desktop](Connect-Tenji-to-Desktop.md)).
- Confirm both devices are on the same network.
- Re-check your computer's local IP — it can change between sessions.

**qBittorrent login fails**
Double check the Web UI host/port/credentials match what's set in qBittorrent's own Web UI settings, and that "Bypass authentication for clients on localhost" isn't conflicting with a remote setup.

**Transcoding is slow or fails**
See [Transcoding & Direct Play](Transcoding-and-Direct-Play.md) — hardware acceleration availability varies a lot by platform and GPU.

## Reporting a bug

When opening an issue on the [GitHub repo](https://github.com/yeonvvx/tori/issues), include:

- Your OS and Tori version
- Steps to reproduce
- The relevant log excerpt (redact anything sensitive, like tokens)
