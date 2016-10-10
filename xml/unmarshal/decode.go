package main

//Convert XML content into struct
import (
	"encoding/xml"
	"fmt"
	"encoding/json"
)

//This struct has no name
//Have Channel obj, Channel has Title ans Description
var XML = []byte(`
			<?xml version="1.0" encoding="UTF-8" ?>
			<Channel>
				<Title>test</Title>
				<Description>this is a test</Description>
			</Channel>
		`)

type Query struct {
	Chan Channel `xml:"Channel"`
}

type Channel struct {
	Title string `xml:"Title"`
	Desc  string `xml:"Description"`
}

func (s Channel) String() string {
	return fmt.Sprintf("%s - %s", s.Title, s.Desc)
}

func (s Channel) PrintJsonStruct() {
	sB, _ := json.Marshal(s)
	fmt.Println(string(sB))
}

func (s Channel) PrintXMLStruct() {
	sB, _ := xml.Marshal(s)
	fmt.Println(string(sB))
}

func main() {
	//Read XML byte content, convert it into struct
	var q Query
	xml.Unmarshal(XML, &q.Chan)

	fmt.Println(q)
	q.Chan.PrintJsonStruct()
	q.Chan.PrintXMLStruct()


}



