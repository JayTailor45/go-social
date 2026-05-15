# Go Social

A social media project to learn go

# Migrations

- CREATE:

```
migrate create -seq -ext sql -dir ./cmd/migrate/migrations <name>
```

- UP:

```
migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" up
```

- DOWN:

```
migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/social?sslmode=disable" down
```

# Seed data

```
go run cmd/migrate/seed/main.go
```

# Run project

```
go run cmd/api/*.go
```
