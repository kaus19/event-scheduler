postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=root -d postgres:17-alpine 

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root event_scheduler

dropdb:
	docker exec -it postgres17 dropdb event_scheduler

migrateup:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/event_scheduler?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:password@localhost:5432/event_scheduler?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server