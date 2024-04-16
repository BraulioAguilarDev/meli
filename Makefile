.PHONY: docker dc-up dc-down

docker:
	@echo building
	docker build --no-cache -t meli .
	@echo Done

dc-up:
	docker compose -f docker-compose.yml up -d

dc-down:
	docker compose -f docker-compose.yml down
	docker volume rm meli-tech_pg_meli_data
