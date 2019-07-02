#!/bin/bash

GITHUB_LATEST_VERSION=$(curl -L -s -H 'Accept: application/json' https://github.com/jmartin82/compatip/releases/latest | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
GITHUB_FILE="compatip_${GITHUB_LATEST_VERSION//v/}_$(uname -s)_$(uname -m).tar.gz"
GITHUB_URL="https://github.com/jmartin82/compatip/releases/download/${GITHUB_LATEST_VERSION}/${GITHUB_FILE}"

wget -O compatip.tar.gz $GITHUB_URL
tar xzvf compatip.tar.gz compatip
sudo mv -f compatip /usr/local/bin/
rm compatip.tar.gz