export PORT=4000
export JWT_SIGN_KEY=secret
export DATABASE=postgresql://kevv:postgres@0.0.0.0:5432/olidb?sslmode=disable
export MIGRATIONS_URL=github://ikanadev/oliai/oliapi/db/migrations
export QDRANT_URL=127.0.0.1:6334
export OPEN_AI_KEY=myapikey

go run cmd/api/api.go
