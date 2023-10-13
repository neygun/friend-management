run:
	@docker-compose up fm-app

migrate-up:
	docker compose run --rm migrate sh -c "migrate -path /migration -database postgres://postgres:postgres@database:5432/fm-pg?sslmode=disable up"

migrate-drop:
	docker compose run --rm migrate sh -c "migrate -path /migration -database postgres://postgres:postgres@database:5432/fm-pg?sslmode=disable drop"

run-dev:
	go run main.go
