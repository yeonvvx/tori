# Connect Tori Tenji to Your Desktop Server

Tori Tenji (the iOS/Android companion app) can connect to a Tori server running on your computer, so you can browse your library and stream from your phone.

## 1. Make sure your server is reachable on your network

By default Tori Denshi's embedded server binds to `127.0.0.1` (localhost only), which is **not** reachable from another device. To fix this:

1. Open your config file: `~/Library/Application Support/Tori/config.toml` (macOS — path differs on Windows/Linux).
2. Under `[server]`, set:
   ```toml
   host = '0.0.0.0'
   ```
3. Fully quit and reopen Tori Denshi so the change takes effect.

## 2. Find your computer's local IP address

- **macOS**: `ipconfig getifaddr en0` in Terminal
- **Windows**: `ipconfig` in Command Prompt, look for "IPv4 Address"
- **Linux**: `ip addr` or `hostname -I`

## 3. Connect from Tori Tenji

Open the app on your phone, and on the server-setup screen enter:

```
http://<your-computer's-local-IP>:43211
```

Your phone and computer must be on the **same Wi-Fi network** for this to work.

## Notes

- Only one Tori server process can bind to the port at a time — don't run a standalone `tori` binary alongside Tori Denshi, they'll conflict.
- Your computer's local IP can change (e.g. after reconnecting to Wi-Fi). If Tenji suddenly can't connect, re-check the IP and update it in the app.
- For a permanently reachable address (not tied to your home network), see [Remote Access](Remote-Access.md).

## Optional: password-protect your server

If your server is reachable beyond just your own devices, set a password under **Settings → General → Server Password** so a token is required to connect.
