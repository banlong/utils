package main
import (

)
import (
	"utils/log"
	"os"
)

var(
	Trace = logger.NewLogger(3, os.Stdout)
)

func main() {
	Trace.SetL1StringPrefix(".")
	Trace.SetL2StringPrefix("..")
	Trace.SetL3StringPrefix("...")
	Trace.Println(1, "This is level 1")
	Trace.Println(2, "This is level 2")
	Trace.Println(3, "This is level 3")
}
