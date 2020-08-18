package models

import (
	"github.com/francobottoni/go-rest-api/src/database"
	_"log"
)

type Todo struct{
	ID int `json:"id"`
	Description string `json:"description"`
}

func Insert(description string) (Todo,bool){
	db := database.GetConnection()

	var todo_id int
	sqlStatement := "INSERT INTO todos(description) VALUES($1) RETURNING id"
	db.QueryRow(sqlStatement, description).Scan(&todo_id)

	if todo_id == 0{
		return Todo{}, false
	}

	return Todo{todo_id, ""}, true
}

func Get(id string) (Todo,bool){
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM todos WHERE id = $1",id)

	var ID int
	var description string 
	err := row.Scan(&ID, &description)
	if err != nil {
		return Todo{}, false
	}

	return Todo{ID,description}, true
}