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

## Development

The backend API is configured to accept requests from the Angular frontend running on port 4200.