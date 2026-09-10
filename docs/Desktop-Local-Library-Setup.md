# Desktop Local Library Setup

This guide covers setting up Tori Denshi to manage and play a local anime/manga library on your computer — no torrenting or streaming providers required.

## 1. Install Tori Denshi

Download the latest build for your OS (Windows, macOS, or Linux) from the [releases page](https://github.com/yeonvvx/tori/releases) and install it like any other desktop app.

## 2. Point Tori at your library

On first launch, open **Settings → Library** and add the folder(s) where your anime lives. Tori scans recursively, so subfolders are fine.

Tori doesn't require strict naming conventions, but consistent folder structure (one folder per series, one subfolder per season if applicable) gives the best matching results. See [Scanner](Scanner.md) for details on how matching works.

## 3. Link AniList (optional but recommended)

Go to **Settings → AniList** and authenticate. This lets Tori match your local files against AniList entries, pull metadata/cover art, and sync your watch progress.

## 4. Set up a media player

Tori Denshi ships with a built-in libmpv-based player (supports SSA/ASS subtitles, shaders, and hardware decoding). Alternatively, configure an external player:

- **MPV** — point Tori at your `mpv` executable path in Settings → Media Player
- **VLC** — same, via the VLC executable path
- **MPC-HC** — Windows only

## 5. Optional: qBittorrent for downloads

If you also want to download torrents into your local library (rather than stream them), install [qBittorrent](https://www.qbittorrent.org/) and connect it under **Settings → Torrent Client** with its Web UI host, port, username, and password.

## Next steps

- [Local Anime Library](Local-Anime-Library.md) — library paths and scan modes in depth
- [Auto Downloader](Auto-Downloader.md) — automatically grab new episodes
- [UI Customization](UI-Customization.md)
