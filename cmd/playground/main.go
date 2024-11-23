package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "/tmp/mydb.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Printf("Failed to set journal mode: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS main (
			id INTEGER PRIMARY KEY,
			starttime INTEGER NOT NULL,
			endtime INTEGER NOT NULL,
			blocklist TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	add(db, 9991, []string{"a", "b"})
	add(db, 9992, []string{"b"})
	add(db, 9993, []string{"a"})

	printRows(db)
}

func add(db *sql.DB, endtime int64, blocklist []string) {
	blocklistJSON, err := json.Marshal(blocklist)
	if err != nil {
		log.Printf("Failed to serialize blocklist: %v", err)
		return
	}

	_, err = db.Exec(`
		INSERT INTO main
		(starttime, endtime, blocklist)
		VALUES
		(?, ?, ?)
	`, time.Now().Unix(), endtime, blocklistJSON)
	if err != nil {
		log.Printf("Failed to insert row: %v", err)
	}
}

func printRows(db *sql.DB) {
	rows, err := db.Query(`SELECT id, starttime, endtime, blocklist FROM main`)
	if err != nil {
		log.Fatalf("Failed to query rows: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var starttime int64
		var endtime int64
		var blocklistJSON string

		err := rows.Scan(&id, &starttime, &endtime, &blocklistJSON)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}

		// Deserialize the blocklist
		var blocklist []string
		err = json.Unmarshal([]byte(blocklistJSON), &blocklist)
		if err != nil {
			log.Printf("Failed to deserialize blocklist: %v", err)
			continue
		}

		// Print the row
		fmt.Printf("ID: %d, Starttime: %d, Endtime: %d, Blocklist: %v\n", id, starttime, endtime, blocklist)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during row iteration: %v", err)
	}
}
