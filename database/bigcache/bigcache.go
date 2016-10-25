//Example, bigcache object cannot be save. Data is on Heap, after app turn off the data will be erased
package main

import (
	"github.com/allegro/bigcache"
	"time"
	"fmt"
	"encoding/json"
	"log"
	"encoding/gob"
	"bytes"

)

type Dog struct{
	Name  string
	Owner string
}

// Fast, concurrent, evicting in-memory cache written to keep big number of entries without impact on performance.
// BigCache keeps entries on heap but omits GC for them. To achieve that operations on bytes arrays take place,
// therefore entries (de)serialization in front of the cache will be needed in most use cases.
func main()  {
	BigCacheExample()
}



func BigCacheExample()  {
	//bigcache provide only get and set
	fmt.Println("STORE {STRING, STRING}")
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	cache.Set("my-unique-key", []byte("value"))
	entry, _ := cache.Get("my-unique-key")
	fmt.Println("Key:", "my-unique-key")
	fmt.Println("Value:", string(entry))
	fmt.Println();

	//--------Store struct using json--------
	fmt.Println("STORE {STRING, STRUCT}")
	dog := Dog{
		Name: "Nancy",
		Owner: "Richard",
	}
	encodedDog, err := json.Marshal(dog)
	if err != nil{
		log.Println(err.Error())
	}
	cache.Set("mydog", encodedDog)

	//Get struct using Json
	var getDog = new(Dog)
	key:= "mydog"
	entry, _ = cache.Get(key)
	err = json.Unmarshal(entry, getDog)
	fmt.Println("Key:",key)
	fmt.Println("Value:",*getDog)
	fmt.Println();

	//--------Store struct using GOB--------
	doggy := Dog{
		Name: "Milo",
		Owner: "Bobby",
	}

	encodedDog2 := new(bytes.Buffer)
	encoder := gob.NewEncoder(encodedDog2)
	err = encoder.Encode(doggy)
	if err != nil{
		log.Println(err.Error())
	}
	key = "mydog"
	cache.Set(key, encodedDog2.Bytes())

	//Get struct
	getDog = new(Dog)
	entry, _ = cache.Get(key)
	decoder := gob.NewDecoder(bytes.NewReader(entry))
	err = decoder.Decode(getDog)
	fmt.Println("Key:",key)
	fmt.Println("Value:",*getDog)
	fmt.Println();


	//--------Store struct using GOB--------
	fmt.Println("STORE {STRUCT, STRUCT}")
	sample3 := Dog{
		Name: "Kitty",
		Owner: "Bobby2",
	}

	encodeVal3 := new(bytes.Buffer)
	encoder3 := gob.NewEncoder(encodeVal3)
	err = encoder3.Encode(sample3)
	if err != nil{
		log.Println(err.Error())
	}

	inputKey := Key{
		Id: "mydog",
		Dept: "it",
	}
	keyStr, err := json.Marshal(inputKey)
	cache.Set(string(keyStr), encodeVal3.Bytes())	//Bigcache only accept the string as key

	//Get struct
	searchKey := Key{
		Id: "mydog",
		Dept: "it",
	}
	searchKeyBytes, _ := json.Marshal(searchKey)
	searchKeyStr := string(searchKeyBytes)
	returnVal3 := new(Dog)
	returnEncodedVal3, _ := cache.Get(searchKeyStr)
	decoder3 := gob.NewDecoder(bytes.NewReader(returnEncodedVal3))
	err = decoder3.Decode(returnVal3)
	fmt.Println("Key:",searchKeyStr)
	fmt.Println("Value:",*returnVal3)
	fmt.Println();
}

type Key struct{
	Id string
	Dept string
}