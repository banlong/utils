package logger

import (

	"io/ioutil"
	"io"
	"log"
)

type Tracer struct {
	l1 *log.Logger
	l2 *log.Logger
	l3 *log.Logger
	l1_prefix string
	l2_prefix string
	l3_prefix string
}

func NewTracer(displaylevel int, output ...io.Writer) *Tracer {
	if displaylevel > 3 || displaylevel < 0 {
		return &Tracer{
			l1: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l2: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l3: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
		}
	}
	multi := io.MultiWriter(output...)
	retLogger := Tracer{}
	switch displaylevel {
	case 0:
		retLogger =  Tracer{
			l1: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l2: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l3: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
		}
	case 1:
		retLogger = Tracer{
			l1: log.New(multi,"", log.Ldate|log.Ltime),
			l2: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l3: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
		}
	case 2:
		retLogger = Tracer{
			l1: log.New(multi,"", log.Ldate|log.Ltime),
			l2: log.New(multi,"", log.Ldate|log.Ltime),
			l3: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
		}
	case 3:
		retLogger = Tracer{
			l1: log.New(multi,"", log.Ldate|log.Ltime),
			l2: log.New(multi,"", log.Ldate|log.Ltime),
			l3: log.New(multi,"", log.Ldate|log.Ltime),
		}
	}

	return  &retLogger
}


func (dl *Tracer) Println(level int, input ...string) {
	var prtStr = ""
	for _, str := range input{
		prtStr += str
	}

	switch level {
	case 1:
		dl.l1.Println(dl.l1_prefix + prtStr)
	case 2:
		dl.l2.Println(dl.l2_prefix + prtStr)
	case 3:
		dl.l3.Println(dl.l3_prefix + prtStr)
	default:
		dl.l1.Println(dl.l1_prefix + prtStr)
	}
}

func (dl *Tracer) Print(level int, data ...string) {
	var input = ""
	for _, str := range data{
		input += str
	}

	switch level {
	case 1:
		dl.l1.Print(dl.l1_prefix + input)
	case 2:
		dl.l2.Print(dl.l2_prefix + input)
	case 3:
		dl.l3.Print(dl.l3_prefix + input)
	default:
		dl.l1.Print(dl.l1_prefix + input)
	}
}

func (dl *Tracer) Printf(level int, data ...string) {
	var input = ""
	for _, str := range data{
		input += str
	}

	switch level {
	case 1:
		dl.l1.Printf(dl.l1_prefix + input)
	case 2:
		dl.l2.Printf(dl.l2_prefix + input)
	case 3:
		dl.l3.Printf(dl.l3_prefix + input)
	default:
		dl.l1.Printf(dl.l1_prefix + input)
	}
}

func (dl *Tracer) Fatal(level int, data ...string) {
	var input = ""
	for _, str := range data{
		input += str
	}

	switch level {
	case 1:
		dl.l1.Fatal(dl.l1_prefix + input)
	case 2:
		dl.l2.Fatal(dl.l2_prefix + input)
	case 3:
		dl.l3.Fatal(dl.l3_prefix + input)
	default:
		dl.l1.Fatal(dl.l1_prefix + input)
	}
}

func (dl *Tracer) Fatalf(level int, data ...string) {

	var input = ""
	for _, str := range data{
		input += str
	}

	switch level {
	case 1:
		dl.l1.Fatalf(dl.l1_prefix + input)
	case 2:
		dl.l2.Fatalf(dl.l2_prefix + input)
	case 3:
		dl.l3.Fatalf(dl.l3_prefix + input)
	default:
		dl.l1.Fatalf(dl.l1_prefix + input)
	}
}

func (dl *Tracer) Fatalln(level int, data ...string) {
	var input = ""
	for _, str := range data{
		input += str
	}

	switch level {
	case 1:
		dl.l1.Fatalln(dl.l1_prefix + input)
	case 2:
		dl.l2.Fatalln(dl.l2_prefix + input)
	case 3:
		dl.l3.Fatalln(dl.l3_prefix + input)
	default:
		dl.l1.Fatalln(dl.l1_prefix + input)
	}
}

func (dl *Tracer) Panic(level int, data ...string) {
	var input = ""
	for _, str := range data{
		input += str
	}

	switch level {
	case 1:
		dl.l1.Panic(dl.l1_prefix + input)
	case 2:
		dl.l2.Panic(dl.l2_prefix + input)
	case 3:
		dl.l3.Panic(dl.l3_prefix + input)
	default:
		dl.l1.Panic(dl.l1_prefix + input)
	}
}

