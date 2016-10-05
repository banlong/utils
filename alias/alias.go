package alias

import (
	"fmt"
)

type MyInt int

type MyMap map[int]int

type MySlice []int

type MyConvertFunc func(int) string

//var i int = 2
//var i2 MyInt = 4
//i = i2			//cannot use i2 (type MyInt) as type int in assignment
//fmt.Println("i:", i)
//fmt.Println("i2:", i)


func TestMap() {
	mapdata := make(MyMap)
	mapdata[0] = 5
	mapdata[1] = 6
	mapdata[2] = 7
	mapdata[3] = 8
	mapdata[4] = 9
	printMap(mapdata)
}
func printMap(input MyMap){
	for key, val := range input{
		fmt.Println("[", key, ",", val, "]")
	}
}

func TestSlice() {
	fmt.Println("-- slice data 1 --")
	var slicedata = make(MySlice, 5)
	slicedata[0] = 5
	slicedata[1] = 6
	slicedata[2] = 7
	slicedata[3] = 8
	slicedata[4] = 9
	printSlice(slicedata)

	fmt.Println("-- slice data 2 --")
	var slicedata2 MySlice
	slicedata2 = append(slicedata2, 5)
	slicedata2 = append(slicedata2, 6)
	slicedata2 = append(slicedata2, 7)
	slicedata2 = append(slicedata2, 8)
	slicedata2 = append(slicedata2, 9)
	printSlice(slicedata2)
}

func printSlice(input MySlice){
	for index, val := range input{
		fmt.Println("[", index, ",", val, "]")
	}
}

//func TestFunction()  {
//	result := quote123(123)
//	fmt.Println(result)
//}
//
//func quote123(fn MyConvertFunc) string {
//	return fmt.Sprintf("%q", fn(123))
//}

