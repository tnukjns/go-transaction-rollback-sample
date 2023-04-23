# go-transaction-rollback-sample

This Go program demonstrates the importance of handling transactions and their rollback in SQLite, particularly when a transaction is canceled. The program connects to an SQLite database, creates a table called users, and then attempts to execute two SQL queries in a single transaction. During the transaction, the context is intentionally canceled between the first and second queries, causing the transaction to be rolled back immediately, emphasizing the need for proper transaction handling in cases of cancellation.

## Requirements

- Go (version 1.15 or higher)
- SQLite3 Go driver (github.com/mattn/go-sqlite3)

## Installation

1. Install Go on your machine by following the instructions on the official Go website.
2. Install the SQLite3 Go driver by running the following command:
```
go get -u github.com/mattn/go-sqlite3
```

## Usage

1. Clone this repository and navigate to the project directory.
2. Compile and run the program with the following command:
```
go run main.go
```

Upon successful execution, the program will connect to the example.db SQLite database file, create a users table, and attempt to insert two rows within a transaction. However, the context is intentionally canceled between the first and second queries, causing the transaction to be rolled back immediately, and an error message will be logged. This highlights the importance of handling transaction rollbacks in a database application, especially when dealing with cancellations.
