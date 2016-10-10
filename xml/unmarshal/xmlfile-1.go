package main

import (
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

type Query struct {
	Chan Channel `xml:"Channel"`
}

type Channel struct {
	Title string `xml:"Title"`
	Desc  string `xml:"Description"`
	Per   Person `xml:"Person"`
}

type Person struct {
	First string `xml:"First"`
	Last  string `xml:"Last"`
}
func main() {

	//open file
	dat, err := ioutil.ReadFile("data/xmlfile-1.xml")
	if err != nil {
		panic(err)
	}

	//Read XML byte content, convert it into struct
	var q Query
	xml.Unmarshal(dat, &q.Chan)

	fmt.Println(q)
}