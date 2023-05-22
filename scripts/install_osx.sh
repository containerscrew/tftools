#!/usr/bin/env bash

TFTOOLS_CLI_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest)
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "arm64" ]; then CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_CLI_VERSION}/XXXXX-${TFTOOLS_CLI_ARCH}.tar.gz{,.sha256sum}
sha256sum --check xxxxxx-${CLI_ARCH}.tar.gz.sha256sum
sudo tar xzvfC xxxxxx-${CLI_ARCH}.tar.gz /usr/local/bin
rm xxxx-${CLI_ARCH}.tar.gz{,.sha256sum}
