package db

import (
  "database/sql"
  "os"

  "github.com/mattn/go-sqlite3"
)

const dbFile = "db/data/sqlite"

type SQLite struct {
  conn *sql.DB
}

func NewSQLite() (*SQLite, error) {
  // Open a database connection to a sqlite db file.
  conn, err := sql.Open("sqlite3", dbFile)
  if err != nil {
    return nil, err
  }

  // Create tables if the database file did not previously exist.
  if _, err = os.Stat(dbFile); os.IsNotExist(err) {
    _, err = conn.Exec("CREATE TABLE digests(digest TEXT PRIMARY KEY, message TEXT NOT NULL);")
    if err != nil {
      os.Remove(dbFile)
      return nil, err
    }
  }

  return &SQLite{conn: conn}, nil
}

func (s *SQLite) Insert(digest *Digest, msg string) error {
  // Insert into the sqlite database. Unique constraint errors are ignored because the digest is already in the database
  // so to the user it is the same as a successful insert.
  _, err := s.conn.Exec("INSERT INTO digests(digest, message) VALUES(?, ?);", digest.Hex, msg)
  if err != nil {
    sqlite3Err := err.(sqlite3.Error)
    if sqlite3Err.Code != sqlite3.ErrConstraint {
      return err
    }
  }
  return nil
}

func (s *SQLite) Get(digest string) (*Message, error) {
  var msg string
  err := s.conn.QueryRow("SELECT message FROM digests WHERE digest = ?", digest).Scan(&msg)
  if err != nil {
    return nil, err
  }
  return &Message{Text: msg}, nil
}
