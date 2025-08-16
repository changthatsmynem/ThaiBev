# ThaiBev

Full-stack application with Angular frontend and Go backend.

## Project Structure

```
ThaiBev/
├── frontend/thaibev-app/    # Angular application
└── backend/thaibev-api/     # Go API server
```

## Quick Start

### Backend (Go API)
```bash
cd backend/thaibev-api
go mod tidy
go run main.go
```
Server runs on `http://localhost:8080`

### Frontend (Angular)
```bash
cd frontend/thaibev-app
npm install
npm start
```
Application runs on `http://localhost:4200`

## Docker Compose (Recommended)

### Start Entire Application
```bash
docker-compose up --build
```
This command starts the complete application stack:
- PostgreSQL database with auto-created tables (`localhost:5432`)
- Go backend API (`localhost:8080`)
- Angular frontend (`localhost:4200`)

### Stop Application
```bash
docker-compose down
```

### Database Only
```bash
docker-compose up postgres
```

### Reset Database (Fresh Start)
```bash
docker-compose down -v
docker-compose up --build
```
Use `-v` flag to remove database volumes and ensure fresh database initialization.

## Development

The backend API is configured to accept requests from the Angular frontend running on port 4200.

### Database Notes
- Database tables are auto-created via `init.sql` and Go fallback
- If tables are missing, restart with `docker-compose down -v && docker-compose up --build`