package datastore

import (
	"database/sql"
	"fmt"
	"os"

	// pq is the postgres driver for the sql package
	_ "github.com/lib/pq"
)

// ConnectionDetails represents the pieces needed to open a connection
// to Postgres, including the host, port, user, password and dbname
type ConnectionDetails struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// config is a given instance of the connection details needed for Postgres
var (
	config ConnectionDetails
	db     *sql.DB
	err    error
)

// ConnectPostgres establishes the connection between Deskmate and
// a local Postgres database. The connection details for Deskmate
// should be loaded from environment variables. If those variables
// aren't present, Deskmate should ask for those details at the command
// line. From there, the connection details are pulled together to open a
// connection with Postgres, and then ping the database to ensure the
// connection is active.
func ConnectPostgres() {
	envVars := retrieveEnvConfig()
	config = ConnectionDetails{
		host:     "itd-deskmate.db.infra.circleci.com",
		port:     5432,
		user:     envVars["POSTGRES_USER"],
		password: envVars["POSTGRES_PASSWORD"],
		dbname:   envVars["POSTGRES_DB"],
	}

	// Open a connection using connection details
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", config.host, config.port, config.user, config.password, config.dbname)
	db, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Error opening connection to Postgres database:", err.Error())
		return
	}

	// Ping database to ensure connection
	err = db.Ping()
	if err != nil {

		fmt.Println("Error pinging Postgres database:", err.Error())
		return
	}
	fmt.Println("Successfully connected to Postgres")
	checkTable()
}

func retrieveEnvConfig() (map[string]string) {
	env := make(map[string]string)
	env["POSTGRES_USER"] = os.Getenv("POSTGRES_USER")
	env["POSTGRES_PASSWORD"] = os.Getenv("POSTGRES_PASSWORD")
	env["POSTGRES_DB"] = os.Getenv("POSTGRES_DB")
	return env
}