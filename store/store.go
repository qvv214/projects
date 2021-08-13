package store

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func (s *Store) Open() error {
	_, err := toml.DecodeFile("configs.toml", &s.config)
	if err != nil {
		log.Fatalf("fatal decode: %v", err)
	}
	conn, err := sql.Open("postgres", s.config.DatabaseConfig)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	s.db = conn

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) AddUrl(long string, short string) {
	var strId string

	max := s.db.QueryRow("select max(id) from url")
	max.Scan(&strId)

	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Fatalf("fatal strconv: %v", err)
	}

	id++

	query := "insert into url(id,long_name,short_name) values($1,$2,$3)"
	if err != nil {
		log.Fatalf("fatal prepare: %v", err)
	}

	_, err = s.db.Exec(query, id, long, short)
	if err != nil {
		log.Fatalf("fatal insert into: %v", err)
	}
}
