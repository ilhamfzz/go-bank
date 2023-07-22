package component

import (
	"database/sql"
	"fmt"
	"log"

	"go-wallet.in/internal/config"

	_ "github.com/lib/pq"
)

func DatabaseConnection(cnf *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		cnf.Database.Host,
		cnf.Database.Port,
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.Name,
	)

	connection, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error when open connection %s", err.Error())
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalf("Error when ping connection %s", err.Error())
	}

	// Query to create table users
	_, err = connection.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		id integer NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		full_name VARCHAR(255) NOT NULL,
		phone VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		email_verified_at TIMESTAMP DEFAULT NULL);

		CREATE TABLE IF NOT EXISTS accounts (
		id integer NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		user_id integer NOT NULL,
		account_number VARCHAR(255) NOT NULL UNIQUE,
		balance numeric(10,2) NOT NULL);

		CREATE TABLE IF NOT EXISTS transactions (
		id integer NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		account_id integer NOT NULL,
		soft_number VARCHAR(255) NOT NULL,
		dof_number VARCHAR(255) NOT NULL,
		amount numeric(19,2) NOT NULL,
		transaction_type VARCHAR(1) NOT NULL,
		transaction_date TIMESTAMP NOT NULL);

		CREATE TABLE IF NOT EXISTS notifications (
		id integer NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		user_id integer NOT NULL,
		status integer NOT NULL,
		title VARCHAR(255) NOT NULL,
		body VARCHAR(255) NOT NULL,
		is_read integer NOT NULL DEFAULT 0,
		created_at TIMESTAMP default NULL);

		CREATE TABLE IF NOT EXISTS topups (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		user_id integer NOT NULL,
		amount integer NOT NULL,
		status integer NOT NULL DEFAULT 0,
		snap_url VARCHAR(255) NOT NULL);
		`)
	if err != nil {
		log.Fatalf("Error when create table users %s", err.Error())
	}

	return connection
}
