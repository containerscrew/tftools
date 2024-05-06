#!/bin/bash

set -eu

echo -en "\e[32m> Starting tfsum...\e[0m\n"

# Default variables
tf_command="terraform"
tf_action="plan"
tf_params='-parallelism=100'

# Function to display help text
usage() {
    echo "Usage: $0 [-c] [-a] [-p] [-h]"
    echo "Options:"
    echo "  -c           Command. Execute terragrunt or terraform. (Required)"
    echo "  -a           Action. Execute plan, apply or destroy. (Optional)(Default: plan)"
    echo "  -p           Params. Terraform params like: -auto-approve, -parallelism ...etc. (Optional)(Default: -parallelism=100)"
    echo "  -h           Display the help message"
    echo -en             "\e[32m> Example: $ tfsum -c terraform -a plan -p '-parallelism=50 -no-color'\e[0m\n"
}

# Parse options using getopts
while getopts "c:a:p:h" option; do
    case "${option}" in
        c)  # Set command
            tf_command=${OPTARG}
            if [ "${OPTARG}" != "terraform" ] && [ "${OPTARG}" != "terragrunt" ]; then
              echo -en "\e[31m> Invalid command ${OPTARG}. Supported commands are: terraform, terragrunt\e[0m\n"
              exit 1
            fi
            ;;
        a)  # Set action
            tf_action=${OPTARG}
            if [ "${OPTARG}" != "plan" ] && [ "${OPTARG}" != "apply" ] && [ "${OPTARG}" != "delete" ]; then
               echo -en "\e[31m> Invalid command. Supported commands are: plan, apply, delete\e[0m\n"
               exit 1
            fi
            ;;
        p) # Terraform params
            tf_params="${OPTARG//\'/}"
            ;;
        h)  # Help option
            usage
            exit 0
            ;;
        \?) # Invalid option
            echo "Invalid option: -${OPTARG}"
            usage
            exit 1
            ;;
    esac
done


# If no flags, print usage
if [ $# -eq 0 ]; then
    echo -en "\e[31m> You should type \$ tfsum -c terraform/terragrunt at least.\e[0m\n"
    usage
    exit 1
fi

# Start program
echo -en "\e[32m> Executing $tf_command $tf_action $tf_params\e[0m\n"
if [ "$tf_action" == "plan" ]; then
  tf_params+=" -out plan.tfplan"
  # shellcheck disable=SC2086
  $tf_command "$tf_action" $tf_params
  # Print summarized plan
  $tf_command show -json plan.tfplan | tftools summarize --show-tags
else
  # shellcheck disable=SC2086
  $tf_command "$tf_action" $tf_params
fi
