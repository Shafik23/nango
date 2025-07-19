#!/bin/bash

# Build script for nano-golang app

set -e

echo "🔨 Building nano-golang app..."

# Create build directory
mkdir -p build

# Build single binary
go build -ldflags="-s -w" -o build/app main.go

# Show file size
echo ""
echo "✅ Build complete!"
echo ""
echo "📦 Binary size:"
ls -lh build/app

echo ""
echo "🚀 To run: ./build/app"
echo "🌐 Server will start on http://localhost:8080"