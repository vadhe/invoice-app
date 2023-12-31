# invoice-app

## Getting Started

install package

- [Templ](https://templ.guide/)
- [Air](https://github.com/cosmtrek/air)


```bash
# Run docker
docker compose up -d

# Install go depedencies 
go mod tidy

# Create migration 
goose -dir=migrations create name_table sql

# Run migration
goose -dir=migrations postgres "user= dbname= password= sslmode=disable host=localhost port=5432" up

# Run seed
go run ./cmd/seed/main.go

# Run Development server
templ generate --watch
pnpm build-css
air
```
