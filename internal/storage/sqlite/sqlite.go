package sqlite

import (
	"database/sql"

	"github.com/bkjonathan/students-go-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	Db *sql.DB
}

func NewSQLiteStorage(cfg *config.Config) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		age INTEGER NOT NULL
	)`)
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{Db: db}, nil
}

func (s *SQLiteStorage) SaveStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