//flag can be log.Ldate|log.Ltime|log.Llongfile | log.Lshortfile | log.Lmicroseconds
func (dl *Tracer) SetFlags(level int, flag int) {
	switch level {
	case 1:
		dl.l1.SetFlags(flag)
	case 2:
		dl.l2.SetFlags(flag)
	case 3:
		dl.l3.SetFlags(flag)
	default:
		dl.l1.SetFlags(flag)
		dl.l2.SetFlags(flag)
		dl.l3.SetFlags(flag)
	}
}

//This prefix is add before the flag --> left most of the log
func (dl *Tracer) SetPrefix(level int, pf string) {
	switch level {
	case 1:
		dl.l1.SetPrefix(pf)
	case 2:
		dl.l2.SetPrefix(pf)
	case 3:
		dl.l3.SetPrefix(pf)
	default:
		dl.l1.SetPrefix(pf)
		dl.l2.SetPrefix(pf)
		dl.l3.SetPrefix(pf)
	}
}

//level value other than 1-3 mean "all level"
func (dl *Tracer) SetOutput(level int, w ...io.Writer) {
	//get flag
	flags := dl.GetFlags(level)

	//get prefix
	prefix := dl.GetPrefix(level)

	//get string prefix
	strPrefix :=  dl.GetStringPrefix(level)

	multi := io.MultiWriter(w...)
	switch level {
	case 1:
		dl.l1 = log.New(multi, prefix, flags)
		dl.SetL1StringPrefix(strPrefix)
	case 2:
		dl.l2 = log.New(multi, prefix, flags)
		dl.SetL2StringPrefix(strPrefix)
	case 3:
		dl.l3 = log.New(multi, prefix, flags)
		dl.SetL3StringPrefix(strPrefix)
	default:
		flags1 := dl.GetFlags(1)
		prefix1 := dl.GetPrefix(1)
		strPrefix1 :=  dl.GetStringPrefix(1)
		dl.l1 = log.New(multi, prefix1, flags1)
		dl.SetL1StringPrefix(strPrefix1)

		flags2 := dl.GetFlags(2)
		prefix2 := dl.GetPrefix(2)
		strPrefix2 :=  dl.GetStringPrefix(2)
		dl.l2 = log.New(multi, prefix2, flags2)
		dl.SetL2StringPrefix(strPrefix2)

		flags3 := dl.GetFlags(3)
		prefix3 := dl.GetPrefix(3)
		strPrefix3 :=  dl.GetStringPrefix(3)
		dl.l3 = log.New(multi, prefix3, flags3)
		dl.SetL3StringPrefix(strPrefix3)
	}
}
func (dl *Tracer) SetStringPrefix(level int, pf string)  {
	switch level {
	case 1:
		dl.l1_prefix = pf
		break
	case 2:
		dl.l2_prefix = pf
		break
	case 3:
		dl.l3_prefix = pf
	default:
		dl.l1_prefix = pf
		dl.l2_prefix = pf
		dl.l3_prefix = pf
	}
}

//Append a string in front of the input string of println
func (dl *Tracer) SetL1StringPrefix(pf string)  {
	dl.l1_prefix = pf
}

func (dl *Tracer) SetL2StringPrefix(pf string)  {
	dl.l2_prefix = pf
}

func (dl *Tracer) SetL3StringPrefix(pf string)  {
	dl.l3_prefix = pf
}

func (dl *Tracer) GetStringPrefix(level int)  string{
	switch level {
	case 1:
		return dl.l1_prefix
	case 2:
		return dl.l2_prefix
	case 3:
		return dl.l3_prefix
	default:
		return ""
	}
}

func (dl *Tracer) GetPrefix(level int)  string{
	switch level {
	case 1:
		return dl.l1.Prefix()
	case 2:
		return dl.l2.Prefix()
	case 3:
		return dl.l3.Prefix()
	default:
		return ""
	}
}

func (dl *Tracer) GetFlags(level int)  int{
	switch level {
	case 1:
		return dl.l1.Flags()
	case 2:
		return dl.l2.Flags()
	case 3:
		return dl.l3.Flags()
	default:
		return dl.l1.Flags()
	}
}