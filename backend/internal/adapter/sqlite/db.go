package sqlite

import (
	"database/sql"
	"fmt"
)

// Схема таблицы
const schema = `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author VARCHAR(100) NOT NULL DEFAULT 'Anonymous',
		message TEXT NOT NULL,
		response TEXT NOT NULL,
		date DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS comment_requests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author VARCHAR(100) NOT NULL DEFAULT 'Anonymous',
		message TEXT NOT NULL,
		telegram_id INTEGER NOT NULL,
		date DATETIME NOT NULL
	);
`

// Инициализация базы данных
func Init(path string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}
	if _, err := conn.Exec(schema); err != nil {
		return conn, fmt.Errorf("apply schema: %w", err)
	}
	return conn, nil
}

// Метод отключения базы данных
func Close(db *sql.DB) error {
	if db == nil {
		return nil
	}
	return db.Close()
}
