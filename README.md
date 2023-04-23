# go-transaction-sample

This Go program demonstrates the importance of handling transactions and their rollback in SQLite. The program connects to an SQLite database, creates a table called users, and then attempts to execute two SQL queries in a single transaction. If an error occurs during the execution of any of the queries, the transaction is immediately rolled back, emphasizing the need for proper transaction handling.

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

Upon successful execution, the program will connect to the example.db SQLite database file, create a users table, and attempt to insert two rows within a transaction. If an error occurs during the execution of any of the SQL queries, the transaction will be rolled back immediately, and an error message will be logged. This highlights the importance of handling transaction rollbacks in a database application.
