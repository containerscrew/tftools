#!/usr/bin/env bash

TFTOOLS_LATEST_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest | jq -r ".name")
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "aarch64" ]; then TFTOOLS_CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_LATEST_VERSION}/tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz{,.sha256sum}
sha256sum --check tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz.sha256sum
sudo tar xzvfC tftools-${CLI_ARCH}.tar.gz /usr/local/bin
rm tftools-${CLI_ARCH}.tar.gz{,.sha256sum}
