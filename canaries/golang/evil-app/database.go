package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func seedUserData(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var users []User

	json.Unmarshal(byteValue, &users)

	db, err := sql.Open("sqlite3", "data/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dropTableQuery := "DROP TABLE IF EXISTS user;"
	_, err = db.Exec(dropTableQuery)
	if err != nil {
		panic(err)
	}

	createTableQuery := "CREATE TABLE IF NOT EXISTS user(id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, company TEXT, title TEXT, email TEXT, phone TEXT, dob TEXT, ssn TEXT, salary NUMERIC, admin BOOLEAN);"
	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(users); i++ {
		user := users[i]
		insertUserQuery := "INSERT INTO user VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
		_, err = db.Exec(insertUserQuery, i+1, user.FirstName, user.LastName, user.Company, user.Title, user.Email, user.Phone, user.DOB, user.SSN, user.Salary, user.Admin)
		if err != nil {
			panic(err)
		}
	}
}

func getUsers(search string) []User {
	db, err := sql.Open("sqlite3", "data/users.db")
	if err != nil {
		panic(err)
	}
	selectUsersQuery := "SELECT * FROM user WHERE (first_name LIKE ? OR last_name LIKE ?) AND admin == 'false';"
	searchPattern := "%" + search + "%"
	rows, err := db.Query(selectUsersQuery, searchPattern, searchPattern)
	if err != nil {
		panic(err)
	}
	var results []User

	for rows.Next() {
		var user User
		_ = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Company, &user.Title, &user.Email, &user.Phone, &user.DOB, &user.SSN, &user.Salary, &user.Admin)
		results = append(results, user)
	}
	rows.Close()
	return results
}
