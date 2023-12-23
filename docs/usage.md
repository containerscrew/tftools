<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Usage](#usage)
  - [Summarize](#summarize)
  - [Function for zsh or bash shell](#function-for-zsh-or-bash-shell)
- [Function for fish shell](#function-for-fish-shell)
- [Example](#example)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Usage

## Summarize

```bash
terraform plan -out plan.tfplan
terraform show -json plan.tfplan | tftools summarize
```

Or if you have the file in json:

```bash
terraform plan -out plan.tfplan
terraform show -json plan.tfplan > plan.json
cat plan.json | tftools summarize
tftools summarize ;
```

## Function for zsh or bash shell
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
    echo -e "\n"
    $1 show -json plan.tfplan | tftools summarize
    # Delete plan out file to avoid git tracking (although is included in .gitignore)
    if [ -f "plan.tfplan" ]; then rm plan.tfplan; fi
  fi
}
```

# Function for fish shell

```shell
function tfsum
    if test -z $argv[1]
        echo "You should type 'tfsum terraform|terragrunt'"
    else
        echo -e "Starting tf summary..."
        # Don't print output of terraform plan
        # If you don't want full plan output: $argv[1] plan -out plan.tfplan > /dev/null
        $argv[1] plan -out plan.tfplan
        echo -e "\n"
        $argv[1] show -json plan.tfplan | tftools summarize
        # Delete plan out file to avoid git tracking (although is included in .gitignore)
        if test -f "plan.tfplan"; rm plan.tfplan; end
    end
end
```

Load new functions:

```shell
source ~/.zshrc
source ~/.bashrc
source ~/.config/fish/config.fish
```

# Example

```shell
cd my-terraform-project/
tfsum terraform
```

Then, you will see full plan/apply of terraform and the summarized output with the corresponding targets.

The example:

![example](../assets/example.png)

> Terragrunt is also allowed

> [!NOTE]
> If using a pipeline, probably you will not want to see all the output. Update the [tfsum functiojn](//scripts/tfsum.sh) as you need.

Take a look to the comment:

```bash
....
# If you don't want full plan output: $1 plan -out plan.tfplan 1> /dev/null
$1 plan -out plan.tfplan
....
```
