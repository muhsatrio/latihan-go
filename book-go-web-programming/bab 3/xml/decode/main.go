package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

// Post struct
type Post struct {
	XMLName  xml.Name `xml:"post"`
	ID       string   `xml:"id,attr"`
	Content  string   `xml:"content"`
	Author   Author   `xml:"author"`
	Comments Comments `xml:"comments"`
	XML      string   `xml:",innerxml"`
}

// Comments struct
type Comments struct {
	Comment []Comment `xml:"comment"`
}

// Comment struct
type Comment struct {
	ID      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
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
	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error decoding XML into tokens:", err)
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(comment)
			}
		}
	}
}
