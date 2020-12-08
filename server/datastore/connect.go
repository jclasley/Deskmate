package datastore

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

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
	config = checkEnvVar()
	// Check ENV variables for connection details

	// If ENV variables don't exist for Postgres, write a message to the log
	// file and prompt the user
	if (ConnectionDetails{}) == config {
		promptForConnectionDetails()
	}
	// Open a connection using connection details
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.host, config.port, config.user, config.password, config.dbname)
	db, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Error opening connection to Postgres database:", err.Error())
		// TODO: Add logging for an error on database connection
	}

	// Ping database to ensure connection
	err = db.Ping()
	if err != nil {

		fmt.Println("Error pinging Postgres database:", err.Error())
		// TODO: Add logging for a failed ping, likely due to a connection issue
	}
	checkTable()
}

// checkEnvVar runs a quick check on the environment variables to see if ones
// relating to Deskmate and containing the Postgres connection details are
// present
func checkEnvVar() (config ConnectionDetails) {
	port, err := strconv.Atoi(os.Getenv("DESKMATE_PSQL_PORT"))
	if err != nil {
		// TODO: Add logging to capture error for port not being able to be
		// converted
	}
	config = ConnectionDetails{
		host:     os.Getenv("DESKMATE_PSQL_HOST"),
		port:     port,
		user:     os.Getenv("DESKMATE_PSQL_USER"),
		password: os.Getenv("DESKMATE_PSQL_PASS"),
		dbname:   os.Getenv("DESKMATE_PSQL_DBNAME"),
	}
	return config
}

// promptForConnectionDetails asks the user to enter the necessary connection
// details for establishing a Postgres connection if the environment variables
// do not exist
func promptForConnectionDetails() {
	fmt.Println("Postgres connection details not found. For Deskmate to run, please enter connection details for a Postgres database.")
	fmt.Println("Please enter the Postgres host")
	fmt.Scanln(&config.host)

	fmt.Println("Please enter the port for Postgres")
	fmt.Scanln(&config.port)

	fmt.Println("Please enter the username for Postgres")
	fmt.Scanln(&config.user)

	fmt.Println("Please enter the password for Postgres")
	fmt.Scanln(&config.password)

	fmt.Println("Please enter the database name")
	fmt.Scanln(&config.dbname)
}
