postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://:@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://:@localhost:5432/simplebank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://:@localhost:5432/simplebank?sslmode=disable" -verbose down 1

migrateup1:
	migrate -path db/migration -database "postgresql://:@localhost:5432/simplebank?sslmode=disable" -verbose up 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/AzfarInan/go-masterclass/simplebank/db/sqlc Store

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test server mock migrateup1 migratedown1