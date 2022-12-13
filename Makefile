postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRE_USER=root -e POSTGRE_PASSWORD=root -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_database

dropdb:
	docker exec -it postgres dropdb simple_database

migrateup:
	migrate -path ./database/migrations -database "postgresql://root:root@localhost:5432/simple_database?sslmode=disable" -verbose up

migrateup1:
	migrate -path ./database/migrations -database "postgresql://root:root@localhost:5432/simple_database?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./database/migrations -database "postgresql://root:root@localhost:5432/simple_database?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./database/migrations -database "postgresql://root:root@localhost:5432/simple_database?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover -count=1 ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination database/mock/store.go simplebank/database/sqlc Store

.PHONY:createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
