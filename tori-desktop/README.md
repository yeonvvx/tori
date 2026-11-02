<p align="center">
<img src="../docs/images/tori-logo.png" alt="preview" width="70px"/>
</p>

<h2 align="center"><b>Tori</b></h2>

<p align="center">
Electron-based desktop client for Tori. Embeds server and web interface. Successor to Tori Desktop.
</p>

<p align="center">
<img src="https://tori.app/bucket/img-2025-10-29-17-13-15.webp?updatedAt=1761758012992" alt="preview" width="80%"/>
</p>

---

## Prerequisites

- Go 1.24+
- Node.js 20+ and npm

---

## Development

### Web Interface

```shell
# Working dir: ./tori-web
npm run dev:denshi
```
 
### Sidecar

1. Build the server

	```shell
	# Working dir: .
 
	# Windows
	go build -o tori.exe -trimpath -ldflags="-s -w" -tags=nosystray
 
	# Linux, macOS
	go build -o tori -trimpath -ldflags="-s -w"
	```
 
2. Move the binary to `./tori/binaries`

3. Rename the binary:

   - For Windows: `tori-server-windows.exe`
   - For macOS/Intel: `tori-server-darwin-amd64`
   - For macOS/ARM: `tori-server-darwin-arm64`
   - For Linux/x86_64: `tori-server-linux-amd64`
   - For Linux/ARM64: `tori-server-linux-arm64`

### Electron

1. Setup

	```shell
	# Working dir: ./tori
	npm install
	```

2. Run

    `TEST_DATADIR` can be used in development mode, it should point to a dummy data directory for testing purposes.

    ```shell
    # Working dir: ./tori-desktop
    TEST_DATADIR="/path/to/data/dir" npm run dev
   ```

---

## Build

### Web Interface
   
```shell
# Working dir: ./tori-web
npm run build
npm run build:denshi
```

Move the output `./tori-web/out` to `./web`
Move the output `./tori-web/out-denshi` to `./tori/web-denshi`

```shell
# UNIX command
mv ./tori-web/out ./web
mv ./tori-web/out-denshi ./tori/web-denshi
```

### Sidecar

1. Build the server

	```shell
	# Working dir: .
 
	# Windows
	go build -o tori.exe -trimpath -ldflags="-s -w" -tags=nosystray
 
	# Linux, macOS
	go build -o tori -trimpath -ldflags="-s -w"
	```
 
2. Move the binary to `./tori/binaries`

3. Rename the binary:

   - For Windows: `tori-server-windows.exe`
   - For macOS/Intel: `tori-server-darwin-amd64`
   - For macOS/ARM: `tori-server-darwin-arm64`
   - For Linux/x86_64: `tori-server-linux-amd64`
   - For Linux/ARM64: `tori-server-linux-arm64`

### Electron

To build the desktop client for all platforms:

```
npm run build
```

To build for specific platforms:

```
npm run build:mac
npm run build:win
npm run build:linux
```

Output is in `./tori/dist/...`
