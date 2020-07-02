package db

import (
	"database/sql"
	"fmt" // pq library for Postgres DB
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

var (
	db *sql.DB
)

// Init : Initializes the PostgreSQL database
func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	panic(err)
	// util.PanicErr(err)

	err = db.Ping()
	panic(err)
	// util.PanicErr(err)

	fmt.Println("Successfully connected!")
}

// GetInstance : Gets the initiated db instance
func GetInstance() *sql.DB {
	return db
}

// Close : Closes the db connection
func Close() {
	db.Close()
}

// Test : Tests the db conn
func Test() {
	rows, err := db.Query("SELECT * FROM test")
	// util.PanicErr(err)
	panic(err)

	for rows.Next() {
		var id int
		var sensorName string
		err = rows.Scan(&id, &sensorName)
		panic(err)
		fmt.Println("id | sensorName")
		fmt.Printf("%3v | %8v\n", id, sensorName)
	}
}
