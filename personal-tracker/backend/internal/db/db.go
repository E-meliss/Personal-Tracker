package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func Open(path string) (*sql.DB, error) {
    database, err := sql.Open("sqlite3", path)
    if err != nil {
        return nil, err
    }

    if err := migrate(database); err != nil {
        _ = database.Close()
        return nil, err
    }

    return database, nil
}

func migrate(database *sql.DB) error {
    _, err := database.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL,
            day TEXT NOT NULL,              -- YYYY-MM-DD
            is_done INTEGER NOT NULL DEFAULT 0,
            created_at TEXT NOT NULL DEFAULT (datetime('now'))
        );

        CREATE INDEX IF NOT EXISTS idx_tasks_day ON tasks(day);
    `)
    return err
}
