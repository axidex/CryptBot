run: swag run-server

swag:
	swag init -g ./cmd/main/main.go

run-server:
	go run cmd/main/main.go

restart-full: remove pull compose

remove:
	docker compose stop
	docker compose rm -f

pull:
	docker compose pull

compose:
	docker compose up -d