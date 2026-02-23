# AyoIndo Football Management API

A RESTful JSON API built with **Go (Golang)** + **GIN** framework for managing football teams, players, matches, and results.

---

## Tech Stack

| Layer      | Technology              |
|------------|-------------------------|
| Language   | Go 1.24+                |
| Framework  | GIN                     |
| ORM        | GORM (with soft delete) |
| Database   | PostgreSQL (Supabase) |
| Auth       | JWT (HS256)             |
| Password   | bcrypt                  |

---

## Project Structure

```
ayoindo/
├── main.go
├── go.mod / go.sum
├── .env
├── config/
│   └── database.go
├── models/
│   ├── user.go
│   ├── team.go
│   ├── player.go
│   ├── match.go
│   ├── match_result.go
│   └── goal.go
├── handlers/
│   ├── auth_handler.go
│   ├── team_handler.go
│   ├── player_handler.go
│   ├── match_handler.go
│   ├── result_handler.go
│   └── report_handler.go
├── middleware/
│   └── auth.go
├── routes/
│   └── routes.go
└── utils/
    └── response.go
```

---

## Setup & Running

### 1. Prerequisites
- Go 1.24+
- A [Supabase](https://supabase.com) project (or any PostgreSQL instance)
- The `.env` file configured with your database credentials (see step 2)

### 2. Configure environment

Create a `.env` file in the project root based on the template below.
For **Supabase**, use the **Transaction Pooler** connection string (port `6543`):

```env
# Supabase (Transaction Pooler — port 6543)
DB_HOST=aws-1-ap-southeast-2.pooler.supabase.com
DB_PORT=6543
DB_USER=postgres.<your-project-ref>
DB_PASSWORD=<your-supabase-db-password>
DB_NAME=postgres

# Auth
JWT_SECRET=<your-strong-secret-min-32-chars>

# Server
GIN_MODE=debug
PORT=8080
```

> ⚠️ **Never commit `.env` to Git.** It is already listed in `.gitignore`.

> ℹ️ For a local PostgreSQL instance instead, set `DB_HOST=localhost`, `DB_PORT=5432`, `DB_NAME=ayoindo_db`, and `sslmode` can be changed to `disable` in `config/database.go`.

```bash
go run main.go
```

The server starts at `http://localhost:8080`. Tables are auto-migrated on startup.

### 5. Build binary

```bash
go build -o ayoindo main.go
./ayoindo
```

---

## Authentication

All endpoints except `POST /api/auth/register` and `POST /api/auth/login` require a **Bearer JWT token**.

```
Authorization: Bearer <token>
```

---

## API Endpoints

### Health Check

| Method | Path      | Description       |
|--------|-----------|-------------------|
| GET    | `/health` | Server health check |

---

### Authentication

| Method | Path                  | Auth | Description       |
|--------|-----------------------|------|-------------------|
| POST   | `/api/auth/register`  | ❌   | Register new admin |
| POST   | `/api/auth/login`     | ❌   | Login & get JWT   |
| GET    | `/api/auth/me`        | ✅   | Get current user profile |

#### Register
```json
POST /api/auth/register
{
  "username": "admin",
  "email": "admin@example.com",
  "password": "secret123"
}
```

#### Login
```json
POST /api/auth/login
{
  "email": "admin@example.com",
  "password": "secret123"
}
```
**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGci...",
    "user": { "id": 1, "username": "admin", "email": "admin@example.com" }
  }
}
```

---

### Teams

| Method | Path             | Auth | Description           |
|--------|------------------|------|-----------------------|
| GET    | `/api/teams`     | ✅   | List all teams        |
| POST   | `/api/teams`     | ✅   | Create team           |
| GET    | `/api/teams/:id` | ✅   | Get team (with players)|
| PUT    | `/api/teams/:id` | ✅   | Update team           |
| DELETE | `/api/teams/:id` | ✅   | Soft-delete team      |
| GET    | `/api/teams/:id/players` | ✅ | Get players of team |

**Query params for GET /api/teams:** `?city=Jakarta`

#### Create / Update Team Body
```json
{
  "name": "Persija Jakarta",
  "logo": "https://example.com/persija.png",
  "founded_year": 1928,
  "address": "Jl. Menteng Raya No.1",
  "city": "Jakarta"
}
```

---

### Players

| Method | Path               | Auth | Description          |
|--------|--------------------|------|----------------------|
| GET    | `/api/players`     | ✅   | List all players     |
| POST   | `/api/players`     | ✅   | Create player        |
| GET    | `/api/players/:id` | ✅   | Get player detail    |
| PUT    | `/api/players/:id` | ✅   | Update player        |
| DELETE | `/api/players/:id` | ✅   | Soft-delete player   |

**Query params for GET /api/players:** `?team_id=1`, `?position=penyerang`

#### Create / Update Player Body
```json
{
  "team_id": 1,
  "name": "Bambang Pamungkas",
  "height": 175.0,
  "weight": 68.5,
  "position": "penyerang",
  "jersey_number": 20
}
```

**Valid positions:** `penyerang`, `gelandang`, `bertahan`, `penjaga_gawang`

> ⚠️ Jersey numbers must be unique within a team.

---

### Matches

| Method | Path               | Auth | Description              |
|--------|--------------------|------|--------------------------|
| GET    | `/api/matches`     | ✅   | List all matches         |
| POST   | `/api/matches`     | ✅   | Create match schedule    |
| GET    | `/api/matches/:id` | ✅   | Get match detail         |
| PUT    | `/api/matches/:id` | ✅   | Update match schedule    |
| DELETE | `/api/matches/:id` | ✅   | Soft-delete match        |

**Query params for GET /api/matches:** `?status=scheduled`, `?status=completed`

#### Create / Update Match Body
```json
{
  "home_team_id": 1,
  "away_team_id": 2,
  "match_date": "2025-03-15",
  "match_time": "19:30"
}
```

---

### Match Results

| Method | Path                     | Auth | Description              |
|--------|--------------------------|------|--------------------------|
| POST   | `/api/matches/:id/result`| ✅   | Submit / update result   |
| GET    | `/api/matches/:id/result`| ✅   | Get match result         |

#### Submit Match Result Body
```json
{
  "home_score": 2,
  "away_score": 1,
  "goals": [
    { "player_id": 5, "minute": 23 },
    { "player_id": 5, "minute": 67 },
    { "player_id": 12, "minute": 45 }
  ]
}
```

> ⚠️ Number of goals per team must equal the reported score.  
> ⚠️ Each player must belong to one of the two teams.  
> Submitting again to the same match **replaces** the existing result.

---

### Reports

| Method | Path                      | Auth | Description                     |
|--------|---------------------------|------|---------------------------------|
| GET    | `/api/reports/matches`    | ✅   | Summary of all completed matches|
| GET    | `/api/reports/matches/:id`| ✅   | Detailed report for one match   |

#### Detailed Report Response
```json
{
  "success": true,
  "data": {
    "match_id": 1,
    "match_date": "2025-03-15",
    "match_time": "19:30",
    "home_team": { "id": 1, "name": "Persija Jakarta" },
    "away_team": { "id": 2, "name": "Arema FC" },
    "home_score": 2,
    "away_score": 1,
    "final_status": "Tim Home Menang",
    "goals": [
      { "player_id": 5, "player": { "name": "Bambang" }, "minute": 23 },
      { "player_id": 5, "player": { "name": "Bambang" }, "minute": 67 },
      { "player_id": 12, "player": { "name": "Singo" }, "minute": 45 }
    ],
    "top_scorers": [
      { "player_id": 5, "player_name": "Bambang", "goals": 2 }
    ],
    "home_team_total_wins": 5,
    "away_team_total_wins": 3
  }
}
```

**`home_team_total_wins`** = cumulative all-time wins for the home team (as home or away) across all completed matches up to and including this match.  
**`away_team_total_wins`** = same for the away team.

---

## Business Rules

1. **One player, one team** — a player can only be registered to one team.
2. **Unique jersey per team** — two players in the same team cannot share a jersey number.
3. **Home ≠ Away** — a team cannot play against itself.
4. **Goal validation** — goal count must match home/away score; goals are only valid for players from the two competing teams.
5. **Soft delete** — all `DELETE` endpoints use GORM soft delete (`deleted_at` timestamp). Records remain in the DB but are excluded from all queries.
6. **JWT expiry** — tokens expire after **24 hours**.

---

## Standard Response Format

```json
{
  "success": true | false,
  "message": "...",
  "data": { ... }
}
```

List responses include a `"total"` field.
