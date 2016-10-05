package main
import (

)
import (
	"utils/logger"
	"os"
	"utils/alias"
)

var(
	Trace = logger.NewTracer(3, os.Stdout)
)

func main() {
	//testTrace()
	alias.TestFunction()
}

func testTrace()  {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Trace.Fatal(4, "Failed to open log file",  err.Error())
	}

	Trace.SetL1StringPrefix(".")
	Trace.SetL2StringPrefix("..")
	Trace.SetL3StringPrefix("...")
	Trace.SetOutput(4, file)
	Trace.Println(1, "This is level 1")
	Trace.Println(2, "This is level 2")
	Trace.Println(3, "This is level 3")
}




