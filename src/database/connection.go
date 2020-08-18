package database

import(
	"database/sql"
	_"github.com/lib/pq"
	"log"
)

func GetConnection() *sql.DB{
	connStr := "postgres://postgres:yourPassword@localhost/todos_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		log.Fatal(err)
	}

	return db
}