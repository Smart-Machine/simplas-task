default: build run

build:
	@docker compose up -d
	@chown -R $(id -u):$(id -g) certs

run: build
	@go run cmd/app/main.go

clean:
	@sudo rm -rfv certs
	@docker compose down -v

gen-proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		service/proto/service.proto