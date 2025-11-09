# setup.ps1 - Install ScaryCommit via Docker on Windows

param(
    [string]$InstallDir = "$env:USERPROFILE\bin"
)

# Check if docker installed
if (-not (Get-Command docker -ErrorAction SilentlyContinue)) {
    Write-Host "❌ Docker not found. Please install Docker Desktop for Windows"
    exit 1
}

# Create the install directory if it doesnt exist
if (-Not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir | Out-Null
}

# Build docker image for windows binary
Write-Host "Building docker image..."
docker build -t scarycommit-builder .

# Create a temporary container
$container = docker create scarycommit-builder

# Copy the binary to install directory
Write-Host "Copying binary to $InstallDir ..."
docker cp "$container:/scarycommit/scarycommit.exe" "$InstallDir\scarycommit.exe"

# Remove the temporary container
docker rm $container | Out-Null

Write-Host "✅ Binary installed to $InstallDir"

# Check if install directory is in PATH
if ($env:PATH -notlike "*$InstallDir*") {
    Write-Host "Warning: $InstallDir is not in your PATH."
    Write-Host "You can add it via System Properties → Environment Variables → PATH."
}

Write-Host "Done! You can now run: scarycommit.exe init"
