package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Author struct
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Comment struct
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Post struct
type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		log.Fatalln("Error open json file: ", err)
	}

	defer jsonFile.Close()

	// with unmarshall

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln("Error reading JSON data: ", err)
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)

	// with decoder

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error decoding JSON: ", err)
		}
		fmt.Println(post)
	}
}
