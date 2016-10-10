package main
import (

)
import (
	"os"
	"utils/golog"
	"time"
)

var(
	terminal = os.Stdout
	errorStd = os.Stderr
)

func main() {
	testGolog()


}

func logSetup()  {
	golog.SetStringPrefix(1, "->")	//This string will appear before the the printed string
	golog.SetStringPrefix(2, "-->>")
	golog.SetStringPrefix(3, "--->>>")
	golog.SetFlags(4, golog.Ltime)		//log header display log time only

	//golog.SetPrefix(1, "L1P ")
	//golog.SetPrefix(2, "L2P ")
	//golog.SetPrefix(3, "L3P ")


	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		golog.Fatal(4, "Failed to open log file",  err.Error())
	}
	golog.SetOutput(4, file, terminal)   //Set all levels output to both file & stdout
	golog.ShowLogUptoLevel(2)	     //Show all 1-2 levels
}

func testGolog()  {
	logSetup()
	defer golog.TimeElapse(time.Now())		//log the elapse time from this point till end of the method
	golog.Enter()
	golog.Println(1, "This is level 1")
	golog.Println(2, "This is level 2")
	golog.Println(3, "This is level 3")
	golog.Exit()
}


