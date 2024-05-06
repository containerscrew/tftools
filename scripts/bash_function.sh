tfsum() {
  (
    # Enable pipefail within the subshell
    set -o pipefail

    # Create plan and pass through any arguments
    # Make a random tfplan filename in /tmp
    TMP_FILE=$(mktemp /tmp/tfplan.XXXXXX)

    # Execute terraform plan and other commands
    terraform plan -lock=false -compact-warnings -out=${TMP_FILE} "$@" |
      # Remove the line mentioning where the plan was saved
      awk '!/Saved the plan to/{print;next} /Saved the plan to/{exit}' &&
        terraform show -json ${TMP_FILE} |
          tftools summarize --show-tags --show-unchanged --compact &&
            rm ${TMP_FILE}
  )
}
