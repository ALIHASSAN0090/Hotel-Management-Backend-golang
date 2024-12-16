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

.PHONY: setup-local
setup-local:
	go run db_migrator/main.go db_migrator/migrations.go
	./init.sh
	go mod tidy
	go run main.go
