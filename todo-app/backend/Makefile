
.PHONY: run
run:
	export GIN_MODE="debug"
	go run main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build:
	docker build -t demo-backend-app .

.PHONY: test
test:
	go test ./...

.PHONY: test-db
test-db:
	go test --tags=database ./...

.PHONY: up
up:
	docker compose up -d 

.PHONY: down
down:
	docker compose down -v

.PHONY: db-up
db-up:
	docker compose up -d db

.PHONY: db-dwon
db-down:
	docker compose rm -fsv db
