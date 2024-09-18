run_db:
	docker compose --profile db up -d
run_dev:
	docker compose --profile backend up -d --build

