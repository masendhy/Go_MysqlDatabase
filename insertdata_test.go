package mysqlgodatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	data := "INSERT INTO customer(id,name,email,balance,rating,birth_date,married) VALUES ('2','sendhy','sendhy@gmail.com',1000,5.0,'1999-10-11',true),('3','axa','axa@gmail.com',100000,5.0,'1999-10-10',false)"
	_, err := db.ExecContext(ctx, data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succes insert data")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	data := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, data)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestManyQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	data := "SELECT id,name,email,balance,rating, created_at,birth_date,married FROM customer"
	rows, err := db.QueryContext(ctx, data)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var created_at, birth_date time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Created at:", created_at)
		fmt.Println("Birth Date:", birth_date)
		fmt.Println("Married:", married)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin"
	password := "admin"

	sqlquery := "SELECT username FROM user WHERE username =? AND password =? LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlquery, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success Login", username)
	} else {
		fmt.Println("Login Failed")
	}

}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "masendhy@gmail.com"
	text := "helo3 bro"

	autoincrement := "INSERT INTO comments (email,comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, autoincrement, email, text)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	insert := "INSERT INTO comments (email,comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, insert)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "masendhy" + strconv.Itoa(i) + "gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	insert := "INSERT INTO comments (email,comment) VALUES(?,?)"

	for i := 0; i < 10; i++ {
		email := "masendhy" + strconv.Itoa(i) + "gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, insert, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id", id)
	}
	tx.Rollback()
}
