<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Usage](#usage)
  - [Summarize](#summarize)
  - [Function for ~/.zshrc](#function-for-zshrc)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Usage

## Summarize

```bash
terraform plan -out plan.tfplan
terraform show -json plan.tfplan | tftools summarize
```

Or if you have the file in json

```bash
terraform plan -out plan.tfplan
terraform show -json plan.tfplan > plan.json
cat plan.json | tftools summarize
```

## Function for ~/.zshrc

Copy [this function](../scripts/tfsum.sh) in your `~/.zshrc` or `~/.bashrc` file.

```bash
function tfsum() {
  if [ -z "$1" ];
  then
    echo "You should type 'tfsum terraform|terragrunt'"
  else
    echo -e "Starting tf summary..."
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
