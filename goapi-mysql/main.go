package main

import (
	"database/sql"
	"fmt"
	"log"

	// _ meaning not using package directly, just want to exe init func of this package

	_ "github.com/go-sql-driver/mysql"
)

func checkSQLError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

type Data struct {
	id   int
	name string
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DBname)
	db, err := sql.Open("mysql", connectionString)
	checkSQLError(err)
	defer db.Close()

	// insert Code
	result, err := db.Exec("INSERT INTO data VALUES(4, 'pqr')")
	checkSQLError(err)
	lastInsertedId, err := result.LastInsertId()
	fmt.Println("Last Inserted Id: ", lastInsertedId)
	checkSQLError(err)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("Rows Affected: ", rowsAffected)
	checkSQLError(err)

	rows, err := db.Query("SELECT * FROM data")
	checkSQLError(err)
	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkSQLError(err)
		fmt.Println(data)
	}
}
