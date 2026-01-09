# Personal Tracker (Go + Vue 3 + PrimeVue)

A small, GitHub-ready full-stack project:
- **Backend:** Go (Gin) + SQLite
- **Frontend:** Vue 3 + Pinia + PrimeVue (Aura theme)
- **Features:** Daily tasks, checkbox completion, add/delete, progress.

## Screens
Add your screenshots here after running locally.

---

## Requirements
- Go **1.22+**
- Node **18+** (recommended)

---

## Run (Local)

### 1) Backend
```bash
cd backend
go mod tidy
go run ./cmd/api
```
Backend runs at: `http://localhost:8080`  
SQLite file is created at: `backend/tracker.db`

### 2) Frontend
```bash
cd frontend
npm install
npm run dev
```
Frontend runs at: `http://localhost:5173`

---

## API Endpoints
- `GET /api/tasks?day=YYYY-MM-DD`
- `POST /api/tasks`  `{ "title": "...", "day": "YYYY-MM-DD" }`
- `PATCH /api/tasks/:id/toggle`
- `DELETE /api/tasks/:id`
- `GET /api/summary/week?start=YYYY-MM-DD`

---

## Notes
- CORS is configured to allow `http://localhost:5173`.
- Minimal, clean structure to keep the repo small but realistic.

## Roadmap ideas
- Categories (Meals/Skincare/Habits)
- Weekly chart (PrimeVue Chart)
- Editing task titles
- Auth (optional)
