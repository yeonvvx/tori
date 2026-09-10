# Standalone Android Setup

Tori Tenji for Android can run its own embedded server directly on your device, so you don't need a separate computer running Tori at all.

## 1. Install Tori Tenji

Install the app on your Android device (sideloaded APK or, if published, from the store listing).

## 2. Start the Mobile Server

In the app, go to **Settings → Mobile Server** and enable it. This starts a local copy of the Tori server backend on your phone/tablet, using on-device storage for its data directory and cache.

## 3. Point the app at itself

Once the Mobile Server is running, use `http://127.0.0.1:<port>` (the port shown in the Mobile Server screen — 43211 by default) as your server address on the setup screen.

## Notes

- Running the server on-device means your library data lives on that device — back it up like you would any other local data.
- Performance for transcoding/streaming depends on your device's hardware; low-end devices may struggle with on-the-fly transcoding of large files.
- The Mobile Server can also be reached by other devices on the same network the same way a desktop server would be — see [Connect Tenji to Desktop](Connect-Tenji-to-Desktop.md) for the same networking logic (find the device's local IP, ensure it's not restricted to loopback, etc).

## Next steps

- [Offline Mode](Offline-Mode.md)
- [Remote Access](Remote-Access.md)
