package main

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search"
	"log"
	"reflect"
	"strconv"
	"utils/hashmap"
	"encoding/json"
)

func main() {
	BleveExample2()
}

func TestHashMap1() {
	//Test map[string]string
	for i := 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		log.Printf("Set %s completed \n", key)
		go hashmap.Set(key, "Hi - "+strconv.Itoa(i))
	}

	for i := 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		val, err := hashmap.Get(key)
		if err != nil {
			log.Printf("Get %s Err: %s\n", key, err.Error())
		} else {
			fmt.Println("Type:", reflect.TypeOf(val).Name())
			fmt.Println("Key/Value:", key, val)
		}

	}
}

func TestHashMap2() {
	type Person struct {
		Name string
		Age  int
	}
	Persons := hashmap.NewHashMap()

	//Test map[string]struct
	for i := 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		value := Person{
			Name: "Nghiep",
			Age:  i,
		}

		go func() {
			err := Persons.Put(key, value)
			if err != nil {
				fmt.Printf("Set failed, %s \n", err.Error())
			} else {
				log.Printf("Set %s completed \n", key)
			}

		}()
	}

	for i := 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		val, err := Persons.Get(key)
		if err != nil {
			log.Printf("Get %s Err: %s\n", key, err.Error())
		} else {
			fmt.Println("Type:", reflect.TypeOf(val).Name())
			fmt.Println("Key/Value:", key, val)
			fmt.Printf("Name: %s, Age: %d \n", val.(Person).Name, val.(Person).Age)

		}

	}

}

func TestHashMap3() {
	//Test Delete
	key := "key"
	hashmap.Set(key, "Hi Nghiep")
	log.Printf("Set %s completed \n", key)
	log.Println("Map Size:", hashmap.GetSize())

	val, err := hashmap.Get(key)
	if err != nil {
		log.Printf("Get %s Err: %s\n", key, err.Error())
	} else {
		fmt.Println("Type:", reflect.TypeOf(val).Name())
		fmt.Println("Key/Value:", key, val)
	}

	//Delete
	err = hashmap.Delete(key)
	if err != nil {
		log.Printf("Delete %s Err: %s\n", key, err.Error())
	} else {
		val, err = hashmap.Get(key)
		if err != nil {
			log.Printf("Err: %s\n", err.Error())
		} else {
			fmt.Println("Type:", reflect.TypeOf(val).Name())
			fmt.Println("Key/Value:", key, val)
		}

	}

}

func BleveExample1() {

	//create new index, if not exist create one
	hashIndex := hashmap.NewHashIndex("hashmap/indexstore")
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
		log.Println("Total:", searchResult.Total)
	}
	log.Println("----------------------------------")
	log.Println("searchResult:", searchResult)
}

