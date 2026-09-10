# Nakama (Watch Together & Sharing)

Nakama lets you link multiple Tori instances together so you can watch together in sync with friends, or share your library/host instance with a peer without giving them full server access.

## Roles

- **Host** — the instance others connect to. Enable under **Settings → Nakama → Host** and set a username.
- **Peer** — connects to a host's Nakama session using the host's remote URL.

## Setting up as host

1. Go to **Settings → Nakama**.
2. Toggle **Enable hosting** on.
3. Set a display username (shown to peers).
4. Share your reachable server address (see [Remote Access](Remote-Access.md) if peers aren't on your local network) with whoever wants to connect.

## Connecting as a peer

1. Go to **Settings → Nakama**.
2. Enter the host's remote URL.
3. Connect — you'll be able to watch in sync with the host and other connected peers.

## Notes

- Nakama sync is best-effort; large latency differences between peers can cause drift.
- Hosting requires your instance to actually be reachable by peers — a purely local (`127.0.0.1`-bound) server won't work for anyone outside your machine.
