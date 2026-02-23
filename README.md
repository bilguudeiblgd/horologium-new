# Horologium

## Stack

- **Backend**: Go + Gin, port `8080`
- **Frontend**: Next.js 16 + Tailwind CSS, port `3000`
- **Database**: PostgreSQL 17, port `5432`

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) with Docker Compose

## Development

Start all services with hot reload:

```bash
docker compose -f docker-compose.dev.yml up --build
```

- Frontend: http://localhost:3000
- Backend: http://localhost:8080

## Run migrations

```bash
docker compose exec -T postgres psql -U postgres -d db < backend/migrations/001_create_users.sql
```

## Production

```bash
docker compose up --build
```

## Connect to the database

```bash
docker compose exec postgres psql -U postgres -d db
```
