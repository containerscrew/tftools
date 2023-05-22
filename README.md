<p align="center" >
    <img src="assets/logo.png" alt="logo" width="250"/>
<h3 align="center">tftools</h3>
<p align="center">Easy CLI with useful terraform/terragrunt tools</p>
<p align="center">Build with ❤ in Golang</p>
</p>

<p align="center" >
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/containerscrew/tftools">
    <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/containerscrew/tftools">
    <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/containerscrew/tftools">
</p>

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Badges](#badges)
- [Installation](#installation)
  - [Using go](#using-go)
  - [Using brew](#using-brew)
  - [Using release binary (pending to finish this)](#using-release-binary-pending-to-finish-this)
    - [Linux](#linux)
    - [Mac OSX](#mac-osx)
    - [Supported OS](#supported-os)
  - [Check version](#check-version)
- [Usage](#usage)
  - [Function for ~/.zshrc](#function-for-zshrc)
- [Available tools in this CLI](#available-tools-in-this-cli)
- [Credits](#credits)
- [Contribution](#contribution)
- [LICENSE](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Badges

![Release Status](https://github.com/containerscrew/tftools/actions/workflows/release.yml/badge.svg)
![Build Status](https://github.com/containerscrew/tftools/actions/workflows/build.yml/badge.svg)
[![License](https://img.shields.io/github/license/containerscrew/tftools)](/LICENSE)
[![Release](https://img.shields.io/github/release/containerscrew/tftools)](https://github.com/containerscrew/tftools/releases/latest)
[![GitHub Releases Stats](https://img.shields.io/github/downloads/containerscrew/tftools/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=containerscrew&repository=tftools)

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

## Using release binary (pending to finish this)

### Linux
```bash
TFTOOLS_CLI_VERSION=$(curl -s https://api.github.com/repos/containerscrew/tftools/releases/latest)
TFTOOLS_CLI_ARCH=amd64
if [ "$(uname -m)" = "aarch64" ]; then CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/tftools/releases/download/${TFTOOLS_CLI_VERSION}/XXXXX-${TFTOOLS_CLI_ARCH}.tar.gz{,.sha256sum}
sha256sum --check xxxxxx-${CLI_ARCH}.tar.gz.sha256sum
sudo tar xzvfC xxxxxx-${CLI_ARCH}.tar.gz /usr/local/bin
rm xxxx-${CLI_ARCH}.tar.gz{,.sha256sum}
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

[Releases](https://github.com/containerscrew/tftools/releases)

### Supported OS

| OS        | ARM64 | AMD64 |
|-----------|:-----:|------:|
| Mac       |  √    |   √   |
| Linux     |  √    |   √   |

## Check version

```bash
tftools version
```

# Usage

Take a look inside [docs](./docs) folder.

```bash
tftools usage
```

## Function for ~/.zshrc

Copy [this function](scripts/tfsum.sh) in your `~/.zshrc` or `~/.bashrc` file.

```bash
function tfsum() {
    if [ -z "$1" ];
    then
        echo "You should type 'tfsum terraform|terragrunt'"
    else
        echo -e "Starting tf summary...\n"
        # Don't print output of terraform plan
        # If you don't want full plan output: $1 plan -out plan.tfplan 1> /dev/null
        $1 plan -out plan.tfplan
        echo -e "\n\n"
        $1 show -json plan.tfplan | tftools summarize
        # Delete plan out file to avoid git tracking (although is included in .gitignore)
        if [ -f "plan.tfplan" ]; then rm plan.tfplan; fi
    fi
}
```

```bash
source ~/.zshrc
tfsum terragrunt or tfsum terraform
```

# Available tools in this CLI

- [x] summarize `tftools summarize`: get a clear output/summary of changes when you perform a terraform/terragrunt plan
- [ ] statemv `tftools statemv`: interactive terraform/terragrunt state mv
- [ ] target-generator `tftools target-generator`: generates the final string of `terraform apply -target='' -target=''`command to execute applies with specific target instead of do it manually.

# Credits
- [Cobra to build beautiful CLI](https://cobra.dev/)
- [Terraform json structs for data parsing](https://github.com/hashicorp/terraform-json)
- [Distroless for container build](https://github.com/GoogleContainerTools/distroless)
- [Glamour markdown render](https://github.com/charmbracelet/glamour)
- [Official issue to solve this concern](https://github.com/hashicorp/terraform/issues/10507)

# Contribution

Pull requests are welcome! Any code refactoring, improvement, implementation.

# LICENSE

[LICENSE](./LICENSE)
