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
- [Installation](#installation)
- [Usage](#usage)
  - [Function for ~/.zshrc](#function-for-zshrc)
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

# TF summarize

Execute `tfsum terraform`, then you will see the original output of a plan/apply and a summary only printing the resource addr (specially when you are working with a lot of changes).

> tfsum is a custom function that you will see below, you can add it into your **.zshrc** or **.bashrc**

![tfsum](assets/example.png)

This summarized output can be useful, for example, for:

* You are migrating a terraform module and there are many changes that may be important in terms of destroying/creating resources (e.g., if you are migrating an EKS module from v17.X to v19.X).
* You use GitOps and deploy terraform from pipeline. The pipeline that makes the `terraform plan` can always show a summary of what is going to change (instead of having a super output of the original terraform plan).

# Installation

Take a look inside docs [install](./docs/install.md)

# Usage

Take a look inside docs [usage](./docs/usage.md)

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
