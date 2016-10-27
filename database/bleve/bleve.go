package main

import (
	"log"
	"github.com/blevesearch/bleve"
	"utils/hashmap"
	"fmt"

	"strconv"
	"github.com/blevesearch/bleve/search"
)

//Bleve create index based on the defined search values & return value. This is very useful to use the indexing with
//NoSQL(Key/Value) because we can search value that more flexible(one or many fields) in compare with bolt/map
//Step Are:
// Step 1: Create Index when save data to Bolt
// Step 2: Search the desired item in the index map, will return the key that later can be used to retrieve the value from Bolt
// Step 3: Using return key to get the data from Bolt or any Key/Value database

type Message  struct{
	Id 	string
	From 	string
	Body 	string
	Value   int
}

func main() {

	BleveExample1()
}

func BleveExample1() {

	//create new index, if not exist create one
	hashIndex := NewHashIndex("hashmap/indexstore")
	for i := 0; i < 11; i++ {
		message := Message{
			Id:    "msgId-" + strconv.Itoa(i),
			From:  "marty-" + strconv.Itoa(i) + ".schoch@gmail.com",
			Body:  "bleve indexing is easy",
			Value: i,
		}

		//add an index, using the message Id
		hashIndex.AddIndex(message.Id, message)
	}

	//find items have value > 5
	//searchResult := hashIndex.ExecQuery("Value:>5")
	searchResult := hashIndex.ExecQuery("easy")

	if searchResult.Hits.Len() > 0 {
		log.Println("SEARCH ALL Value CONTAINS '>5'")
		log.Println("---------------------------------")
		PrintDocumentMatchCollection(searchResult.Hits)
		log.Println("Total:", searchResult.Total)
	}
	log.Println("----------------------------------")
	log.Println("searchResult:", searchResult)
}

