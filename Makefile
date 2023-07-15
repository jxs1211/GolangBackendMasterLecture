DATABASESOURCE=postgres://root:root@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postpres --network bank-network -e POSTGRES_DB=simple_bank -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 postgres:12-alpine

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

dbup:
	docker compose up -d postgres

dbdown:
	docker compose down 

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database $(DATABASESOURCE) up

migratedown:
	migrate -path db/migration/ -database $(DATABASESOURCE) down

test:
	go test -v -cover -short ./...

server:
	docker compose up

.PHONY: network postgres mysql dbup dbdown createdb dropdb migrateup migratedown test server
