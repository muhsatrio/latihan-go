package main

import (
	"encoding/json"
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
	post := Post{
		ID:      1,
		Content: "Hello World",
		Author: Author{
			ID:   2,
			Name: "Testing 123",
		},
		Comments: []Comment{
			Comment{
				ID:      3,
				Content: "Halo",
				Author:  "Orang 1",
			},
			Comment{
				ID:      4,
				Content: "Halo juga",
				Author:  "Orang 2",
			},
		},
	}
	// with marshall
	output, err := json.Marshal(&post)
	if err != nil {
		log.Fatalln("Marshalling failed: ", err)
	}
	err = ioutil.WriteFile("post.json", output, 0644)
	if err != nil {
		log.Fatalln("Writing file failed: ", err)
	}

	// with decode
	jsonFile, err := os.Create("post.json")
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		log.Fatalln("Error encoding JSON file: ", err)
	}
}
