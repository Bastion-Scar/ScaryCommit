#!/bin/bash
set -e

# Building docker image
docker build -t sco-builder .

# Creating temporary container
docker create --name temp sco-builder
sudo docker cp temp:/usr/local/bin/sco /usr/local/bin/sco #Copy binary from container to our machine
docker rm temp

# Making file executable
sudo chmod +x /usr/local/bin/sco

echo "scarycommit installed в /usr/local/bin"
 # IIIITSSS TTVV TIIIIMEEE