package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"userapi.com/models"
)

var _users []*models.User

func main() {
	_users = []*models.User{
		{Id: 1, Firstname: "Teste"},
		{Id: 2, Firstname: "Teste2"},
		{Id: 3, Firstname: "Teste3"},
	}

	router := mux.NewRouter()
	router.HandleFunc("/user", GetAll).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&_users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParameter, error := strconv.Atoi(params["id"])

	if error != nil {
		panic(error)
	}

	for _, item := range _users {
		if item.Id == idParameter {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	fmt.Fprint(w, "[]")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	fmt.Println(&newUser)

	_users = append(_users, &newUser)

	json.NewEncoder(w).Encode(&_users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParameter, error := strconv.Atoi(params["id"])

	if error != nil {
		panic(error)
	}

	_users, error = remove(idParameter, _users)

	if err := json.NewEncoder(w).Encode(&_users); error != nil {
		panic(err)
	}
}

func remove(id int, usersParam []*models.User) (users []*models.User, err error) {
	for i, item := range usersParam {
		if item.Id == id {
			_users = append(_users[:i], _users[i+1:]...)
			return _users, nil
		}
	}
	return []*models.User{}, err
}
