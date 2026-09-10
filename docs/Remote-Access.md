# Remote Access

Ways to reach your Tori server from outside your home network.

## Option 1: Port forwarding (home server)

Forward your router's external port to your Tori machine's local IP and port (default `43211`). Your ISP-assigned public IP becomes your access point — a dynamic DNS service (e.g. DuckDNS, No-IP) is recommended since home IPs often change.

**Downsides**: exposes your server directly to the internet; requires your computer to stay on and connected; some ISPs block inbound ports or use CGNAT (no public IP at all).

## Option 2: VPN mesh (Tailscale / ZeroTier)

Install Tailscale (or similar) on both your Tori host and any device connecting to it. Devices get a private mesh IP that works from anywhere, with no ports exposed to the public internet.

**Best for**: personal use across your own devices, without wanting a public URL.

## Option 3: Cloud VPS

Run Tori entirely on a rented VPS (e.g. a small Ubuntu instance) instead of your home machine. This gives you a stable public IP, no dependence on your home network/ISP, and true 24/7 uptime.

1. Provision a Linux VPS (2 vCPU / 4GB RAM / 40GB disk is a reasonable baseline).
2. Build or upload the `tori` Linux binary and run it as a `systemd` service so it survives reboots.
3. Set `host = '0.0.0.0'` in `config.toml` (see [Configuration](Configuration.md)).
4. Point Tori Denshi/Tenji at the VPS's public IP and port.
5. **Set a server password** (Settings → General) since this instance is reachable by anyone who finds the IP.
6. Optionally put it behind a domain + HTTPS (e.g. via Caddy or nginx + Let's Encrypt) instead of raw `http://ip:port`.

**Note**: since your media library and/or torrent activity happens on the VPS itself, this only makes sense if you're not relying on files stored on your own computer — see [Desktop Streaming Setup](Desktop-Streaming-Setup.md) for a debrid/torrent-based setup that doesn't need local storage.

## Security checklist for any remote setup

- Set a server password
- Keep Tori updated to the latest release
- Don't expose your torrent client's Web UI (qBittorrent, etc) publicly — only the Tori server port needs to be reachable
