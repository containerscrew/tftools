<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Installation](#installation)
  - [Using go](#using-go)
  - [Using brew](#using-brew)
  - [Using release binary (pending to finish this)](#using-release-binary-pending-to-finish-this)
    - [Linux](#linux)
    - [Mac OSX](#mac-osx)
    - [Supported OS](#supported-os)
  - [Check version](#check-version)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Installation

## Using go

Go install is not a good solution because in order to compile the binary you must first run go generate ./....

```bash
go install github.com/containerscrew/tftools
```

> **NOTE:** this will fail as I mentioned 🛠️

You can do the following:

```bash
git clone https://github.com/containerscrew/tftools.git && cd tftools
go generate ./...
go install .
```

## Using brew

```bash
brew tap containerscrew/tftools https://github.com/containerscrew/tftools
brew install tftools
```

## Using release binary (pending to finish this)

### Linux
```bash
TFTOOLS_LATEST_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest | jq -r ".name")
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "aarch64" ]; then TFTOOLS_CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_LATEST_VERSION}/tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz
sudo tar xzvfC tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz /usr/local/bin
rm tftools-linux-${TFTOOLS_CLI_ARCH}.tar.gz
```

### Mac OSX

```bash
TFTOOLS_CLI_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest)
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "arm64" ]; then CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_CLI_VERSION}/XXXXX-${TFTOOLS_CLI_ARCH}.tar.gz{,.sha256sum}
sha256sum --check xxxxxx-${CLI_ARCH}.tar.gz.sha256sum
sudo tar xzvfC xxxxxx-${CLI_ARCH}.tar.gz /usr/local/bin
rm xxxx-${CLI_ARCH}.tar.gz{,.sha256sum}
```

You will find apk, rpm and deb packages in [releases](https://github.com/containerscrew/tftools/releases)

### Supported OS

| OS        | ARM64 | AMD64 |
|-----------|:-----:|------:|
| Mac       |  √    |   √   |
| Linux     |  √    |   √   |

## Check version

```bash
tftools version
```