func BleveExample2() {
	var dataIndex bleve.Index
	indexPath := "hashmap/indexstore"

	dataIndex, err := bleve.Open(indexPath)
	if err != nil {
		mapping := bleve.NewIndexMapping()
		dataIndex, err = bleve.New(indexPath, mapping)
	}

	//Create sample data
	for i := 0; i < 11; i++ {
		message := Message{
			Id:    "msgId-" + strconv.Itoa(i),
			From:  "marty-" + strconv.Itoa(i) + ".schoch@gmail.com",
			Body:  "bleve indexing is easy",
			Value: i,
		}

		//Add new index, field to be used as Index, we can define multiple index at the same time for one object
		dataIndex.Index(message.Id, message)
		//This message.Id is the return value from the search result.
		//Assume I save message object into key-value bolt (key= message.Id, value = message).
		//If I want to search the messages in bolt that satisfy my conditions, if match found,
		//the search result will contain the key that I will then use it to retrive the object
		//from BoltDB.

	}

	// Case 1: search for the "easy". Plain terms without any other syntax are
	// interpreted as a match query for the term in the default field.
	// The default field is "_all" unless overridden in the index mapping.
	searchPlanValue := bleve.NewQueryStringQuery("easy")

	//Declare a search request
	searchRequest := bleve.NewSearchRequest(searchPlanValue)

	//Execute search
	searchResult, _ := dataIndex.Search(searchRequest)

	//Display result
	if searchResult.Hits.Len() > 0 {
		log.Println("1 - SEARCH ALL FIELDS CONTAINS 'easy'")
		log.Println("---------------------------------")
		PrintDocumentMatchCollection(searchResult.Hits)
		log.Println("Total:", searchResult.Total)
		log.Println("ID:", searchResult.Hits[0].ID)
		log.Println("Index:", searchResult.Hits[0].Index)
		log.Println("Fields:", searchResult.Hits[0].Fields)
		log.Println("Fragments:", searchResult.Hits[0].Fragments)
		log.Println("Locations: ", searchResult.Hits[0].Locations)
		log.Println("Score:", searchResult.Hits[0].Score)
	}
	log.Println("----------------------------------")
	log.Println("searchResult:", searchResult)

	// Case 2: Field Scoping, search for the "marty.schoch@gmail.com" in field "From".
	searchValueInField := bleve.NewQueryStringQuery("From:marty.schoch@gmail.com")

	//Declare a search request
	searchRequest = bleve.NewSearchRequest(searchValueInField)

	//Execute search
	searchResult, _ = dataIndex.Search(searchRequest)

	//Display result
	//Display result
	if searchResult.Hits.Len() > 0 {
		log.Println("2A - SEARCH FIELDS 'From' CONTAINS 'marty.schoch@gmail.com'")
		log.Println("---------------------------------")
		PrintDocumentMatchCollection(searchResult.Hits)

		log.Println("Total:", searchResult.Total)
		log.Println("ID:", searchResult.Hits[0].ID)
		log.Println("Index:", searchResult.Hits[0].Index)
		log.Println("Fields:", searchResult.Hits[0].Fields)
		log.Println("Fragments:", searchResult.Hits[0].Fragments)
		log.Println("Locations: ", searchResult.Hits[0].Locations)
		log.Println("Score:", searchResult.Hits[0].Score)
	}
	log.Println("----------------------------------")
	log.Println("searchResult:", searchResult)

	// Case 2: Field Scoping, search for the "marty.schoch@gmail.com" in field "From".
	searchMatchField := bleve.NewQueryStringQuery("+From:marty.schoch@gmail.com")

	//Declare a search request
	searchRequest = bleve.NewSearchRequest(searchMatchField)

	//Execute search
	searchResult, _ = dataIndex.Search(searchRequest)

	//Display result
	//Display result
	if searchResult.Hits.Len() > 0 {
		log.Println("2B - SEARCH FIELDS 'From' MATCHED 'marty.schoch@gmail.com'")
		log.Println("---------------------------------")
		log.Println("Total:", searchResult.Total)
		PrintDocumentMatchCollection(searchResult.Hits)
		log.Println("ID:", searchResult.Hits[0].ID)
		log.Println("Index:", searchResult.Hits[0].Index)
		log.Println("Fields:", searchResult.Hits[0].Fields)
		log.Println("Fragments:", searchResult.Hits[0].Fragments)
		log.Println("Locations: ", searchResult.Hits[0].Locations)
		log.Println("Score:", searchResult.Hits[0].Score)
	}
	log.Println("----------------------------------")
	log.Println("searchResult:", searchResult)

	// Case 3: Required, Optional, and Exclusion.
	// Example: +description:water -light beer will perform a Boolean Query that MUST satisfy
	// the Match Query for the term water in the description field, MUST NOT satisfy the Match
	// Query for the term light in the default field, and SHOULD satisfy the Match Query for
	// the term beer in the default field.
	searchBooleanValue := bleve.NewQueryStringQuery("+From:marty.schoch@gmail.com -easy bleve")

	//Declare a search request
	searchRequest = bleve.NewSearchRequest(searchBooleanValue)

	//Execute search
	searchResult, _ = dataIndex.Search(searchRequest)

	//Display result
	if searchResult.Hits.Len() > 0 {
		log.Println("3 - SEARCH FIELDS 'Value' CONTAINS '+From:marty.schoch@gmail.com -easy bleve'")
		log.Println("Search message that From field has 'marty.schoch@gmail.com', not have 'easy', have 'bleve'")
		log.Println("---------------------------------")
		log.Println("Total:", searchResult.Total)
		PrintDocumentMatchCollection(searchResult.Hits)
		log.Println("ID:", searchResult.Hits[0].ID)
		log.Println("Index:", searchResult.Hits[0].Index)
		log.Println("Fields:", searchResult.Hits[0].Fields)
		log.Println("Fragments:", searchResult.Hits[0].Fragments)
		log.Println("Locations: ", searchResult.Hits[0].Locations)
		log.Println("Score:", searchResult.Hits[0].Score)
	}
	log.Println("----------------------------------")
	log.Println("searchResult:", searchResult)

	//Case 4: Numeric Ranges
	//You can perform numeric ranges by using the >, >=, <, and <= operators, followed by a numeric value.
	//Example: abv:>10 will perform an Numeric Range Query on the abv field for values greater than ten.
	searchRannge := bleve.NewQueryStringQuery("Value:>5 Value:<11 ")

	//Declare a search request
	searchRequest = bleve.NewSearchRequest(searchRannge)

	//Execute search
	searchResult, _ = dataIndex.Search(searchRequest)

	//Display result
	if searchResult.Hits.Len() > 0 {
		log.Println("4 - SEARCH FIELDS 'Value' CONTAINS 'Value:>5 Value:<11'")
		log.Println("---------------------------------")
		log.Println("Total:", searchResult.Total)
		PrintDocumentMatchCollection(searchResult.Hits)
		log.Println("ID:", searchResult.Hits[0].ID)
		log.Println("Index:", searchResult.Hits[0].Index)
		log.Println("Fields:", searchResult.Hits[0].Fields)
		log.Println("Fragments:", searchResult.Hits[0].Fragments)
		log.Println("Locations: ", searchResult.Hits[0].Locations)
		log.Println("Score:", searchResult.Hits[0].Score)
	}
	log.Println("----------------------------------")
	log.Println("searchResult:", searchResult)
}


func PrintDocumentMatchCollection(data search.DocumentMatchCollection) {
	fmt.Print("Len:", data.Len(), " - ")
	fmt.Print("[")
	for i := 0; i < data.Len(); i++ {
		fmt.Print(data[i].ID, " ")
	}
	fmt.Print("]")
	fmt.Println()
}