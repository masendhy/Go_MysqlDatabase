package mysqlgodatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gomysql_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// let's try to insert our go_test table on gomysql_database
	insert, err := db.Query("INSERT INTO go_test VALUES (1,'John','TEST')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
