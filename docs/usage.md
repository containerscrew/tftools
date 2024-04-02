<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Usage](#usage)
  - [Summarize](#summarize)
  - [Function for zsh](#function-for-zsh)
  - [Function for fish](#function-for-fish)
  - [Function for bash](#function-for-bash)
  - [Load new functions](#load-new-functions)
  - [Using tfsum as a custom binary](#using-tfsum-as-a-custom-binary)
- [Example](#example)
- [JSON output support](#json-output-support)
  - [JSON output with arns](#json-output-with-arns)
  - [JSON output only metrics](#json-output-only-metrics)
  - [JSON output pretty](#json-output-pretty)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Usage

## Summarize

```shell
terraform plan -out plan.tfplan
terraform show -json plan.tfplan | tftools summarize --show-tags
```

Or if you have the file in json:

```shell
terraform plan -out plan.tfplan
terraform show -json plan.tfplan > plan.json
cat plan.json | tftools summarize
```

## Function for zsh

Edit your `~/.zshrc`

```bash
function tfsum() {
  if [ -z "$1" ];
  then
    echo "You should type 'tfsum terraform|terragrunt'"
  else
    echo -en "Starting tf summary... Please wait"
    # If you want to print full plan output: $1 plan -out plan.tfplan
    $1 plan -out plan.tfplan 1> /dev/null
    $1 show -json plan.tfplan | tftools summarize --show-tags
    # Delete plan out file to avoid git tracking (although is included in .gitignore)
    if [ -f "plan.tfplan" ]; then rm plan.tfplan; fi
  fi
}
```

## Function for fish

Edit your `~/.config/fish/config.fish` or create a new file inside `~/.config/fish/functions/tfsum.fish`

```shell
function tfsum
    if test -z $argv[1]
        echo "You should type 'tfsum terraform|terragrunt'"
    else
        echo -en "Starting tf summary... Please wait"
        # If you want to print full plan output: $argv[1] plan -out plan.tfplan
        $argv[1] plan -out plan.tfplan 1> /dev/null
        $argv[1] show -json plan.tfplan | tftools summarize --show-tags
        # Delete plan out file to avoid git tracking (although is included in .gitignore)
        if test -f "plan.tfplan"; rm plan.tfplan; end
    end
end
```

## Function for bash

Edit your `~/.bashrc`

```shell
tfsum() {
  if [ -z "$1" ]; then
    echo "You should type 'tfsum terraform|terragrunt'"
    exit 1
  fi

  echo -en "Starting tf summary... Please wait\n"

  if [ -n "$2" ] && [ "$2" == "-v" ]; then
    "$1" plan -out plan.tfplan
  else
    "$1" plan -out plan.tfplan 1> /dev/null
  fi

  "$1" show -json plan.tfplan | tftools summarize --show-tags
  if [ -f "plan.tfplan" ]; then rm plan.tfplan; fi
}
```

> [!NOTE]
> Note that the bash function has the possibility of activating the output or not using the `-v` flag

```shell
tfsum terraform -v # show full output
```

> [!WARNING]
> Adapt the rest of zsh or fish functions according to your needs.

## Load new functions

```shell
source ~/.zshrc
source ~/.bashrc
source ~/.config/fish/config.fish
```

## Using tfsum as a custom binary

Copy [tfsum](../scripts/tfsum.sh) to `/usr/local/bin/tfsum`

```shell
sudo cp scripts/tfsum.sh /usr/local/bin/tfsum
```

> `/usr/local/bin` or other directory included in your path

Inside the container image of [infratools](https://github.com/containerscrew/infratools), this custom binary is already contemplated.

# Example

```shell
cd my-terraform-project/
tfsum terraform
```

Then, you will see the summarized output with the corresponding targets.

The example:

> [!NOTE]
> The following example is using the full output command

```shell
tftools summarize --show-tags --show-unchanged <Documents/plan.json
```

![example](../assets/example.png)

> Terragrunt is also supported

```shell
$ tfsum terragrunt
```

# JSON output support

## JSON output with arns
```shell
tftools summarize --json --pretty-json <demo.json
```

## JSON output only metrics

```shell
tftools summarize --json --metrics <demo.json
```

## JSON output pretty

```shell
tftools summarize --json --metrics --pretty-json <demo.json
```

![example-json-outputs](../assets/example-json-output.png)
