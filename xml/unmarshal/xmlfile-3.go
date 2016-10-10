package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"encoding/json"
)

type Query struct{
	Dat Data `xml:"Data"`
}

func (s Query) PrintJsonStruct() {
	sB, _ := json.Marshal(s)
	fmt.Println(string(sB))
}

func (s Query) PrintXMLStruct() {
	sB, _ := xml.Marshal(s)
	fmt.Println(string(sB))
}


type Data struct {
	Ser Series `xml:"Series"`
	EpisodeList []Episode `xml:"Episode"`
}


func (s Data) PrintJsonStruct() {
	sB, _ := json.Marshal(s)
	fmt.Println(string(sB))
}

func (s Data) PrintXMLStruct() {
	sB, _ := xml.Marshal(s)
	fmt.Println(string(sB))
}


type Series struct {
	//We can name field name differently from the XML tag. Same rule applied for the struct's name.
	//The xml tag point to the xml tag in the file --> must be same as the tags appear in the XML file
	Id 	int  		`xml:"id"`
	Title 	string  	`xml:"SeriesName"`
	Actors  string 		`xml:"Actors"`
	ShowDay string		`xml:"Airs_DayOfWeek"`
	Keywords map[string] bool
}

type Episode struct {
	SeasonNumber 	int
	EpisodeNumber 	int
	EpisodeName 	string
	FirstAired 	string
}

func (s Series) String() string {
	return fmt.Sprintf("%d - %s", s.Id, s.Title)
}

func (e Episode) String() string {
	return fmt.Sprintf("S%02dE%02d - %s - %s", e.SeasonNumber, e.EpisodeNumber, e.EpisodeName, e.FirstAired)
}

func main() {
	dat, _ := ioutil.ReadFile("data/xmlfile-3.xml")

	var q Query
	//xml.Unmarshal(bValue, &q)
	xml.Unmarshal(dat, &q.Dat)
	_ = "breakpoint"
	fmt.Println("Data:", q.Dat)
	fmt.Println("Actors:", q.Dat.Ser.Actors)
	fmt.Println("ShowDay:", q.Dat.Ser.ShowDay)

	for _, episode := range q.Dat.EpisodeList {
		fmt.Printf("\t%s\n", episode)
	}

	fmt.Println("XML struct: ")
	q.PrintXMLStruct()
	fmt.Println("JSON struct: ")
	q.PrintJsonStruct()
}