export PORT=4000
export JWT_SIGN_KEY=secret
export DATABASE=postgresql://kevv:postgres@0.0.0.0:5432/oliai?sslmode=disable

go run cmd/api/api.go
