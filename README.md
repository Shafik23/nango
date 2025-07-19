# Nano Golang App

A minimal template for creating single-binary web applications with Go.

## Features

- Single self-contained binary with embedded static files
- RESTful API endpoints
- Static file serving for HTML/CSS/JS
- Zero external dependencies
- Ready for Caddy/nginx reverse proxy deployment

## Quick Start

### Development
```bash
go run main.go
```
Visit http://localhost:8080

### Build for Linux
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app main.go
```

### Build for current platform
```bash
go build -ldflags="-s -w" -o app main.go
```

## Deployment with Caddy

Example Caddyfile:
```
your.domain.com {
    reverse_proxy 127.0.0.1:8080
}
```

## Project Structure

```
├── main.go          # Main application with embedded files
├── go.mod           # Go module file
├── static/          # All static assets (embedded at compile time)
│   ├── index.html
│   ├── styles.css
│   └── app.js
└── README.md
```

## API Endpoints

- `GET /` - Serves static files
- `GET /api/hello` - Returns JSON hello message
- `GET /health` - Health check endpoint

## How it Works

The `//go:embed static/*` directive tells Go to include all files from the `static/` directory directly into the compiled binary. This means you can deploy just the single executable file - no need to copy static assets separately.