.PHONY: help run lint test
.DEFAULT_GOAL: help
default: help

help: ## Output available commands
	@echo "Available commands:"
	@echo
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run: ## Run the proxy service
	@docker-compose build service
	@docker-compose run --rm --service-ports service

test: ## Run all tests
	@docker-compose build test
	@docker-compose run --rm test ./scripts/test.sh
