migrate-sql:
	 migrate -source ./models_sql -database postgres://postgres:1234@localhost:5432/travel?sslmode=disable up

migration-down:
	migrate -path ./api/models_sql/ -database "postgresql://postgres:1234@localhost:5432/travel?sslmode=disable" -verbose down

migration-up:
	migrate -path ./api/models_sql/ -database "postgresql://postgres:1234@localhost:5432/travel?sslmode=disable" -verbose up

gen-swag:
	swag init -g ./api/api.go -o ./api/docs

run:
	go run cmd/main.go

