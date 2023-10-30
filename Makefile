.ONESHELL:
SHELL = /bin/bash
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

runlocal: ## Run's the Azure func locally
	@source ./scripts/build.sh && runLocal

publish: ## Publish to Azure functions
	@source ./scripts/build.sh && publish
