//Example, bigcache object cannot be save. Data is on Heap, after app turn off the data will be erased
package maind

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

	//bigcache provide only get and set
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	cache.Set("my-unique-key", []byte("value"))
	entry, _ := cache.Get("my-unique-key")
	fmt.Println("String value:", string(entry))

	//--------Store struct using json--------
	dog := Dog{
		Name: "Nancy",
		Owner: "Richard",
	}
	encodedDog, err := json.Marshal(dog)
	if err != nil{
		log.Println(err.Error())
	}
	cache.Set("mydog", encodedDog)

	//Get struct
	var getDog = new(Dog)
	entry, _ = cache.Get("mydog")
	err = json.Unmarshal(entry, getDog)
	fmt.Println("Struct Value:",*getDog)

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
	cache.Set("mydog", encodedDog2.Bytes())

	//Get struct
	getDog = new(Dog)
	entry, _ = cache.Get("mydog")
	decoder := gob.NewDecoder(bytes.NewReader(entry))
	err = decoder.Decode(getDog)
	fmt.Println("Struct Value:",*getDog)


}



