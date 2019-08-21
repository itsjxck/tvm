#!/bin/bash

case "$OSTYPE" in
  darwin*) PLATFORM=darwin ;;
  linux*) PLATFORM=linux ;;
  *) echo "Can't detect OS"; exit ;;
esac

DOWNLOAD_URL=$(curl -s "https://api.github.com/repos/itsjxck/hashiman/releases/latest" | grep -o "http.*${PLATFORM}_amd64")
curl -L --silent -o /usr/local/bin/hashiman ${DOWNLOAD_URL}
chmod +x /usr/local/bin/hashiman