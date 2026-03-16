# FlexCLI

FlexCLI is a Go-based command-line interface for the FlexCoach AI fitness platform. It allows users and developers to interact with the backend API to manage profiles, training plans, goals, and system status.

## Features

- **Profile Management**: View and update user profile, body vitals, and preferences.
- **Training Plans**: Retrieve, update, or skip daily training plans.
- **Goal & Constraint Tracking**: Manage structured training goals and user constraints.
- **Device Connections**: Monitor and sync Garmin, Withings, and Rouvy data.
- **System Status**: Check backend health and service connectivity.

## Installation

### Homebrew (Recommended)

```bash
brew install f1dot4/flexcli/flexcli
```

### From Source

Ensure you have Go 1.22+ installed.

```bash
cd flexcli
make build
# Binaries will be in bin/flexcli-mac and bin/flexcli-linux
```

## Quick Start

Connect directly to a server without saving a configuration:

```bash
flexcli --server https://flexcoach.example.com --key YOUR_API_KEY profile body vitals get
```

Or configure a persistent context:

```bash
flexcli config --server https://flexcoach.example.com --key YOUR_API_KEY --name production
flexcli context use production
flexcli status
```

## Global Flags

- `--server`: Override the FlexCoach server URL.
- `--key`: Override the API Key (can also use `FLEXCLI_API_KEY` environment variable).
- `--context`: Use a specific context from the configuration file.
- `--config`: Specify a custom configuration file (default: `~/.flexcli.json`).

## Development

### Building
```bash
make build
```

### Releasing
The release process is fully automated via the `Makefile`. A single command handles testing, cross-compilation, documentation generation, version bumping, and git tagging:

```bash
make release v=0.1.7
```

**What this does:**
1. Runs all Go tests (`make test`).
2. Updates version strings in `main.go`.
3. Updates the Homebrew formula (`Formula/flexcli.rb`) with the new version and tag URL.
4. Cross-compiles binaries for macOS and Linux (`make build`).
5. Re-generates the full CLI reference documentation (`make docs`).
6. Calculates the SHA256 of the new release and updates the formula.
7. Commits all changes and creates a local Git tag (`v0.1.7`).

**Finalizing the release:**
After running `make release`, push the changes and the tag to GitHub:
```bash
git push origin main
git push origin v0.1.7
```
Once the tag is pushed, Homebrew users can upgrade via `brew upgrade flexcli`.
