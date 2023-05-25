<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Installation](#installation)
  - [Using go](#using-go)
  - [Using brew](#using-brew)
  - [Using release binary](#using-release-binary)
    - [Linux](#linux)
    - [Mac OSX](#mac-osx)
    - [Supported OS](#supported-os)
  - [Check version](#check-version)
- [Updating](#updating)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Installation

## Using go

```bash
go install github.com/containerscrew/tftools
```

## Using brew

```bash
brew tap containerscrew/tftools https://github.com/containerscrew/tftools
brew install tftools
```

## Using release binary

### Linux
```bash
TFTOOLS_LATEST_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest | jq -r ".name")
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "aarch64" ]; then TFTOOLS_CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_LATEST_VERSION}/tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz
tar -xzf tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz tftools
sudo mv tftools /usr/local/bin/tftools
rm tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz
```

### Mac OSX

```bash
TFTOOLS_LATEST_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest | jq -r ".name")
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "arm64" ]; then TFTOOLS_CLI_ARCH=arm64; fi
curl -L --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_LATEST_VERSION}/tftools-darwin-${TFTOOLS_CLI_ARCH}.tar.gz
tar -xzf tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz tftools
sudo mv tftools /usr/local/bin/tftools
rm tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz
```

You will find apk, rpm and deb packages in [releases](https://github.com/containerscrew/tftools/releases)

For example, a deb package:

```bash
TFTOOLS_LATEST_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest | jq -r ".name")
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "aarch64" ]; then TFTOOLS_CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_LATEST_VERSION}/tftools-linux-${TFTOOLS_CLI_ARCH}.deb
sudo dpkg -i tftools-linux-${TFTOOLS_CLI_ARCH}.deb
rm dpkg -i tftools-linux-${TFTOOLS_CLI_ARCH}.deb
```

See full [scripts](../scripts)

### Supported OS

| OS        | ARM64 | AMD64 |
|-----------|:-----:|------:|
| Mac       |  √    |   √   |
| Linux     |  √    |   √   |

## Check version

```bash
tftools version
```

# Updating

```bash
brew update
brew upgrade containerscrew/tftools/tftools
```