func BleveExample2() {
	//var dataIndex bleve.Index
	indexPath := "hashmap/indexstore"
	//dataIndex := CreateSample(indexPath)

	dataIndex, err := bleve.Open(indexPath)
	if err != nil {
		mapping := bleve.NewIndexMapping()
		dataIndex, err = bleve.New(indexPath, mapping)
	}
	//Create sample data
	//CreateSampleData(dataIndex)

	//==========================================================================================================//
	// CASE I: search for the "easy". The NewQueryStringQuery will search for field Body, hit if search string
	// are contained in Body field
	searchStrQuery := bleve.NewQueryStringQuery("Body:indexing")

	//Declare a search request.
	searchRequest := bleve.NewSearchRequest(searchStrQuery)
	//--------------------------------------------------------------------------------------------------------//
	//Notice that the default NewSearchRequestOptions(q, 10, 0, false) with Size = 10, From index = 0, Explain = false
	//therefore by default the Hits array just return maximum 10 found items. But we can change it by set the :
	/*
	searchRequest.Size = 200		//Show max 200 matched items
	searchRequest.Explain = true		//print out the array of return values
	searchRequest.From = 10			//return from item 10
	*/
	//--------------------------------------------------------------------------------------------------------//
	//Execute search
	searchResult, _ := dataIndex.Search(searchRequest)

	//Display result
	log.Println("CASE I - SEARCH FIELD 'Body' CONTAINS 'indexing'")
	log.Println("---------------------------------------------------")
	log.Println("Total:", searchResult.Total)
	log.Println("---------------------------------------------------")
	log.Println("searchResult:", searchResult)



	//==========================================================================================================//
	// CASE II: search EXACT MATCHED the "easy". The NewQueryStringQuery will search for field Body, hit if search string
	// are contained in Body field
	searchPhaseQuery := bleve.NewMatchPhraseQuery("Body:easy")

	//Declare a search request.
	searchRequest = bleve.NewSearchRequest(searchPhaseQuery)

	//Execute search
	searchResult, _ = dataIndex.Search(searchRequest)

	//Display result
	log.Println()
	log.Println("CASE II(a) - SEARCH FIELD 'Body' CONTAINS EXACT 'easy'")
	log.Println("---------------------------------------------------")
	log.Println("Total:", searchResult.Total)
	log.Println("---------------------------------------------------")
	log.Println("searchResult:", searchResult)
	//Return 0 hit



	//Case II(b)
	//matchQuery := bleve.NewMatchPhraseQuery("msgId10")
	matchQuery := bleve.NewMatchQuery("msgId10")
	searchRequest = bleve.NewSearchRequest(matchQuery)
	searchResult, _ = dataIndex.Search(searchRequest)
	//Return 101, not compare match whole string if there is a '.-'. Only compare the first. How to?


	//Display result
	log.Println()
	log.Println("CASE II(b) - SEARCH FIELD 'Body' CONTAINS EXACT 'bleve indexing'")
	log.Println("---------------------------------------------------")
	log.Println("Total:", searchResult.Total)
	log.Println("---------------------------------------------------")
	log.Println("searchResult:", searchResult)
	//Return 0 hit

	//==========================================================================================================//
	searchPhaseQuery = bleve.NewMatchPhraseQuery("easy")
	searchRequest = bleve.NewSearchRequest(searchPhaseQuery)
	searchResult, _ = dataIndex.Search(searchRequest)
	//Display result
	log.Println("CASE III - SEARCH ALL FIELDS WHICH HAVE EXACT 'easy'")
	log.Println("---------------------------------------------------")
	log.Println("Total:", searchResult.Total)
	log.Println("---------------------------------------------------")
	log.Println("searchResult:", searchResult)
	//Return 101 hits
	//==========================================================================================================//
	//// Case III: Field Scoping, search for the "marty.schoch@gmail.com" in field "From".
	//searchValueInField := bleve.NewQueryStringQuery("From:marty.schoch@gmail.com")
	//
	////Declare a search request
	//searchRequest = bleve.NewSearchRequest(searchValueInField)
	//
	////Execute search
	//searchResult, _ = dataIndex.Search(searchRequest)
	//
	////Display result
	////Display result
	//if searchResult.Hits.Len() > 0 {
	//	log.Println("2A - SEARCH FIELDS 'From' CONTAINS 'marty.schoch@gmail.com'")
	//	log.Println("---------------------------------")
	//	PrintDocumentMatchCollection(searchResult.Hits)
	//
	//	log.Println("Total:", searchResult.Total)
	//	log.Println("ID:", searchResult.Hits[0].ID)
	//	log.Println("Index:", searchResult.Hits[0].Index)
	//	log.Println("Fields:", searchResult.Hits[0].Fields)
	//	log.Println("Fragments:", searchResult.Hits[0].Fragments)
	//	log.Println("Locations: ", searchResult.Hits[0].Locations)
	//	log.Println("Score:", searchResult.Hits[0].Score)
	//}
	//log.Println("----------------------------------")
	//log.Println("searchResult:", searchResult)
	//
	//// Case 2: Field Scoping, search for the "marty.schoch@gmail.com" in field "From".
	//searchMatchField := bleve.NewQueryStringQuery("+From:marty.schoch@gmail.com")
	//
	////Declare a search request
	//searchRequest = bleve.NewSearchRequest(searchMatchField)
	//
	////Execute search
	//searchResult, _ = dataIndex.Search(searchRequest)
	//
	////Display result
	////Display result
	//if searchResult.Hits.Len() > 0 {
	//	log.Println("2B - SEARCH FIELDS 'From' MATCHED 'marty.schoch@gmail.com'")
	//	log.Println("---------------------------------")
	//	log.Println("Total:", searchResult.Total)
	//	PrintDocumentMatchCollection(searchResult.Hits)
	//	log.Println("ID:", searchResult.Hits[0].ID)
	//	log.Println("Index:", searchResult.Hits[0].Index)
	//	log.Println("Fields:", searchResult.Hits[0].Fields)
	//	log.Println("Fragments:", searchResult.Hits[0].Fragments)
	//	log.Println("Locations: ", searchResult.Hits[0].Locations)
	//	log.Println("Score:", searchResult.Hits[0].Score)
	//}
	//log.Println("----------------------------------")
	//log.Println("searchResult:", searchResult)
	//
	//// Case 3: Required, Optional, and Exclusion.
	//// Example: +description:water -light beer will perform a Boolean Query that MUST satisfy
	//// the Match Query for the term water in the description field, MUST NOT satisfy the Match
	//// Query for the term light in the default field, and SHOULD satisfy the Match Query for
	//// the term beer in the default field.
	//searchBooleanValue := bleve.NewQueryStringQuery("+From:marty.schoch@gmail.com -easy bleve")
	//
	////Declare a search request
	//searchRequest = bleve.NewSearchRequest(searchBooleanValue)
	//
	////Execute search
	//searchResult, _ = dataIndex.Search(searchRequest)
	//
	////Display result
	//if searchResult.Hits.Len() > 0 {
	//	log.Println("3 - SEARCH FIELDS 'Value' CONTAINS '+From:marty.schoch@gmail.com -easy bleve'")
	//	log.Println("Search message that From field has 'marty.schoch@gmail.com', not have 'easy', have 'bleve'")
	//	log.Println("---------------------------------")
	//	log.Println("Total:", searchResult.Total)
	//	PrintDocumentMatchCollection(searchResult.Hits)
	//	log.Println("ID:", searchResult.Hits[0].ID)
	//	log.Println("Index:", searchResult.Hits[0].Index)
	//	log.Println("Fields:", searchResult.Hits[0].Fields)
	//	log.Println("Fragments:", searchResult.Hits[0].Fragments)
	//	log.Println("Locations: ", searchResult.Hits[0].Locations)
	//	log.Println("Score:", searchResult.Hits[0].Score)
	//}
	//log.Println("----------------------------------")
	//log.Println("searchResult:", searchResult)
	//
	////Case 4: Numeric Ranges
	////You can perform numeric ranges by using the >, >=, <, and <= operators, followed by a numeric value.
	////Example: abv:>10 will perform an Numeric Range Query on the abv field for values greater than ten.
	//searchRannge := bleve.NewQueryStringQuery("Value:>5 Value:<11 ")
	//
	////Declare a search request
	//searchRequest = bleve.NewSearchRequest(searchRannge)
	//
	////Execute search
	//searchResult, _ = dataIndex.Search(searchRequest)
	//
	////Display result
	//if searchResult.Hits.Len() > 0 {
	//	log.Println("4 - SEARCH FIELDS 'Value' CONTAINS 'Value:>5 Value:<11'")
	//	log.Println("---------------------------------")
	//	log.Println("Total:", searchResult.Total)
	//	PrintDocumentMatchCollection(searchResult.Hits)
	//	log.Println("ID:", searchResult.Hits[0].ID)
	//	log.Println("Index:", searchResult.Hits[0].Index)
	//	log.Println("Fields:", searchResult.Hits[0].Fields)
	//	log.Println("Fragments:", searchResult.Hits[0].Fragments)
	//	log.Println("Locations: ", searchResult.Hits[0].Locations)
	//	log.Println("Score:", searchResult.Hits[0].Score)
	//}
	//log.Println("----------------------------------")
	//log.Println("searchResult:", searchResult)
}

