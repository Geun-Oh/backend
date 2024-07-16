test:
	go test -v -cover ./...

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose up