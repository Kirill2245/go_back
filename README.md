# BackToDo

Монорепозиторий приложения для управления задачами и подзадачами.

## Структура

```
.
├── backend/    # Go + Echo + GORM, REST API и Swagger
├── frontend/   # Vue 3 + Vite SPA
├── docker-compose.yaml
└── .env.example
```

## Backend

Требования: Go 1.27+ или Docker.

### Запуск через Docker Compose

```bash
cp .env.example .env
docker compose up --build
```

API будет доступен на `http://localhost:8080`, Swagger UI — на `http://localhost:8080/swagger`.

### Локальный запуск

```bash
cd backend
go run .
```

Для локального запуска нужна PostgreSQL. Строка подключения и разрешённый origin задаются переменными окружения `DATABASE_URL` и `ALLOWORIGIN` (см. `.env.example`).

## Frontend

Требования: Node.js 20+.

```bash
cd frontend
npm install
npm run dev
```

Приложение откроется на `http://localhost:5173` и будет обращаться к API на `http://localhost:8080`. Перед запуском фронтенда backend должен быть запущен.

Сборка:

```bash
cd frontend
npm run build
npm run preview
```

## API

- `GET /get-tasks` — список задач;
- `POST /create-task` — создание задачи;
- `PATCH /update/{id}` — обновление задачи;
- `DELETE /delete-task/{id}` — удаление задачи;
- `GET /get-small-tasks/{taskID}` — список подзадач задачи;
- `POST /create-small-task` — создание подзадачи;
- `PATCH /update-small-task/{id}` — обновление подзадачи;
- `DELETE /delete-small-task/{id}` — удаление подзадачи.
