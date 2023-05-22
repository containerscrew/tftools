#!/usr/bin/env bash

TFTOOLS_CLI_VERSION=latest
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "aarch64" ]; then CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_CLI_VERSION}/tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz{,.sha256sum}
sha256sum --check tftools-${CLI_ARCH}.tar.gz.sha256sum
sudo tar xzvfC tftools-${CLI_ARCH}.tar.gz /usr/local/bin
rm tftools-${CLI_ARCH}.tar.gz{,.sha256sum}
