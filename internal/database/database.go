package database

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	path string
	db   *sql.DB
}

func NewDB(path string) (DB, error) {
	// Ensure the directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return DB{}, err
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return DB{}, err
	}

	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		return DB{}, err
	}

	err = createTableMain(db)
	if err != nil {
		return DB{}, err
	}

	err = createMainEntry(db)
	if err != nil {
		return DB{}, err
	}

	return DB{
		path: path,
		db:   db,
	}, nil
}

func createTableMain(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS main (
			id INTEGER PRIMARY KEY,
			endtime INTEGER,
			urls TEXT
		)
	`)

	return err
}

func createMainEntry(db *sql.DB) error {
	_, err := db.Exec(`
	INSERT OR IGNORE INTO main
	(id, endtime, urls)
	VALUES
	(1, 0, '[]')
	`)

	return err
}

func (d *DB) CloseDB() {
	d.db.Close()
}

// endtime

func (d *DB) GetEndtime() (int64, error) {
	var endtime int64
	err := d.db.QueryRow(`
	SELECT endtime
	FROM main
	WHERE id=?
	`, 1).Scan(&endtime)
	return endtime, err
}

func (d *DB) SetEndtime(endtime int64) error {
	_, err := d.db.Exec(`
	UPDATE main
	SET endtime=?
	WHERE id=?
	`, endtime, 1)
	return err
}

// blocklist

func (d *DB) GetBlocklist() ([]string, error) {
	var urls string
	err := d.db.QueryRow(`
	SELECT urls
	FROM main
	WHERE id=?
	`, 1).Scan(&urls)
	if err != nil {
		return nil, err
	}

	return decodeStringSlice([]byte(urls))
}

func (d *DB) SetBlocklist(blocklist []string) error {
	encoded, err := encodeStringSlice(blocklist)
	if err != nil {
		return err
	}

	_, err = d.db.Exec(`
	UPDATE main
	SET urls=?
	WHERE id=?
	`, string(encoded), 1)
	return err
}

// helpers

func encodeStringSlice(data []string) ([]byte, error) {
	return json.Marshal(data)
}

func decodeStringSlice(data []byte) ([]string, error) {
	result := []string{}
	err := json.Unmarshal(data, &result)
	return result, err
}
