package main

import (
	"encoding/xml"
	"fmt"
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
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		log.Fatalln("Error opening file post.xml")
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatalln("Error reading XML File")
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
