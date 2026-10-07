# MagicHours

A terminal-based synchronicity tracker that displays meaningful messages at specific mirror times throughout the day. Built with Go and Bubble Tea for a beautiful TUI (Terminal User Interface) experience.

## About

MagicHours monitors the current time and displays spiritual messages associated with synchronistic times (like 11:11, 12:12, etc.). Each mirror time carries a unique message focused on self-awareness, spiritual growth, and personal development.

## Features

- Real-time synchronicity detection
- Beautiful terminal UI with dynamic color themes
- Countdown to the next synchronistic moment
- Interactive color customization (keyboard or mouse)
- Full-screen terminal experience

## Installation

### Prerequisites

- Go 1.24 or later

### Build from Source

```bash
git clone https://github.com/luismascotto/magic-hours.git
cd magic-hours
go build -o magic-hours
```

## Usage

Run the application:

```bash
./magic-hours
```

### Controls

- **q** or **Esc** or **Ctrl+C**: Quit
- **c** or **Mouse Click/Scroll**: Change color theme
- **Middle Mouse Button**: Quit

## Web Version

A static HTML version is available in the `web/` directory. See [web/README.md](web/README.md) for deployment instructions and how to add it to your iPhone home screen as a progressive web app.

## License

GPL-3.0 License - see [LICENSE](LICENSE) file for details.
