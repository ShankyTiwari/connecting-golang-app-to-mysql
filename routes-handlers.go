package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func renderHome(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "views/index.html")
}

func getUsers(response http.ResponseWriter, request *http.Request) {
	var (
		user  User
		users []User
	)

	username := mux.Vars(request)["name"]

	rows, err := db.Query("SELECT * FROM users where name like '%" + username + "%'")

	if err != nil {
		fmt.Println(err)
		returnErrorResponse(response, request)
	}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Lname, &user.Country)
		users = append(users, user)
	}
	defer rows.Close()

	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		returnErrorResponse(response, request)
	}

	if jsonResponse == nil {
		returnErrorResponse(response, request)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request) {
	jsonResponse, err := json.Marshal("It's not you it's me.")
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	response.Write(jsonResponse)
}
