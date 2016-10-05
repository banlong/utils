package main
import (

)
import (
	"utils/logger"
	"os"
)

var(
	Trace = logger.NewTracer(3, os.Stdout)
	terminal = os.Stdout
	errorStd = os.Stderr
)

func main() {
	testTrace()
	//alias.TestFunction()
}

func testTrace()  {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Trace.Fatal(4, "Failed to open log file",  err.Error())
	}

	Trace.SetStringPrefix(1, ".")
	Trace.SetStringPrefix(2, "..")
	Trace.SetStringPrefix(3, "...")
	Trace.SetOutput(4, file, terminal)
	Trace.Println(1, "This is level 1")
	Trace.Println(2, "This is level 2")
	Trace.Println(3, "This is level 3")
}




