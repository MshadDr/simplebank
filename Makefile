postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRE_USER=root -e POSTGRE_PASSWORD=root -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_database

dropdb:
	docker exec -it postgres dropdb simple_database

migrateup:
	migrate -path ./database/migrations -database "postgresql://root:root@localhost:5432/simple_database?sslmode=disable" -verbose up

migratedown:
	migrate -path ./database/migrations -database "postgresql://root:root@localhost:5432/simple_database?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY:createdb dropdb migrateup migratedown sqlc
