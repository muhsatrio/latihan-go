package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", returnAllUsers).Methods("GET")
	router.HandleFunc("/user/add", insertUsers).Methods("POST")
	router.HandleFunc("/user/update/{id}", updateUsers).Methods("PUT")
	router.HandleFunc("/user/delete/{id}", deleteUser).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to PORT :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
