package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

// SetupDB initializes the database and creates the necessary tables if they do not exist
func SetupDB() *sql.DB {
	dsn := "username:password@tcp(localhost:3306)/dbname?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Create the TODO table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
        id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,     
        task TEXT,
        completed BOOLEAN
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateToDo(task string, db *sql.DB) error {

	insertSQL := `insert into todos(task, completed) values (?, false)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = statement.Exec(task)
	return err
}

type Todo struct {
	ID        int
	task      string
	completed bool
}

func getAllToDos(db *sql.DB) ([]Todo, error) {

	selectAllSQL := `select * from todos`
	rows, err := db.Query(selectAllSQL)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo

		if err := rows.Scan(&todo.ID, &todo.task, &todo.completed); err != nil {
			log.Fatal(err)
			return nil, err
		}

		todos = append(todos, todo)
	}
	return todos, nil

}

func updateToDo(ID int, db *sql.DB) error {
	updateSQL := `update todos set completed=1 where ID=?`
	statement, err := db.Prepare(updateSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = statement.Exec(ID)

	return err
}

func deleteToDo(ID int, db *sql.DB) error {
	deleteSQL := `delete from todos where ID=?`
	statement, err := db.Prepare(deleteSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = statement.Exec(ID)

	return err
}

func main() {
	//Connect to database
	db := SetupDB()
	defer db.Close()

	//Create a TODO
	task := fmt.Sprintf("%s%d", "TODO-", rand.Intn(100))
	CreateToDo(task, db)

	//Get all TODOs
	todos, err := getAllToDos(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todos)

	//Update a TODO - mark it completed(1)
	updateToDo(2, db)
	updateToDo(3, db)
	//todos=nil
	todos, err = getAllToDos(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todos)

	//Delete a TODO
	deleteToDo(1, db)
	todos, err = getAllToDos(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todos)

}
