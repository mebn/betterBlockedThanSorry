package main

import (
	"database/sql"
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

	createTableRuns(db)
	createTableURLS(db)

	addURL(db, "a")
	addURL(db, "b")
	addURL(db, "c")

	start(db, 9991)
	start(db, 9992)
	start(db, 9993)

	printRuns(db)
	printURLS(db)
}

func createTableRuns(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS runs (
			id INTEGER PRIMARY KEY,
			starttime INTEGER NOT NULL,
			endtime INTEGER NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func createTableURLS(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY,
			url TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func addURL(db *sql.DB, url string) {
	_, err := db.Exec(`
		INSERT INTO urls
		(url)
		VALUES
		(?)
	`, url)
	if err != nil {
		log.Printf("Failed to insert row: %v", err)
	}
}

func start(db *sql.DB, endtime int64) {
	_, err := db.Exec(`
		INSERT INTO runs
		(starttime, endtime)
		VALUES
		(?, ?)
	`, time.Now().Unix(), endtime)
	if err != nil {
		log.Printf("Failed to insert row: %v", err)
	}
}

func printRuns(db *sql.DB) {
	rows, err := db.Query(`SELECT id, starttime, endtime FROM runs`)
	if err != nil {
		log.Fatalf("Failed to query rows: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var starttime int64
		var endtime int64

		err := rows.Scan(&id, &starttime, &endtime)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}

		// Print the row
		fmt.Printf("ID: %d, Starttime: %d, Endtime: %d\n", id, starttime, endtime)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during row iteration: %v", err)
	}
}

func printURLS(db *sql.DB) {
	rows, err := db.Query(`SELECT id, url FROM urls`)
	if err != nil {
		log.Fatalf("Failed to query rows: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var url string

		err := rows.Scan(&id, &url)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}

		// Print the row
		fmt.Printf("ID: %d, url: %s\n", id, url)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during row iteration: %v", err)
	}
}
