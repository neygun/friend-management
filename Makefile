build:
	docker build -t fm-app .

run:
	docker-compose --env-file ./.env up

migrate-up:
	docker compose run --rm migrate sh -c "migrate -path /migration -database postgres://postgres:postgres@database:5432/fm-pg?sslmode=disable up"

migrate-drop:
	docker compose run --rm migrate sh -c "migrate -path /migration -database postgres://postgres:postgres@database:5432/fm-pg?sslmode=disable drop"

run-dev:
	go run main.go
