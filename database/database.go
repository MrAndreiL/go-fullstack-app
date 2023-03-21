package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var isCreated bool
var db *sql.DB

func Connect() {
	// Establish connection.
	var err error
	db, err = sql.Open("mysql", "user:user_password@tcp(127.0.0.1:6033)/app_db")
	if err != nil {
		fmt.Println("Error occurred when connecting to database.")
		panic(err.Error())
	}
	isCreated = true
	fmt.Println("Connection to database established successfully.")

	// Create tables if not created and seed them with default values.
	if err = seedDatabase(); err != nil {
		fmt.Println("Error occurred when seeding database with default values.")
		panic(err.Error())
	}
}

func seedDatabase() error {
	accounts_create := `CREATE TABLE IF NOT EXISTS accounts (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                                             name TEXT NOT NULL,
                                                             username TEXT NOT NULL,
                                                             email VARCHAR(50) NOT NULL,
                                                             password TEXT NOT NULL)`

	if _, err := db.Exec(accounts_create); err != nil {
		fmt.Println("Error occurred when creating table.")
		return err
	}

	// if table is empty, seed it with default values.
	res, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		fmt.Println("Query error")
		return err
	}

	if res.Next() == false {
		studentsInsert := `INSERT INTO accounts (name, username, email, password) VALUES (?, ?, ?, ?)`

		_, err := db.Exec(studentsInsert, "Lungu Andrei", "Andreul", "lunguandrei759@gmail.com", "testpass")
		if err != nil {
			fmt.Println("Account seed inserting error")
			return err
		}
	}
	return nil
}

func GetDbConnection() *sql.DB {
	if !isCreated {
		Connect()
	}
	return db
}

func CloseConnection() {
	err := db.Close()

	if err != nil {
		fmt.Println("Error occurred when closing database connection")
		panic(err.Error())
	}
	fmt.Println("Connection to database closed successfully.")
}
