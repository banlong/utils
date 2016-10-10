package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Entry struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Updated   string `xml:"updated"`
	Published string `xml:"published"`
	Category  string `xml:"category"`
	Summary   string `xml:"summary"`
}
type Feed struct {
	XMLName  xml.Name `xml:"feed"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	Id       string   `xml:"id"`
	Updated  string   `xml:"updated"`
	Logo     string   `xml:"logo"`
	Icon     string   `xml:"icon"`
	Rights   string   `xml:"rights"`
	Entries  []Entry  `xml:"entry"`
}

func (s Feed) PrintJsonStruct() {
	sB, _ := json.Marshal(s)
	fmt.Println(string(sB))
}

func (s Feed) PrintXMLStruct() {
	sB, _ := xml.Marshal(s)
	fmt.Println(string(sB))
}
func main() {
	//open file
	dat, err := ioutil.ReadFile("data/xmlfile-2.xml")
	if err != nil {
		panic(err)
	}

	//decode file into Feed
	var f Feed
	err2 := xml.Unmarshal(dat, &f)
	if err2 != nil {
		panic(err2)
	}

	f.PrintXMLStruct()

	f.PrintJsonStruct()
}
