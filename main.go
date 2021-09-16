package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connection user=gummy789j")
	if err != nil {
		log.Fatal(fmt.Sprintf("Uable to connect: %v/n", err))
	}

	defer conn.Close()

	log.Println("Connected to database!")

	// test my connection (ping the database)
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database\n")
	}

	log.Println("Ping database!")
	// get rows from table

	err = getAllRows(conn)
	if err != nil {
		log.Fatal("Get rows error\n")
	}

	// insert a row
	query := `insert into users (first_name, last_name) values ($1, $2)`

	_, err = conn.Exec(query, "Jack", "Brown")

	if err != nil {
		log.Fatal(err)
	}

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal("Get rows error\n")
	}

	// update a row
	stmt := `update users set first_name = $1 where id = $2`

	_, err = conn.Exec(stmt, "Jackie", 5)
	if err != nil {
		log.Fatal(err)
	}

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal("Get rows error\n")
	}

	// get a row by id
	query = `select id, first_name, last_name from users where id = $1`
	row := conn.QueryRow(query, 1)

	var firstName, lastName string

	var id int

	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("QueryRow return", id, firstName, lastName)

	// delete a row
	query = `delete from users where id = $1`
	_, err = conn.Exec(query, 6)
	if err != nil {
		log.Fatal(err)
	}

	// get rows again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal("Get rows error\n")
	}
}

func getAllRows(conn *sql.DB) error {

	rows, err := conn.Query("select id, first_name, last_name from users")

	if err != nil {
		log.Println(err)
		return err
	}

	defer rows.Close()

	var firstName, lastName string

	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("--------------------------------------")

	return nil
}
