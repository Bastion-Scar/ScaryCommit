#!/bin/bash
set -e

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64) ARCH="arm64" ;;
    *) ARCH="amd64" ;;
esac

echo "Detected OS: $OS, Architecture: $ARCH"

# Building docker image
echo "Building docker image..."
docker build -t sco-builder .

# Creating temporary container
echo "Creating container..."
container_id=$(docker create sco-builder)

# Determine binary name and install path
if [ "$OS" = "linux" ]; then
    BINARY_NAME="sco-linux"
    INSTALL_PATH="/usr/local/bin/sco"
    COPY_CMD="sudo docker cp"
    CHMOD_CMD="sudo chmod"
elif [ "$OS" = "darwin" ]; then
    BINARY_NAME="sco-macos" 
    INSTALL_PATH="/usr/local/bin/sco"
    COPY_CMD="docker cp"
    CHMOD_CMD="chmod"
else
    echo "Unsupported OS: $OS"
    exit 1
fi

# Copy appropriate binary
echo "Copying $BINARY_NAME to $INSTALL_PATH..."
$COPY_CMD "$container_id:/usr/local/bin/$BINARY_NAME" "$INSTALL_PATH"

# Remove container
docker rm "$container_id"

# Make executable
$CHMOD_CMD +x "$INSTALL_PATH"

echo "✅ scarycommit installed to $INSTALL_PATH"
echo "Done! You can now run: sco"
 # IIIITSSS TTVV TIIIIMEEE