type Message struct {
	Id    string
	From  string
	Body  string
	Value int
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


type Key struct{
	Id 	string
	To 	string
	Body 	string
	Value   int
}

func CreateSampleData(indexVault bleve.Index)  {
	//bleve.Index is an interface, must not be pass using the pointer. Because you will not able to
	//access the interface methods via its pointer. Why? Recall that the interface is the collection of signature
	//These functions do not actually implements in the interface and therefore work like a pointer to the method of
	//the real object that this interface variable represents. For example:	var index *bleve.Index to declare
	// a variable of interface Index. Then we cannot use index.<method's name> to access the method
	for i := 0; i <= 100; i++ {

		key := Key{
			Id:    "msgId" + strconv.Itoa(i),
			To:  "marty" + strconv.Itoa(i) + ".schoch@gmail.com",
			Body:  "bleve indexing is easy",
			Value: i,
		}

		keyStr, _ := json.Marshal(key)

		message := Message{
			Id:    "msgId" + strconv.Itoa(i),
			From:  "marty" + strconv.Itoa(i) + ".schoch@gmail.com",
			Body:  "bleve indexing is easy",
			Value: i,
		}

		//Add new index, field to be used as Index, we can define multiple index at the same time for one object
		indexVault.Index(string(keyStr), message)
		//This keyStr is the return value from the search result.
		//Assume I save message object into key-value bolt (key= message.Id, value = message).
		//If I want to search the messages in bolt that satisfy my conditions, if match found,
		//the search result will contain the key that I will then use it to retrive the object
		//from BoltDB.
	}
}