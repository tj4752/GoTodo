package main

import (
	"GoTodo/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todo-completed", GetCompletedItems).Methods("GET")
	router.HandleFunc("/todo-incomplete", GetIncompleteItems).Methods("GET")
	return router
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestCreateItem(t *testing.T) {
	samp := &model.TodoItemModel{Description: "Testcase 1", Completed: false}
	jsonSamp, _ := json.Marshal(samp)
	request, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonSamp))
	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetCompletedItems(t *testing.T) {
	request, _ := http.NewRequest("GET", "/todo-completed", nil)
	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetIncompleteItems(t *testing.T) {
	request, _ := http.NewRequest("GET", "/todo-incomplete", nil)
	response := executeRequest(request)
	checkResponseCode(t, http.StatusOK, response.Code)
}
