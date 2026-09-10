# Standalone iOS Setup

Tori Tenji for iOS supports the same idea as Android — running its own embedded server on-device — though iOS's sandboxing makes this more limited than Android.

## Requirements

- A sideloaded build of Tori Tenji (via Xcode with a free personal Apple ID, or a proper Apple Developer Program build for a real distributable `.ipa`)
- iOS 16+ recommended

## 1. Enable the Mobile Server

In the app, go to **Settings → Mobile Server** and enable it. As on Android, this runs a local instance of the Tori backend using the app's sandboxed storage.

## 2. Connect to it

Use `http://127.0.0.1:<port>` on the setup screen (default port 43211).

## Limitations on iOS

- Background execution is restricted by iOS — the embedded server generally only runs while the app is in the foreground or briefly backgrounded, unlike a desktop server which runs indefinitely.
- Free personal-team sideloads expire after 7 days and need to be reinstalled from Xcode; a paid Apple Developer account removes this limit.
- For an always-on server reachable from your phone, a desktop or [Remote Access](Remote-Access.md) VPS setup is more reliable than the on-device iOS server.

## Next steps

- [Connect Tenji to Desktop](Connect-Tenji-to-Desktop.md) — the more reliable alternative
- [Remote Access](Remote-Access.md)
