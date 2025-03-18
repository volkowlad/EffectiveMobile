run:
	docker compose --env-file .env -f docker-compose.yml up

down:
	docker compose -f docker-compose.yml down