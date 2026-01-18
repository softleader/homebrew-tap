# brew-tapper

brew-tapper is a bot to automatic upgrade Homebrew Formula.

## Features
- Automatic upgrade Homebrew Formula version and sha256.
- Support multi-platform and multi-architecture (including Apple Silicon arm64).
- Modernized with Go 1.25.6.

## Requirements
- Go 1.25.6 or later.

## Installation
```bash
go get -u github.com/softleader/homebrew-tap/tapper/cmd/tapper
```

## Usage
```bash
tapper --owner softleader --token <GITHUB_TOKEN> --formula slctl --version 3.8.2 --dist _dist
```
