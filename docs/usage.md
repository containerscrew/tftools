<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Usage](#usage)
  - [Summarize](#summarize)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Usage

## Summarize

```bash
terraform plan -out plan.tfplan
terraform show -json plan.tfplan | tftools summarize
```

Or if you have the file in json

```bash
cat plan.json | tftools summarize
```
