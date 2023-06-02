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
- [TF summarize](#tf-summarize)
  - [Example](#example)
  - [Steps](#steps)
- [Installation](#installation)
  - [Quick installation (latest version)](#quick-installation-latest-version)
- [Usage](#usage)
  - [Built-in subcommand](#built-in-subcommand)
- [Credits](#credits)
- [TO DO](#to-do)
- [Contribution](#contribution)
- [LICENSE](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Badges

![Release Status](https://github.com/containerscrew/tftools/actions/workflows/release.yml/badge.svg)
![Build Status](https://github.com/containerscrew/tftools/actions/workflows/build.yml/badge.svg)
![Git Leaks Status](https://github.com/containerscrew/tftools/actions/workflows/gitleaks.yml/badge.svg)
[![License](https://img.shields.io/github/license/containerscrew/tftools)](/LICENSE)
[![Release](https://img.shields.io/github/release/containerscrew/tftools)](https://github.com/containerscrew/tftools/releases/latest)
[![GitHub Releases Stats](https://img.shields.io/github/downloads/containerscrew/tftools/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=containerscrew&repository=tftools)

# TF summarize

**The concern is:** I have a lot of changes in terraform and I need a clear way of the concepts that are going to be **deleted|changed|created** only with the resource address. It can be messy to have a super tf plan output when there are **many changes**.


## Example

Imagine you are going to:

- Create a new s3 test bucket
- Change your ALB ingress controller policy and vpc-cni addon
- Delete grafana backup tool helm chart

## Steps

1. Makes the appropriate changes from code.

2. Execute `tfsum terraform`, then you will see the original output of a plan/apply and a summary only printing the resource addr and the action.

> tfsum is a custom function. See [usage](#usage) or [tfsum.sh](scripts/tfsum.sh)

![tfsum](assets/example.png)

*This summarized output can be useful, for example, for:*

* You are migrating a terraform module and there are many changes that may be important in terms of destroying/creating resources (e.g., if you are migrating an EKS module from v17.X to v19.X).
* You use GitOps and deploy terraform from pipeline. The pipeline that makes the `terraform plan` can always show a summary of what is going to change (instead of having a super output of the original terraform plan).

# Installation

## Quick installation (latest version)

```bash
curl --proto '=https' --tlsv1.2 -sSfL https://raw.githubusercontent.com/containerscrew/tftools/main/scripts/install.sh | bash
```

Take a look inside docs [install](./docs/install.md)

# Usage

Take a look inside docs [usage](./docs/usage.md)

## Built-in subcommand

`tfsum usage` is subcommand that prints the contents of [usage.md](docs/usage.md) in pretty terminal markdown render

```bash
tftools usage
```

> Requires internet connectivity, as it fetches the [usage.md](https://raw.githubusercontent.com/containerscrew/tftools/main/docs/usage.md) file.

# Credits
- [Cobra to build beautiful CLI](https://cobra.dev/)
- [Terraform json structs for data parsing](https://github.com/hashicorp/terraform-json)
- [Distroless for container build](https://github.com/GoogleContainerTools/distroless)
- [Glamour markdown render](https://github.com/charmbracelet/glamour)
- [Official issue to solve this concern](https://github.com/hashicorp/terraform/issues/10507)
- [Git leaks](https://github.com/gitleaks/gitleaks-action)
- [To my teacher of Golang @gilmiriam](https://github.com/gilmiriam)

# TO DO

* Improve error handling
* Add tests, although I have no experience
* Code refactor is certainly needed!

# Contribution

Pull requests are welcome! Any code refactoring, improvement, implementation.

# LICENSE

[LICENSE](./LICENSE)
