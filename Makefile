.PHONY: run

run:
	@go run main.go \
		--concurrent=$(concurrent) \
		--api_url=$(api_url) \
		--api_key=$(api_key)
