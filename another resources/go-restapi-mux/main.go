package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// Article struct
type Article struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

// Articles var
var Articles []Article

var idNow int

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	for index, article := range Articles {
		if article.ID == idInt {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	var article Article

	decoder := schema.NewDecoder()

	decoder.Decode(&article, r.Form)

	if err != nil {
		log.Fatal(err)
	}

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	keyInt, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range Articles {
		if article.ID == keyInt {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/article", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{ID: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{ID: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
