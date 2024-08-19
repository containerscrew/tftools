SHELL:=/bin/sh
.PHONY: all

help: ## this help
	@awk 'BEGIN {FS = ":.*?## ";  printf "Usage:\n  make \033[36m<target> \033[0m\n\nTargets:\n"} /^[a-zA-Z0-9_-]+:.*?## / {gsub("\\\\n",sprintf("\n%22c",""), $$2);printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

doctoc: ## Create table of contents with doctoc
	doctoc .

goreleaser: ## Generate go binaries using goreleaser (brew install goreleaser)
	goreleaser release --snapshot --clean -p 2

golangci-lint: ## Lint Golang code (brew install golangci-lint)
	golangci-lint run --fix

pre-commit: ## Run pre-commit
	pre-commit run -a

go-generate: ## Run go generate
	go generate ./...

gosec: ## Run gosec
	gosec -exclude=G104,G204,G107 ./...

update-dependencies: ## Update dependencies
	go get -u ./...

generate-changelog: ## Generate changelog using git cliff
	git cliff --output CHANGELOG.md

##https://github.com/moovweb/gvm
##go get golang.org/x/tools/cmd/goimports
##go install golang.org/x/tools/cmd/goimports
