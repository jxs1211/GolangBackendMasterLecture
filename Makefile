dbup:
	docker compose up -d postgres
dbdown:
	docker compose down 
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres dropdb simple_bank
migrateup:
	migrate -path db/migration/ -database postgres://root:root@localhost:5432/simple_bank?sslmode=disable up
migratedown:
	migrate -path db/migration/ -database postgres://root:root@localhost:5432/simple_bank?sslmode=disable down
test:
	go test -v -cover -short ./...
server:
	go run main.go

.PHONY: dbup dbdown createdb dropdb migrateup migratedown test server
