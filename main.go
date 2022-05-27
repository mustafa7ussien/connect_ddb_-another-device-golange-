package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Go MySQL Tutorial")

	// if there is an error opening the connection, handle it

	pswd := "P@ssword123"
	db, err := sql.Open("mysql", "user2:"+pswd+"@tcp(192.168.43.191:3306)/hotel_db")
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	fmt.Println("opened")
	fmt.Println(pswd)
	// 	// perform a db.Query insert

	insert, err := db.Query("INSERT INTO hotel VALUES (1,'Mostafa',3,'Assuit')")
	// 	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

}
