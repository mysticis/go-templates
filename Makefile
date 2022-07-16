migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/tasks?sslmode=disable" -verbose up  
migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/tasks?sslmode=disable" -verbose down
server:
	go run main.go
.PHONY: migrateup migratedown server