# VibeCheck API

> A mood and status sharing REST API — post your daily vibe, follow people, react to their vibes, and see what's trending.

Built as **Project 2 of 3** in a structured backend engineering curriculum. The primary goals are to solidify all Phase 2 backend patterns from Project 1 (ThreadPulse) and to learn six new production concepts through dedicated checkpoints built into the project milestones.

---

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go |
| Framework | Gin |
| Database | PostgreSQL (via sqlx + golang-migrate) |
| Cache | Redis (via go-redis) |
| Auth | JWT (golang-jwt) |
| Containerisation | Docker + Docker Compose |
| CI | GitHub Actions |
| Deployment | Railway / Render |

---

## Features

- **JWT Auth** — Register and login with bcrypt-hashed passwords
- **Post a Vibe** — One vibe per user per day, enforced by a DB unique constraint on `(user_id, date)`
- **Follow System** — Follow and unfollow users; vibes from followed users appear in your feed
- **Feed with Pagination** — Cursor/offset-paginated feed of vibes from people you follow
- **React to Vibes** — Reactions are processed by a background worker with two outcomes: reaction count update + notification creation
- **Delete Vibe** — Cascade deletion of reactions and notifications wrapped in a single DB transaction
- **Redis Caching** — Cache-aside strategy with TTL and targeted invalidation
- **Trending Vibes** — Most-reacted vibes in the last 24 hours via SQL `ORDER BY` + time filter + pagination
- **Rate Limiting** — Redis-based rate limiting on the react endpoint
- **Structured Logging** — JSON-formatted logs throughout using Go's built-in `slog`
- **Graceful Shutdown** — OS signal listener drains in-flight requests before exit
- **CORS Middleware** — Configured via `gin-contrib/cors` for live frontend compatibility
- **Integration Tests** — Full flow test (post vibe → react → verify notification in DB) running in CI

---

## New Production Concepts Learned Here

| Concept | Where Applied |
|---------|--------------|
| Context & Timeouts | Every DB call has a 3-second deadline |
| Database Transactions | Delete vibe cascade — atomic, rolls back on failure |
| Structured Logging (slog) | Every significant action logged as structured JSON |
| Graceful Shutdown | `main.go` listens for SIGTERM, drains requests cleanly |
| CORS Middleware | Configured for live deployment URL |
| Integration Tests | Full post → react → notify flow tested against real test DB |

---

## Getting Started

### Prerequisites

- Go 1.22+
- Docker and Docker Compose

### Run Locally

```bash
git clone https://github.com/deeep8250/vibecheck-api.git
cd vibecheck-api

cp .env.example .env
# Fill in your secrets in .env

docker compose up --build
```

The API will be available at `http://localhost:8080`.

### Run Migrations

```bash
migrate -path db/migrations -database "postgres://user:password@localhost:5432/vibecheck?sslmode=disable" up
```

---

## Environment Variables

```env
DB_URL=postgres://user:password@localhost:5432/vibecheck?sslmode=disable
REDIS_URL=redis://localhost:6379
JWT_SECRET=your_secret_here
PORT=8080
```

---

## API Endpoints

### Auth

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/api/v1/auth/register` | Register a new user | ❌ |
| POST | `/api/v1/auth/login` | Login and receive JWT | ❌ |

### Vibes

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/api/v1/vibes` | Post today's vibe (text + mood tag + emoji) | ✅ |
| GET | `/api/v1/vibes/:id` | Get a specific vibe | ✅ |
| DELETE | `/api/v1/vibes/:id` | Delete your vibe (cascades reactions + notifications) | ✅ |

### Feed

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/api/v1/feed` | Paginated feed of vibes from followed users | ✅ |

### Follow

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/api/v1/users/:id/follow` | Follow a user | ✅ |
| DELETE | `/api/v1/users/:id/follow` | Unfollow a user | ✅ |

### Reactions

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/api/v1/vibes/:id/react` | React to a vibe (rate limited) | ✅ |

### Trending

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/api/v1/trending` | Most-reacted vibes in the last 24 hours, paginated | ✅ |

---

## Architecture

The project follows a strict **3-layer architecture** throughout:

```
Handler  →  Service  →  Repository
  (HTTP)     (logic)      (DB/Redis)
```

- **Handlers** parse requests and write responses — no business logic
- **Services** own all business logic — no direct DB access
- **Repositories** own all DB and Redis access — no HTTP awareness

### Background Worker

Reactions are processed asynchronously via a Go channel-based worker pool. Each reaction job has two outcomes:

1. Increment the reaction count on the vibe
2. Create a notification record for the vibe owner

Both writes are independent (not transactional) — the count and notification are best-effort async updates. Context with timeout is passed to every DB call inside the worker.

### Database Transaction (Delete Vibe)

Deleting a vibe performs three operations atomically using `sqlx.BeginTxx`:

1. Delete all reactions for the vibe
2. Delete all notifications for the vibe
3. Delete the vibe itself

If any step fails, the entire transaction is rolled back.

---

## Project Structure

```
vibecheck-api/
├── cmd/api/main.go              # Entry point — server setup, graceful shutdown
├── internal/
│   ├── handler/                 # HTTP layer
│   ├── service/                 # Business logic layer
│   ├── repository/              # DB + Redis layer
│   ├── middleware/              # Auth, rate limiting, CORS
│   ├── worker/                  # Background reaction worker
│   ├── models/                  # Structs (User, Vibe, Reaction, Notification)
│   └── config/                  # Env config loader
├── db/migrations/               # golang-migrate SQL files
├── tests/integration/           # Integration tests (real DB)
├── .github/workflows/ci.yml     # GitHub Actions CI
├── docker-compose.yml
├── Dockerfile
└── .env.example
```

---

## CI / CD

GitHub Actions runs on every push to `main`:

1. Start PostgreSQL and Redis services
2. Run migrations against the test DB
3. Run unit tests
4. Run integration tests (post vibe → react → verify notification in DB)

Deployed to **Railway** (or Render) — live URL added here after first deployment.

---

## Part of the Phase 2 Backend Roadmap

| Project | Focus | Status |
|---------|-------|--------|
| ThreadPulse | Learn all Phase 2 concepts | ✅ Complete |
| **VibeCheck** | **Solidify + learn missing production concepts** | 🔨 In Progress |
| Nexus | Resume showstopper — own everything | ⏳ Upcoming |
