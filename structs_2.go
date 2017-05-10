package main

import (
	"database/sql"
	"fmt"
	"os"
)

// START OMIT

// structs can have methods associated with them

type User struct {
	Name string
	Age  int
	// note that the exported/unexported rule works the same way for structs as well!
}

// Save saves the user in a given database
// sql support is also in the standard library,
// but it's an abstraction. The implementation for a given database
// comes from a 3rd party driver.
func (u *User) Save(db *sql.DB) error {
	return db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", u.Name, u.Age)
}

func main() {
	db, err := sql.Open("postgres", "...")
	if err != nil {
		fmt.Println("error connecting to database:", err)
		os.Exit(1)
	}

	u := &User{
		Name: "John",
		Age:  29,
	}

	if err := u.Save(db); err != nil {
		fmt.Println("error saving the user:", err)
		os.Exit(1)
	}
}

// END OMIT
