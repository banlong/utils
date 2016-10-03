package main
import (

)
import (
	"utils/log"
	"os"
)

var(
	Trace = logger.NewDebugLog(3, os.Stdout)
)

func main() {
	Trace.SetL1Prefix(".")
	Trace.SetL2Prefix("..")
	Trace.SetL3Prefix("...")
	Trace.Println(1, "This is level 1")
	Trace.Println(2, "This is level 2")
	Trace.Println(3, "This is level 3")
}
