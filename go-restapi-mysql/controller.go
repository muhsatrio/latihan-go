package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/schema"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arrUser []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err := rows.Scan(&users.ID, &users.FirstName, &users.LastName)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	decoder := schema.NewDecoder()

	decoder.Decode(&users, r.Form)

	if err != nil {
		log.Fatal(err)
	}

	// arrUser = append(arrUser, users)
	_, err = db.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", users.FirstName, users.LastName)

	if err != nil {
		log.Fatal(err)
	}

	response.Status = 1
	response.Message = "Success Added Data"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var response Response
	var users Users
	vars := mux.Vars(r)

	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	decode := schema.NewDecoder()

	decode.Decode(&users, r.Form)

	_, err = db.Exec("UPDATE person SET first_name = ?, last_name = ? WHERE id = ?", users.FirstName, users.LastName, vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	response.Status = 1
	response.Message = "Success Update Data"

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var response Response

	vars := mux.Vars(r)

	_, err := db.Exec("DELETE FROM person WHERE id=?", vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	response.Status = 1
	response.Message = "Delete User Successfull"

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}
