.PHONY: run

# IMPROVEMENT: Set safe fallback defaults if the user leaves variables blank
concurrent ?= 1
api_url ?= 
api_key ?= 

run:
	@# Construct the execution command dynamically based on what was provided
	@CMD="go run main.go --concurrent=$(concurrent)"; \
	if [ -n "$(api_url)" ]; then CMD="$$CMD --api_url=$(api_url)"; fi; \
	if [ -n "$(api_key)" ]; then CMD="$$CMD --api_key=$(api_key)"; fi; \
	$$CMD
