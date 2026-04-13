#!/bin/bash

set -e

REPO="adem02/epse"
BINARY="epse"
INSTALL_DIR="/usr/local/bin"

# Detect OS
OS=$(uname -s)
case "$OS" in
  Linux)  OS="Linux" ;;
  Darwin) OS="Darwin" ;;
  *)
    echo "❌ Unsupported OS: $OS"
    exit 1
    ;;
esac

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
  x86_64)  ARCH="x86_64" ;;
  arm64)   ARCH="arm64" ;;
  aarch64) ARCH="arm64" ;;
  *)
    echo "❌ Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

# Get latest version
echo "🔍 Fetching latest version..."
VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$VERSION" ]; then
  echo "❌ Could not fetch latest version"
  exit 1
fi

echo "📦 Installing EPSE $VERSION ($OS/$ARCH)..."

# Build download URL
FILENAME="${BINARY}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILENAME"

# Download and install
TMP_DIR=$(mktemp -d)
curl -sL "$URL" -o "$TMP_DIR/$FILENAME"
tar -xzf "$TMP_DIR/$FILENAME" -C "$TMP_DIR"
mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/$BINARY"
chmod +x "$INSTALL_DIR/$BINARY"
rm -rf "$TMP_DIR"

echo "✅ EPSE $VERSION installed successfully!"
echo "   Run 'epse --help' to get started."