# Author: Michael Bopp
# Filename: Dockerfile
# Description: Dockerfile for CS330 Presentation.
# Date Created: 8/23/23
# Last Edited: 8/29/23

# Base the container on the latest Alpine Linux image.
FROM alpine:latest

# Install Busybox, Go, Neovim, Python, and Vim
RUN apk add --no-cache busybox go neovim python3 vim

# Optional: Set the working directory to your Go workspace
WORKDIR workspace

# source -> destination
COPY script.py /script.py

RUN python3 /script.py

EXPOSE 80