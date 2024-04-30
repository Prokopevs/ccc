build:
	go build github.com/Prokopevs/ccc/users/cmd/server

migrateup:
	migrate -path schema -database "postgres://postgres:...@localhost:5432/gameDB?sslmode=disable" -verbose up

migratedown:
	migrate -path schema -database "postgres://postgres:...@localhost:5432/gameDB?sslmode=disable" -verbose down


