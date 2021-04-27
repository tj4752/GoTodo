package main

import (
	"GoTodo/model"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

var db, _ = gorm.Open("mysql", "root:root@/GoTodo?charset=utf8&parseTime=True&loc=Local")

func main() {
	defer db.Close()

	db.Debug().DropTableIfExists(&model.TodoItemModel{})
	db.Debug().AutoMigrate(&model.TodoItemModel{})

	router := mux.NewRouter()
	router.HandleFunc("/todo-completed", GetCompletedItems).Methods("GET")
	router.HandleFunc("/todo-incomplete", GetIncompleteItems).Methods("GET")
	router.HandleFunc("/todo", CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", UpdateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", DeleteItem).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8000", handler)
}
