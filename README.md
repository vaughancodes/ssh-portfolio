# ssh-portfolio

A terminal-based portfolio served over SSH, built with Go and the [Charm](https://charm.sh) ecosystem.

## Features

- Tabbed navigation across **About**, **Experience**, **Projects**, **Skills**, **Education**, and **Contact** sections
- ASCII art banner with gradient coloring
- Clickable hyperlinks (in supported terminals)
- Responsive layout that adapts to terminal size
- Scrollable content via a viewport

## Tech Stack

- [Wish](https://github.com/charmbracelet/wish) — SSH server framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — terminal UI framework (TUI)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — styling and layout
- [Bubbles](https://github.com/charmbracelet/bubbles) — pre-built TUI components (viewport)

## Running Locally

```bash
go build -o ssh-portfolio .
./ssh-portfolio
```

The server starts on `0.0.0.0:22` by default. Connect with:

```bash
ssh -p 21 localhost
```

To use a different port:

```bash
./ssh-portfolio -p 2222
```

A host key is generated automatically in `.ssh/id_ed25519` on first run.

## Running via Docker Compose

Populate `.env` with a listening port:
```bash
LISTEN_PORT=2222
```

Then, run via Docker Compose:
```bash
docker compose up --build
```

## Controls

| Key | Action |
|---|---|
| `Tab` / `→` / `l` | Next tab |
| `Shift+Tab` / `←` / `h` | Previous tab |
| `↑` / `↓` / `j` / `k` | Scroll content |
| `q` / `Ctrl+C` | Quit |
