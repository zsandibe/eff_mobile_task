create-migrations:
	migrate create -ext sql -dir ./migrations -seq init_schema

migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/effective?sslmode=disable" -verbose up 

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@127.0.0.1:5432/effective?sslmode=disable" -verbose down

run:
	go run cmd/main.go