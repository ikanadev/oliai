# usage migrate.sh up|down

export DATABASE=postgresql://kevv:postgres@0.0.0.0:5432/olidb?sslmode=disable
migrate -database ${DATABASE} -path db/migrations $1
