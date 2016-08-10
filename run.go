package main

import "fmt"
import _ "github.com/go-sql-driver/mysql"
import "database/sql"
import "os"

func main() {
	db := getDb()
    rows, err := db.Query("SELECT name  FROM some-table")

    if err != nil {
    	fmt.Println("You have some trouble with SQL Sintaxe");
    	panic(err)
    }

    for rows.Next() {
    	var name string
    	_ = rows.Scan(&name)
    	stmt, err := db.Prepare("INSERT INTO <INSTRUCTION> values(?,91)")
    	if (err != nil) {
    		fmt.Println(err)
    		fmt.Printf("%s has not Migrated \n", name)
    		continue
    	}

    	_, err = stmt.Exec(name)
    	if (err != nil) {
    		fmt.Println(err)
    		fmt.Printf("%s has not Migrated \n", name)
    		continue
    	}

    	fmt.Printf("%s has Migrated \n", name)
    }    
}

func getDb() (*sql.DB) {
    db, err := sql.Open("mysql", "root:root@tcp(<HOSTNAME>:3306)/<DBNAME>?charset=utf8")

	if (err != nil) {
		panic(err)
	}

	return db
}