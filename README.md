# ⚽ League Simulation API

A backend API written in Go for simulating a football league,  
including fixture generation, week-by-week simulation, league table, and champion probability prediction.

---

## 📚 Table of Contents

- [Features](#features)  
- [Project Structure](#project-structure)  
- [Installation & Run](#installation--run)  
- [Database Schema](#database-schema)  
- [API Endpoints](#api-endpoints)  
- [Example API Requests](#example-api-requests)  
- [Testing](#testing)  
- [Notes](#notes)

---

## 🚀 Features

- Team and match entities with dynamic stats  
- Round-robin fixture generation (supports 4 teams)  
- Week-by-week match simulation with realistic goal calculation  
- League table with advanced tie-breaker rules (points, goal difference, head-to-head, goals for, name)  
- Monte Carlo–based champion probability calculation and DB storage  
- PostgreSQL database support (persistent data)  
- Full RESTful API with Gin  
- Easy reset for new seasons (without resetting team IDs)  
- Modular, layered architecture (entity, repository, service, controller/api)

---

## 🗂️ Project Structure

```text
entities/           # Data structures (Team, Match, ChampionProbability)
repository/         # Database and memory repositories, interfaces
service/            # Business logic, simulation, probability, statistics
controller/         # API/HTTP handlers (Gin)
main.go             # Program entry (Gin setup and router)
```

---

## 🛠️ Installation & Run

1. **Install Go 1.22+ and PostgreSQL.**  
2. **Clone the repository:**  
   ```bash
   git clone https://github.com/<your-username>/league-simulation.git
   cd league-simulation
   ```  
3. **Install dependencies:**  
   ```bash
   go mod download
   ```  
4. **Configure your database connection** in `repository/db.go`  
   ```go
   const POSTGRES_DSN = "postgres://user:password@localhost:5432/yourdb?sslmode=disable"
   ```  
5. **Create your database schema** (see below).  
6. **Run the server:**  
   ```bash
   go run main.go
   ```  
   The API will be available at: `http://localhost:8080`

---

## Database Schema

```sql
CREATE TABLE IF NOT EXISTS public.teams
(
    id integer NOT NULL,
    name varchar(30) NOT NULL,
    strength integer DEFAULT 0,
    matchesplayed integer DEFAULT 0,
    matcheswon integer DEFAULT 0,
    matchesdrawn integer DEFAULT 0,
    matcheslost integer DEFAULT 0,
    goalsfor integer DEFAULT 0,
    goalsagainst integer DEFAULT 0,
    goaldifference integer DEFAULT 0,
    points integer DEFAULT 0,
    CONSTRAINT teams_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.matches
(
    id integer NOT NULL DEFAULT nextval('matches_id_seq'::regclass),
    week integer NOT NULL,
    hometeamid integer,
    awayteamid integer,
    homegoals integer DEFAULT 0,
    awaygoals integer DEFAULT 0,
    played boolean DEFAULT false,
    CONSTRAINT matches_pkey PRIMARY KEY (id),
    CONSTRAINT matches_awayteamid_fkey FOREIGN KEY (awayteamid)
        REFERENCES public.teams (id),
    CONSTRAINT matches_hometeamid_fkey FOREIGN KEY (hometeamid)
        REFERENCES public.teams (id)
);

CREATE TABLE IF NOT EXISTS public.champion_probabilities
(
    id integer NOT NULL DEFAULT nextval('champion_probabilities_id_seq'::regclass),
    team_id integer,
    season integer NOT NULL,
    probability real NOT NULL,
    calculated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT champion_probabilities_pkey PRIMARY KEY (id),
    CONSTRAINT champion_probabilities_team_id_fkey FOREIGN KEY (team_id)
        REFERENCES public.teams (id)
);
```

---

## API Endpoints

All endpoints return `application/json`.

1. **Get All Teams**  
   ```http
   GET /teams
   ```  
   **Response:**
   ```json
   [
     {
       "id": 1,
       "name": "Team A",
       "strength": 80,
       "matchesPlayed": 0,
       "matchesWon": 0,
       "matchesDrawn": 0,
       "matchesLost": 0,
       "goalsFor": 0,
       "goalsAgainst": 0,
       "goalDifference": 0,
       "points": 0
     }
     // ...
   ]
   ```

2. **Get League Fixture**  
   ```http
   GET /fixture
   ```  
   **Response:**
   ```json
   [
     {
       "id": 1,
       "week": 1,
       "homeTeamID": 1,
       "awayTeamID": 2,
       "homeGoals": 2,
       "awayGoals": 1,
       "played": true
     }
     // ...
   ]
   ```

3. **Get League Table**  
   ```http
   GET /league-table
   ```  
   **Response:** List of all teams, sorted by tie-breaker rules.

4. **Create Fixture**  
   ```http
   POST /create-fixture
   ```  
   **Response:**
   ```json
   { "message": "Fixture created!" }
   ```

5. **Simulate Week**  
   ```http
   POST /simulate-week/:week
   ```  
   **Param:** `week` (int)  
   **Response:**
   ```json
   [
     {
       "id": 5,
       "week": 2,
       "homeTeamID": 3,
       "awayTeamID": 4,
       "homeGoals": 1,
       "awayGoals": 1,
       "played": true
     }
     // ...
   ]
   ```

6. **Get Champion Probabilities (in-memory)**  
   ```http
   GET /champion-probabilities
   ```  
   **Response:**
   ```json
   {
     "Team A": 0.72,
     "Team B": 0.11,
     "Team C": 0.09,
     "Team D": 0.08
   }
   ```

7. **Get Champion Probabilities (from DB)**  
   ```http
   GET /champion-probabilities-db?season=2024
   ```  
   **Param:** `season` (int)  
   **Response:**
   ```json
   [
     {
       "id": 1,
       "teamID": 1,
       "season": 2024,
       "probability": 0.72,
       "calculatedAt": "2024-06-01T15:22:00Z"
     }
     // ...
   ]
   ```

8. **Reset Teams and Champion Probabilities**  
   ```http
   POST /reset-teams
   ```  
   **Response:**
   ```json
   { "message": "All teams' stats reset successfully!" }
   ```

**General Error Response:**
```json
{
  "error": "Error message"
}
```

---

## 💡 Example API Requests

- **Create a fixture:**
  ```bash
  curl -X POST http://localhost:8080/create-fixture
  ```
- **Simulate week 1:**
  ```bash
  curl -X POST http://localhost:8080/simulate-week/1
  ```
- **Get league table:**
  ```bash
  curl http://localhost:8080/league-table
  ```
- **Get champion probabilities:**
  ```bash
  curl http://localhost:8080/champion-probabilities
  ```
- **Reset the teams' stats:**
  ```bash
  curl -X POST http://localhost:8080/reset-teams
  ```

---

## 🧪 Testing

Run all tests:
```bash
go test ./test
```

---

## 📝 Notes

- All endpoints return JSON.  
- Business logic is decoupled via interface-based repository pattern.  
- The API is stateless; no session/cookie required.  
- Currently only 4 teams supported for fixture creation (can be extended).
