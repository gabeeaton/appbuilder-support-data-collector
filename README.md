# AppBuilder Support Data Collector

A desktop application built with [Wails3](https://v3.wails.io/), using a Go backend and a Svelte/TypeScript frontend.

## Prerequisites

- [Go](https://go.dev/dl/) 1.25 or later
- [Node.js](https://nodejs.org/) (LTS) and npm
- Platform build tools:
  - **macOS**: Xcode command line tools (`xcode-select --install`)
  - **Windows**: WebView2 (usually preinstalled on Windows 10/11). No C compiler is needed — Wails3's Windows backend is pure Go/CGO-free; one would only be required if a future dependency needs cgo.
  - **Linux**: `gtk4` and `webkitgtk-6.0` development packages (e.g. `libgtk-4-dev libwebkitgtk-6.0-dev` on Ubuntu 24.04+/Debian 13+)

## Installing the Wails3 CLI

Install the Wails3 CLI with:

```bash
go install github.com/wailsapp/wails/v3/cmd/wails3@v3.0.0
```

Confirm it installed correctly:

```bash
wails3 version
```

Then check your environment is set up correctly for Wails development:

```bash
wails3 doctor
```

## Getting Started

Clone the repository and install frontend dependencies:

```bash
git clone git@github.com:sillsdev/appbuilder-support-data-collector.git
cd appbuilder-support-data-collector
cd frontend && npm install && cd ..
```

Run the app in development mode with hot reload for both the Go backend and the frontend:

```bash
wails3 dev
```

## Command-line options

```bash
sdc -project "/path/to/project.appDef" -theme dark -lang fr
```

All flags are optional: omit `-project` to choose a folder in the wizard; theme defaults to the OS preference and language to `en`. `-theme dark` selects DaisyUI’s `dim` theme and `-theme light` selects `corporate`. Settings are exposed through `SettingsService.GetSettings()`. Use `-help` for usage.

## Building

To produce a production build for your current platform:

```bash
wails3 build
```

Build output is written to `build/bin/`.

## Releases

Pushing a `v*` tag triggers [`.github/workflows/release.yml`](.github/workflows/release.yml), which builds a plain binary natively for Linux (amd64), macOS (arm64 + amd64), and Windows (amd64), then uses [goreleaser](https://goreleaser.com/) to attach them to a GitHub Release along with a `checksums.txt`.

These are unpackaged binaries only, meant to be consumed by [Scripture App Builder](https://software.sil.org/scriptureappbuilder/)'s own packaging (e.g. its Linux Flatpak, which supplies the GTK4/WebKitGTK 6.0 runtime dependencies) - not installers in their own right. There's no Linux `.deb`/`.rpm`/AppImage packaging and no iOS/Android build support, since this app isn't distributed through those channels. A signed macOS `.dmg` and a Windows installer (via InnoSetup) may be added later if this app needs to be distributed independently of SAB.

## Project Structure

A typical Wails3 project is laid out as:

```
.
├── main.go          # Application entry point
├── app.go           # App struct and Go methods exposed to the frontend
├── frontend/         # Svelte/TypeScript frontend source
│   ├── src/
│   └── package.json
├── build/            # Platform build assets and output (ignored in git)
└── wails.json        # Wails project configuration
```

## Learn More

- [Wails3 Documentation](https://v3.wails.io/)
- [Wails3 GitHub Repository](https://github.com/wailsapp/wails)

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file.
