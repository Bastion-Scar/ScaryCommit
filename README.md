# ScaryCommit v0.2.0

**ScaryCommit** is a CLI tool written in Go that uses **AI** to generate meaningful and stylish commit messages based on the content of your Git diff.  It analyzes the changes, consults an LLM (such as OpenRouter), and creates a "ready to go" commit 

# Features

- 🔍 Analyzes staged changes (`git diff --cached`)
- 🤖 Generates commit messages using AI (via OpenRouter API)
- 💬 Supports different commit styles (like Conventional)
- 🗣️ Supports languages (`en`, later — `ru`, `jp`, etc)
- 🪄 Minimalistic CLI interface in Cobra
- 💾 Automatic configuration initialization (`scarycommit init`)
- 🧠 Secure confirmation before creating a commit

# Installation

# For Linux
- git clone https://github.com/Bastion-Scar/ScaryCommit.git
- cd ScaryCommit
- go build -o scarycommit main.go
- sudo mv scarycommit /usr/local/bin/

# For Windows

- git clone https://github.com/Bastion-Scar/ScaryCommit.git
- cd ScaryCommit
- go build -o scarycommit.exe main.go

Optionally add to PATH(not necessary)

# Using
- scarycommit init - Creates a yml configuration in which you must specify the API key and AI model
- scarycommit commit - commits automatically (need to add the necessary files to the index git add)


