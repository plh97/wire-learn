# wire-learn

Small Gin + GORM (SQLite) sample project with a simple user lookup API.

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
