# Local Anime Library

## Library paths

Add one or more folders under **Settings → Library**. Tori scans them recursively for video files and matches them against AniList entries.

## Scan modes

- **Standard scan** — matches existing files against your AniList collection; unmatched files are grouped separately.
- **Deep scan** — re-scans all files from scratch, useful after bulk-renaming or restructuring your library.

## Recommended structure

While Tori doesn't enforce strict naming, matching accuracy improves with:

```
/Anime
  /Series Name
    /Season 1
      Series Name - S01E01.mkv
      Series Name - S01E02.mkv
```

Single-season shows can skip the season subfolder.

## Locked/ignored files

Individual files or folders can be locked (excluded from re-scans) or ignored entirely from the library view if you don't want Tori to touch them.

## Related

- [Scanner](Scanner.md) — how file-to-entry matching actually works
- [Auto Downloader](Auto-Downloader.md)
