package main

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//Get the collection reference
	personCollection := session.DB("test").C("people")

	//Add 2 persons (document) into collection
	err = personCollection.Insert(
		&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"},
	)
	if err != nil {
		log.Fatal(err)
	}

	//Query data
	result := Person{}
	results := []Person{}
	err = personCollection.Find(bson.M{"name": "Ale"}).One(&result)
	err = personCollection.Find(bson.M{"name": "Ale"}).All(&results)
	if err != nil {
		log.Fatal(err)
	}


	dbnames, _ := session.DatabaseNames()
	fmt.Println("Databases:", dbnames)
	fmt.Println("One Result:", result)
	fmt.Println("All Results:", results)
}