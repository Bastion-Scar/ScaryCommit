# ScaryCommit v0.2.3

**ScaryCommit** is a CLI tool written in Go that uses **AI** to generate meaningful and stylish commit messages based on the content of your Git diff.  It analyzes the changes, consults an LLM (such as OpenRouter), and creates a "ready to go" commit 

# Features

- 🔍 Analyzes staged changes (`git diff --cached`)
- 🤖 Generates commit messages using AI (via OpenRouter API)
- 💬 Supports different commit styles (like Conventional)
- 🗣️ Supports languages (`en`, later — `ru`, `jp`, etc)
- 🪄 Minimalistic CLI interface in Cobra
- 💾 Automatic configuration initialization (`sco init`)
- 🧠 Secure confirmation before creating a commit

# Installation

# For Linux
- git clone https://github.com/Bastion-Scar/ScaryCommit.git
- cd ScaryCommit
- go build -o sco main.go
- sudo mv sco /usr/local/bin/

# For Windows

- git clone https://github.com/Bastion-Scar/ScaryCommit.git
- cd ScaryCommit
- go build -o sco.exe main.go

Optionally add to PATH(not necessary)

# Using
- sco init - Creates a yml configuration in which you must specify the API key and AI model
- sco commit - commits automatically (need to add the necessary files to the index git add)

# LINUX/MAC. If you are unable to download Go SDK for go build:
Option 1
- sudo chmod +x setup.sh
- ./setup.sh

Option 2
- make build
- sudo cp sco-linux /usr/local/bin/sco
- sudo cp sco-macos /usr/local/bin/sco   for Mac
- sudo chmod +x /usr/local/bin/sco


This will install the binary without installing Go

# WINDOWS. If you are unable to download Go SDK for go build:
Option 1
- come to PowerShell and type: Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned
- come to scarycommit folder like: cd C:\path\to\ScaryCommit
- .\setup.ps1
- [Environment]::SetEnvironmentVariable("PATH", "$env:USERPROFILE\bin;$env:PATH", "User") 
- restart PowerShell
- You can use sco

Option 2
- make build
- $InstallDir = "$env:USERPROFILE\bin"
- mkdir $InstallDir -Force
- copy sco-windows.exe $InstallDir\sco.exe
- [Environment]::SetEnvironmentVariable("PATH", "$InstallDir;$env:PATH", "User")
- restart powershell
- Use sco
