package main
import (

)
import (
	"os"
	"utils/golog"
	"utils/env"
)

var(
	terminal = os.Stdout
	errorStd = os.Stderr
)

func main() {
	testGolog()


}

func testGolog()  {
	golog.Enter()
	golog.SetStringPrefix(1, "LV1-")
	golog.SetStringPrefix(2, "LV2---")
	golog.SetStringPrefix(3, "LV3------")

	golog.SetFlags(4, golog.Ltime|golog.Lshortfile)
	golog.SetPrefix(1, "L1P ")
	golog.SetPrefix(2, "L2P ")
	golog.SetPrefix(3, "L3P ")

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		golog.Fatal(4, "Failed to open log file",  err.Error())
	}
	golog.SetOutput(4, file, terminal)   //4 out of range, --> default is all
	golog.ShowLogUptoLevel(2)

	golog.Println(1, "This is level 1")
	golog.Println(2, "This is level 2")
	golog.Println(3, "This is level 3")

	env.Hello()

	golog.Exit()
}


