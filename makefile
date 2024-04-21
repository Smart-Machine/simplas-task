default: build run

build:
	@docker compose up -d
	@chown -R $(id -u):$(id -g) certs

run: build
	@go run cmd/app/main.go

clean:
	@sudo rm -rfv certs
	@docker compose down -v