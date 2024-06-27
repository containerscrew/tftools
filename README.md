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
- [Tftools summarize](#tftools-summarize)
- [Installation](#installation)
  - [Install latest version](#install-latest-version)
  - [Install specific release](#install-specific-release)
  - [Container image](#container-image)
- [Usage](#usage)
  - [Built-in subcommand](#built-in-subcommand)
- [Example](#example)
- [TO DO](#to-do)
- [Contribution](#contribution)
- [LICENSE](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Badges

![Release Status](https://github.com/containerscrew/tftools/actions/workflows/release.yml/badge.svg)
![Build Status](https://github.com/containerscrew/tftools/actions/workflows/build.yml/badge.svg)
![Git Leaks Status](https://github.com/containerscrew/tftools/actions/workflows/gitleaks.yml/badge.svg)
![Lint Status](https://github.com/containerscrew/tftools/actions/workflows/lint.yml/badge.svg)
![Gosec Status](https://github.com/containerscrew/tftools/actions/workflows/gosec.yml/badge.svg)
![Test Status](https://github.com/containerscrew/tftools/actions/workflows/test.yml/badge.svg)
[![License](https://img.shields.io/github/license/containerscrew/tftools)](/LICENSE)
[![Release](https://img.shields.io/github/release/containerscrew/tftools)](https://github.com/containerscrew/tftools/releases/latest)
[![GitHub Releases Stats](https://img.shields.io/github/downloads/containerscrew/tftools/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=containerscrew&repository=tftools)

# Tftools summarize

**The concern is:** I have a lot of changes in terraform and I need a clear way of the concepts that are going to be **deleted|changed|created** only with the resource address. It can be messy to have a super tf plan output when there are **many changes**.

# Installation

## Install latest version

```shell
curl --proto '=https' --tlsv1.2 -sSfL https://raw.githubusercontent.com/containerscrew/tftools/main/scripts/install.sh | sh
```

## Install specific release

```shell
curl --proto '=https' --tlsv1.2 -sSfL https://raw.githubusercontent.com/containerscrew/tftools/main/scripts/install.sh | sh -s -- -v "v0.8.0"
```

> [!NOTE]
> If you don't specify `-v` flag, by default will install always latest version.

## Container image

[In this other repo](https://github.com/containerscrew/infratools) I have a container image where you can find this tool `tftools` installed.

https://hub.docker.com/r/containerscrew/infratools/

> Take a look inside [install](./docs/install.md) documentation for other installation methods.

# Usage

Take a look inside docs [usage](./docs/usage.md)

## Built-in subcommand

`tftools usage` is subcommand that prints the contents of [usage.md](docs/usage.md) in pretty terminal markdown render

```bash
tftools usage
```

> Requires internet connectivity, as it fetches the [usage.md](docs/usage.md) file.

# Example

![example](assets/example.png)

*This summarized output can be useful, for example, for:*

* You are migrating a terraform module and there are many changes that may be important in terms of destroying/creating resources (e.g., if you are migrating an EKS module from v17.X to v19.X).
* You use GitOps and deploy terraform from pipeline. The pipeline that makes the `terraform plan` can always show a summary of what is going to change (instead of having a super output of the original terraform plan).

# TO DO

* Improve error handling
* Add tests, although I have no experience
* Code refactor is certainly needed!
* Create new subcommand for an interative terraform state mv target migration (when you need to move a lot of resources)
* Other subcommand when you need to apply only certain targets (terraform apply -target=x -target=x ...)

# Contribution

Pull requests are welcome! Any code refactoring, improvement, implementation.

# LICENSE

[LICENSE](./LICENSE)
