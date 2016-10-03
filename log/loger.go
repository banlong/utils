package logger

import (
	"log"
	"io/ioutil"
	"io"
)

type DebugLog struct {
	l1 *log.Logger
	l2 *log.Logger
	l3 *log.Logger
	l1_prefix string
	l2_prefix string
	l3_prefix string
}

func NewDebugLog(displaylevel int, output ...io.Writer) *DebugLog {
	if displaylevel > 3 || displaylevel < 0 {
		return &DebugLog{
			l1: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l2: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l3: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
		}
	}
	multi := io.MultiWriter(output...)
	retLogger := DebugLog{}
	switch displaylevel {
	case 0:
		retLogger =  DebugLog{
			l1: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l2: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l3: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
		}
	case 1:
		retLogger = DebugLog{
			l1: log.New(multi,"", log.Ldate|log.Ltime),
			l2: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
			l3: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
		}
	case 2:
		retLogger = DebugLog{
			l1: log.New(multi,"", log.Ldate|log.Ltime),
			l2: log.New(multi,"", log.Ldate|log.Ltime),
			l3: log.New(ioutil.Discard,"", log.Ldate|log.Ltime),
		}
	case 3:
		retLogger = DebugLog{
			l1: log.New(multi,"", log.Ldate|log.Ltime),
			l2: log.New(multi,"", log.Ldate|log.Ltime),
			l3: log.New(multi,"", log.Ldate|log.Ltime),
		}
	}

	return  &retLogger
}

func (dl *DebugLog) SetL1Prefix(pf string)  {
	dl.l1_prefix = pf
}

func (dl *DebugLog) SetL2Prefix(pf string)  {
	dl.l2_prefix = pf
}

func (dl *DebugLog) SetL3Prefix(pf string)  {
	dl.l3_prefix = pf
}

func (dl *DebugLog) Println(level int, input string) {
	switch level {
	case 1:
		dl.l1.Println(dl.l1_prefix + input)
	case 2:
		dl.l2.Println(dl.l2_prefix + input)
	case 3:
		dl.l3.Println(dl.l3_prefix + input)
	default:
		dl.l1.Println(dl.l1_prefix + input)
	}
}

func (dl *DebugLog) Print(level int, input string) {
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

func (dl *DebugLog) Printf(level int, input string) {
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

func (dl *DebugLog) Fatal(level int, input string) {
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

func (dl *DebugLog) Fatalf(level int, input string) {
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

func (dl *DebugLog) Fatalln(level int, input string) {
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

func (dl *DebugLog) Panic(level int, input string) {
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
func (dl *DebugLog) SetFlags(level int, flag int) {
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

func (dl *DebugLog) SetPrefix(level int, pf string) {
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
func (dl *DebugLog) SetOutput(level int, w ...io.Writer) {
	multi := io.MultiWriter(w...)
	switch level {
	case 1:
		dl.l1.SetOutput(multi)
	case 2:
		dl.l2.SetOutput(multi)
	case 3:
		dl.l3.SetOutput(multi)
	default:
		dl.l1.SetOutput(multi)
		dl.l2.SetOutput(multi)
		dl.l3.SetOutput(multi)
	}

}