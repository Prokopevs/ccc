build:
	go build github.com/Prokopevs/ccc/game/cmd/server
swagger:
	swag init -g ./cmd/server/main.go
migrateup:
	migrate -path schema -database "postgres://postgres:...@localhost:5432/gamedb?sslmode=disable" -verbose up
migratedown:
	migrate -path schema -database "postgres://postgres:...@localhost:5432/gamedb?sslmode=disable" -verbose down