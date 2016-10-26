package main

import (
	"utils/hashmap"
	"log"
	"fmt"
	"strconv"
	"reflect"
)

func main()  {
	TestHashMap3()
}

func TestHashMap1()  {
	//Test map[string]string
	for i:= 0; i < 100; i++{
		key := "key" + strconv.Itoa(i)
		log.Printf("Set %s completed \n", key)
		go hashmap.Set(key, "Hi - " + strconv.Itoa(i))
	}

	for i:= 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		val, err := hashmap.Get(key)
		if err != nil {
			log.Printf("Get %s Err: %s\n", key, err.Error())
		}else{
			fmt.Println("Type:", reflect.TypeOf(val).Name())
			fmt.Println("Key/Value:", key, val)
		}

	}
}

func TestHashMap2()  {
	type Person struct{
		Name string
		Age  int
	}
	Persons := hashmap.NewHashMap()



	//Test map[string]struct
	for i:= 0; i < 100; i++{
		key := "key" + strconv.Itoa(i)
		value := Person{
			Name:"Nghiep",
			Age: i,
		}

		go func() {
			err := Persons.Put(key, value)
			if err != nil{
				fmt.Printf("Set failed, %s \n", err.Error())
			}else{
				log.Printf("Set %s completed \n", key)
			}

		}()
	}

	for i:= 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		val, err := Persons.Get(key)
		if err != nil {
			log.Printf("Get %s Err: %s\n", key, err.Error())
		}else{
			fmt.Println("Type:", reflect.TypeOf(val).Name())
			fmt.Println("Key/Value:", key, val)
			fmt.Printf("Name: %s, Age: %d \n", val.(Person).Name, val.(Person).Age)

		}

	}


}

func TestHashMap3()  {
	//Test Delete
	key := "key"
	hashmap.Set(key, "Hi Nghiep")
	log.Printf("Set %s completed \n", key)
	log.Println("Map Size:", hashmap.GetSize())

	val, err := hashmap.Get(key)
	if err != nil {
		log.Printf("Get %s Err: %s\n", key, err.Error())
	}else{
		fmt.Println("Type:", reflect.TypeOf(val).Name())
		fmt.Println("Key/Value:", key, val)
	}

	//Delete
	err = hashmap.Delete(key)
	if err != nil {
		log.Printf("Delete %s Err: %s\n", key, err.Error())
	}else{
		val, err = hashmap.Get(key)
		if err != nil {
			log.Printf("Err: %s\n", err.Error())
		}else{
			fmt.Println("Type:", reflect.TypeOf(val).Name())
			fmt.Println("Key/Value:", key, val)
		}

	}

}