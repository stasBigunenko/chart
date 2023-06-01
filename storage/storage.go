package storage

import (
	"chart/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Storage struct
type Storage struct {
	db *sql.DB
}

func New(connStr config.ConfigDB) (*Storage, error) {
	db, err := sql.Open("postgres", connStr.ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database %w\n", err)
	}

	database := &Storage{db: db}

	return database, nil
}

func (s *Storage) GetDB() *sql.DB {
	return s.db
}

func (s *Storage) Close() {
	s.db.Close()
}
