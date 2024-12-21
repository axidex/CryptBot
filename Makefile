run: swag run-server

swag:
	swag init -g ./cmd/main/main.go

run-server:
	go run cmd/main/main.go

re-pull:
	docker compose stop
	docker compose rm -f
	docker compose up -d --build