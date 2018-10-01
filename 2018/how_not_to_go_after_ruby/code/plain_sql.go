package dao

import (
	"database/sql"
	"log"
	"fmt"
)
// START OMIT
func Get(age int)  {
	var db *sql.DB
	rows, _ := db.Query("SELECT name FROM users WHERE age=?", age)
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
// END OMIT
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}