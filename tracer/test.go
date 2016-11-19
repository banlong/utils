package main

import "github.com/sabhiram/go-tracey"

// Setup global enter exit trace functions (default options)
var Exit, Enter = tracey.New(nil)


func foo(i int) {
	// $FN will get replaced with the function's name
	defer Exit(Enter("$FN(%d)", i))
	if i != 0 {
		foo(i - 1)
	}
}

func main() {
	defer Exit(Enter())
	foo(2)
}