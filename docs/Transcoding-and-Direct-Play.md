# Transcoding & Direct Play

## Direct Play

If your device/browser supports a file's codec and container natively, Tori serves it as-is — no CPU/GPU cost, best quality, fastest start time.

## Transcoding

When direct play isn't possible (unsupported codec, remote device with limited codec support, bandwidth constraints), Tori transcodes on the fly using FFmpeg.

### Hardware acceleration

Tori can use GPU-based hardware encoding/decoding where available:

- **macOS**: VideoToolbox
- **Windows**: NVENC (NVIDIA), QuickSync (Intel), AMF (AMD)
- **Linux**: VAAPI / NVENC depending on GPU and drivers

Enable under **Settings → Transcoding → Hardware Acceleration**. If unavailable or misconfigured, Tori falls back to software (CPU) transcoding, which is slower and more CPU-intensive — this matters if you're running Tori on a low-power VPS.

### Quality/bitrate settings

Adjust target resolution and bitrate under **Settings → Transcoding** to balance quality against your (or your viewers') bandwidth.

## Related

- [Streaming](Streaming.md)
- [Remote Access](Remote-Access.md) — bandwidth matters more once you're streaming outside your LAN
