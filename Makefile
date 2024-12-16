.PHONY: run
run:
	go run main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: user
user:
	./init.sh

.PHONY: migrate-up
migrate-up:
	go run db_migrator/main.go db_migrator/migrations.go

.PHONY: setup-docker
setup-docker:
	docker compose up --build -d

.PHONY: down-docker
down-docker:
	docker compose down

.PHONY: rebuild-docker
rebuild-docker:
	docker compose down
	docker compose up --build -d

.PHONY: logs
logs:
	docker compose logs -f

.PHONY: clean
clean:
	docker compose down -v
	docker system prune -f
	docker volume rm $(docker volume ls -q)

.PHONY: setup-local
setup-local:
	go run db_migrator/main.go db_migrator/migrations.go
	./init.sh
	go mod tidy
	go run main.go