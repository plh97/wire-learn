# Wire Learn Tutorial

A minimal Go + Gin + GORM tutorial project demonstrating dependency injection with Google Wire. Includes a simple User API backed by SQLite.

## Tech Stack
- Go
- Gin
- GORM (SQLite)
- Google Wire
- Zap

## Project Structure
- api/: HTTP handlers
- core/: shared providers (DB, logger)
- model/: data models
- router/: routes setup
- service/: business logic
- wire/: Wire injector

## Prerequisites
- Go 1.20+ (or your local version)
- Wire installed: `go install github.com/google/wire/cmd/wire@latest`

## Install Dependencies
```bash
go mod tidy
```

## Step-by-step Tutorial
This project is a progression from global variables to manual wiring, then to Google Wire.

### 1) Start with global variables (baseline)
Goal: get things working quickly.

- Create global singletons (DB, logger) in a package like `core/`.
- Use them directly in services and handlers.

Example (concept only):
- `core` exposes `DB` and `Logger`
- `service` and `api` use `core.DB` / `core.Logger` directly

Downside: tight coupling, hard to test, hidden dependencies.

### 2) Remove globals with manual wiring
Goal: make dependencies explicit and testable.

- Add constructors:
	- `core.NewDB()`
	- `core.NewLogger()`
	- `service.NewUserService(db, logger)`
	- `api.NewUserApi(userService)`
	- `router.NewUserRouter(api)`
- Wire everything in `main.go` manually.

Now each layer gets its dependencies via parameters (no globals).

### 3) Introduce Google Wire for auto injection
Goal: avoid repetitive manual wiring.

1. Create a Wire injector in `wire/wire.go` (build tag `wireinject`).
2. Register providers with `wire.Build(...)`:
	 - `core.NewDB`
	 - `core.NewLogger`
	 - `service.NewUserService`
	 - `api.NewUserApi`
	 - `router.NewUserRouterProvider`
3. Run Wire to generate `wire_gen.go`:

```bash
wire ./wire
```

Wire will create `wire_gen.go` with the compiled dependency graph and provide an `InitWire()` that returns `*gin.Engine`.

### 4) Use the generated injector
- In `main.go`, call `wire.InitWire()` instead of manual wiring.

That’s it—dependencies are explicit, testable, and automatically assembled.

## Generate Wire
```bash
wire ./wire
```

## Run
```bash
go run .
```

## API
### Get User
`GET /api/user?id=1`

Example:
```bash
curl "http://localhost:8080/api/user?id=1"
```

## Notes
- SQLite database file: `test.db`
- Update providers in `wire/` when adding new services.
