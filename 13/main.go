package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_NAME     = "members"
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_HOST     = "localhost"
)

func main() {
	fmt.Println("data persistence")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Try to ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	// _, err = db.Exec(`
	// 	CREATE TABLE members (
	// 		id SERIAL PRIMARY KEY,
	// 		name VARCHAR(255) NOT NULL,
	// 		email VARCHAR(255) NOT NULL
	// 	);
	// `)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	id := 1
	query := fmt.Sprintf("SELECT id,name, email FROM members WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		// var id int
		// var name, email string

		var member Members
		//		err := rows.Scan(&id, &name, &email)
		err := rows.Scan(&member.id, &member.name, &member.email)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", member.id, member.name, member.email)

	}

}

type Members struct {
	id    int
	name  string
	email string
}
