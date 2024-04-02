#!/bin/bash

if [ -z "$1" ]; then
  echo "You should type 'tfsum terraform|terragrunt'"
  exit 1
fi

echo -en "Starting tf summary... Please wait"

if [ -n "$2" ] && [ "$2" == "-v" ]; then
  "$1" plan -out plan.tfplan
else
  "$1" plan -out plan.tfplan 1> /dev/null
fi

"$1" show -json plan.tfplan | tftools summarize --show-tags
if [ -f "plan.tfplan" ]; then rm plan.tfplan; fi