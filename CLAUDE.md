# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands

- **Build the project**: `./build.sh` - Creates a binary in `build/app` with embedded static files
- **Run for development**: `go run main.go` - Starts server on http://localhost:8080
- **Manual build**: `go build -ldflags="-s -w" -o build/app main.go`
- **Cross-compile for Linux**: `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o build/app main.go`

## Architecture

This is a single-binary Go web application template that embeds all static assets at compile time using Go's `embed` package. The key architectural principle is zero external dependencies - everything needed to run the application is compiled into one executable.

### How Static Embedding Works

The `//go:embed static/*` directive in `main.go` tells the Go compiler to include all files from the `static/` directory into the binary. These files are then served using `http.FileServer` with an embedded filesystem, making deployment as simple as copying a single binary.

### API Structure

- Static files served from root path (`/`)
- JSON API endpoints under `/api/*`
- Health check at `/health`

When adding new API endpoints, follow the pattern in `main.go` of using `http.HandleFunc` with the `/api/` prefix.