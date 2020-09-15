package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

// Post struct
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	XML     string   `xml:",innerxml"`
}

// Author struct
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		ID:      "1",
		Content: "Hello World",
		Author: Author{
			ID:   "2",
			Name: "Satrio",
		},
	}

	// by marshall way

	output, err := xml.Marshal(&post)
	if err != nil {
		log.Fatalln("Error marshaling to XML: ", err)
	}
	err = ioutil.WriteFile("post.xml", output, 0644)
	if err != nil {
		log.Fatalln("Error writing XML to file: ", err)
	}

	// by encode way

	xmlFile, err := os.Create("post.xml")
	if err != nil {
		log.Fatalln("Error creating file xml: ", err)
	}
	encoder := xml.NewEncoder(xmlFile)
	err = encoder.Encode(&post)
	if err != nil {
		log.Fatalln("Error encode file xml: ", err)
	}
}
