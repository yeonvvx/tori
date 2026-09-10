# Scanner

Tori's scanner parses filenames and folder structure to match local files against AniList entries, without requiring a strict naming convention.

## What it reads

- Series title (from the folder or filename)
- Season/part indicators
- Episode number
- Release group / resolution tags (used for display, not matching)

## Matching process

1. Tori extracts a probable title from the file/folder name.
2. It searches your AniList collection (and, if enabled, AniList's broader database) for the closest match.
3. If confident, it auto-matches. If ambiguous, the file is left for manual matching in the library UI.

## Manual matching

Unmatched or misidentified entries can be corrected from the library view — select the file(s) and choose **Match to AniList entry** to search and assign manually.

## Tips for better matching

- Keep the series name consistent across all files in that series.
- Include the episode number in a recognizable format (`E01`, `Episode 1`, `- 01`, etc).
- Avoid mixing multiple series' files in one folder without season/series subfolders.

## Related

- [Local Anime Library](Local-Anime-Library.md)
