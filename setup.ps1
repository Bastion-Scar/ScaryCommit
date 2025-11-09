# setup.ps1 - Install ScaryCommit via Docker on Windows

param(
    [string]$InstallDir = "$env:USERPROFILE\bin"
)

# Check if docker installed and running
try {
    docker version | Out-Null
} catch {
    Write-Host "❌ Docker not found or not running. Please install/start Docker Desktop"
    exit 1
}

# Create the install directory if it doesn't exist
if (-Not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir | Out-Null
}

# Build docker image
Write-Host "Building docker image..."
docker build -t scarycommit-builder .

# Create a temporary container
$container = docker create scarycommit-builder

# Copy the Windows binary
Write-Host "Copying Windows binary to $InstallDir ..."
docker cp "${container}:/usr/local/bin/sco-windows.exe" "$InstallDir\sco.exe"

# Remove the temporary container
docker rm $container | Out-Null

Write-Host "✅ Binary installed to $InstallDir"

# Add to PATH if not already there
$currentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($currentPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$currentPath;$InstallDir", "User")
    Write-Host "Added $InstallDir to user PATH"
}

Write-Host "Done! You can now run: sco.exe"