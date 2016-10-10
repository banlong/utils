package main

//Convert XML content into struct

import (
	"encoding/xml"
	"fmt"
	"log"

	"encoding/json"
)

type Dictionary struct {
	XMLName   xml.Name   `xml:"dictionary"`
	Grammemes []Grammeme `xml:"grammemes>grammeme"`
}

type Grammeme struct {
	Name   string `xml:",chardata"`
	Parent string `xml:"parent,attr"`
}

func (s Dictionary) PrintXMLStruct() {
	sB, _ := xml.Marshal(s)
	fmt.Println(string(sB))
}

func (s Dictionary) PrintJsonStruct() {
	sB, _ := json.Marshal(s)
	fmt.Println(string(sB))
}

//XML has name : dictionary
//There should be a struct Dictionary. Within it, we have array of grammeme and name it grammmes
//THe Grammeme struct can include 2 attibutes: Name, Parent
var XML  =      []byte(`<dictionary version="0.8" revision="403605">
			    <grammemes>
				<grammeme parent="">POST</grammeme>
				<grammeme parent="POST">NOUN</grammeme>
			    </grammemes>
			</dictionary>
			`)
func main() {
	var dict Dictionary
	if err := xml.Unmarshal(XML, &dict); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v\n", dict)
	dict.PrintXMLStruct()
}

//Result
//main.Dictionary{XMLName:xml.Name{Space:"", Local:"dictionary"}, Grammemes:[]main.Grammeme{main.Grammeme{Name:"POST", Parent:""}, main.Grammeme{Name:"NOUN", Parent:"POST"}}}
//{"XMLName":{"Space":"","Local":"dictionary"},"Grammemes":[{"Name":"POST","Parent":""},{"Name":"NOUN","Parent":"POST"}]}
