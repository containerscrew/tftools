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
