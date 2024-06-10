# Updated: <>

export

help:
	@echo "Usage:"
	@printf "  %s\n" 'PASSWORD="xxx" make leak  ## insert PASSWORD into code and git commit/push'
	@printf "  %s\n" 'make leak                 ## insert random PASSWORD into code and git commit/push'
	@printf "  %s\n" 'make clear                ## no secrets in code and git commit/push'
	@echo

leak:
	bash simulate-developer-commit-push-secret.sh $(PASSWORD)

clear:
	bash simulate-developer-clear-secret-commit-push.sh

.PHONY: Makefile help leak clear
