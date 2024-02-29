export PORT=4000
export JWT_SIGN_KEY=secret
export DATABASE=postgresql://kevv:postgres@0.0.0.0:5432/olidb?sslmode=disable
export MIGRATIONS_URL=github://ikanadev/oliai/oliapi/db/migrations

go run cmd/api/api.go
