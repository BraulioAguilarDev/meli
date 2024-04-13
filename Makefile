.PHONY: docker-build

docker-build:
	@echo building
	docker build --no-cache -t meli .
	@echo Done