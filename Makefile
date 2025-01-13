run: swag run-server

swag:
	swag init -g ./cmd/main/main.go

run-server:
	go run cmd/main/main.go

restart-full: remove pull compose

remove:
	docker compose stop
	docker compose rm -f -v

pull:
	docker compose pull

compose:
	docker compose up -d

tidy:
	go mod tidy
	go fmt ./...