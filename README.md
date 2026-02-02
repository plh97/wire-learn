# wire-learn

Small Gin + GORM (SQLite) sample project with a simple user lookup API.

## Learn Steps

1. Manually build a Gin + GORM project using global variables.
2. Refactor to manual dependency wiring (no globals).
3. Use a simple `wire` package to centralize construction and injection.

## Manual Wiring (step by step)

1. **Create core providers**
 - `core.NewDB()` opens SQLite and runs migrations.
 - `core.NewLogger()` builds the app logger.
2. **Create services**
 - `service.NewUserService(db, logger)` depends on DB + logger.
3. **Create handlers**
 - `api.NewUserApi(userService)` depends on the service.
4. **Create router**
 - `router.NewRouter(userApi)` registers routes with handlers.
5. **Start server**
 - `main.go` calls `wire.InitWire()` and then `Run()`.

This keeps dependencies explicit and avoids package-level globals.

## How `wire` Works Here

This project uses a small manual wiring function (not Google Wire).

- `wire.InitWire()` is a **composition root**.
- It constructs dependencies in order (DB → Logger → Service → API → Router).
- It passes instances down the chain so each layer receives what it needs.

Flow:

```
NewDB → NewLogger → NewUserService → NewUserApi → NewRouter
```

## Requirements

- Go 1.25+

## Install

```bash
go mod tidy
```

## Run

```bash
go run main.go
```

Server starts on `:80`.

## API

- **GET** `/api/user?id=<id>`

Example:

```bash
curl "http://localhost:80/api/user?id=1"
```

## Project Structure

```
api/        HTTP handlers
core/       DB and logger setup
model/      GORM models
router/     Gin routes
service/    Business logic
wire/       Manual wiring/initialization
```

## Notes

- SQLite database file: `db.test`
- If a user ID does not exist, the API returns `404`.
