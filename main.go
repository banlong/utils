package main
import (

)
import (
	"os"
	"utils/golog"
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

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		golog.Fatal(4, "Failed to open log file",  err.Error())
	}
	golog.SetOutput(4, file, terminal)   //4 out of range, --> default is all

	golog.ShowLogUptoLevel(2)

	golog.Println(1, "This is level 1")
	golog.Println(2, "This is level 2")
	golog.Println(3, "This is level 3")

	golog.Exit()
}


