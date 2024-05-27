[![ReadMeSupportPalestine](https://raw.githubusercontent.com/Safouene1/support-palestine-banner/master/banner-project.svg)](https://github.com/Safouene1/support-palestine-banner)

# Exchango

Exchango is a command-line interface (CLI) based currency conversion tool
written in Go. It allows users to quickly convert between different currencies
using real-time exchange rates.

## Key Features

- Real-time[^1] currency conversion
- Support for multiple currencies
- Simple and intuitive command-line interface
- Lightweight and fast
- Works seamlessly on Linux, Windows, and macOS systems.
- Free and Open Source

[^1]: Currency rate is cached for 1 day so you don't have to hit the API on every
single request

## Getting Started

### Installation

- **Build from source**: `go install github.com/tfkhdyt/exchango`
- **Arch Linux (AUR)**: `yay -S exchango-bin`
- **Standalone binary:** Download the binary file from
  [release page](https://github.com/tfkhdyt/geminicommit/releases) and move the
  binary to one of the `PATH` directories in your system, for example:
  - **Linux:** `$HOME/.local/bin/`, `/usr/local/bin/`
  - **Windows:** `%LocalAppData%\Programs\`
  - **macOS:** `/usr/local/bin/`

### Usage

```bash
# list currencies
exchango list

# convert IDR to USD
exchango --from IDR --to USD 69420
```

More details in `exchango --help`

## License

This project is licensed under the GPLv3 License. See the LICENSE file for more details.
