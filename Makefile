default: run

run:
	go run src/main.go

docker:
	docker compose up -d 