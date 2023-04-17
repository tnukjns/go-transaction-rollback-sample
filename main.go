package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// context key for storing transaction ID
type transactionIDKey struct{}

func main() {
	// connect to SQLite database
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// create a table
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// generate a transaction ID (e.g., UUID)
	transactionID := "12345678-1234-5678-1234-567812345678"

	// add the transaction ID to the context
	ctx := context.WithValue(context.Background(), transactionIDKey{}, transactionID)

	// set a cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// execute the first SQL query
	_, err = tx.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {
		handleRollback(ctx, tx, err)
		return
	}

	// cancel in the middle
	cancel()

	// execute the second SQL query
	_, err = tx.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", "Bob")
	if err != nil {
		handleRollback(ctx, tx, err)
		return
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("transaction with ID %s committed successfully", transactionID)
}

func handleRollback(ctx context.Context, tx *sql.Tx, err error) {
	if rbErr := tx.Rollback(); rbErr != nil {
		log.Fatalf("rollback failed: %v", rbErr)
	}
	transactionID, _ := ctx.Value(transactionIDKey{}).(string)
	log.Printf("transaction with ID %s rolled back due to: %v", transactionID, err)
}
