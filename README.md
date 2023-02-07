# Golang starter

## primary things
1. go environment with [fiver](https://github.com/gofiber/fiber)
2. gin
2. gorm
3. api
4. config
5. migrations

### issues
issue: Extension exists but uuid_generate_v4 fails
fix:
```SQL
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### migrations
```shell
go run migrate/migrate.go
```

### running at local
1/ install [Air - Live reload for Go apps](https://github.com/cosmtrek/air)

2/ run
```shell
air
```

### linter
```shell
golangci-lint run
```