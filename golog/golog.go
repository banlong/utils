package golog

import (
	"log"
	"io"
	"io/ioutil"
	"runtime"
	"regexp"
	"fmt"
	"os"
)

var(
	standardLogger = New(defaulInitLevel, defaultPrefix, defaultFlags, os.Stdout)
	invisibleWriter = ioutil.Discard
	terminal = os.Stdout
	errorStd = os.Stderr
)


// These flags define which text to prefix to each log entry generated by the Logger.
const (
	// Bits or'ed together to control what's printed.
	// There is no control over the order they appear (the order listed
	// here) or the format they present (as described in the comments).
	// The prefix is followed by a colon only when Llongfile or Lshortfile
	// is specified.
	// For example, flags Ldate | Ltime (or LstdFlags) produce,
	//	2009/01/23 01:23:23 message
	// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	defaultFlags = Ldate | Ltime // initial values for the standard logger
	defaultLevel  	= 1

	defaulInitLevel = 3
	defaultPrefix   = ""


)

type Golog struct {
	loggers map[int] *log.Logger
	headers map[int]interface{}			//add before the string, vs prefix add before flags
}

func New(level int, prefix string, flag int, w ...io.Writer) *Golog {
	multi := io.MultiWriter(w...)
	var ret = Golog{
		loggers: make(map[int] *log.Logger),
		headers: make(map[int]interface{}),
	}

	for index := 1; index <= level; index++{
		ret.loggers[index] = log.New(multi, prefix, flag)
	}
	return &ret
}

func (dl *Golog) Println(level int, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}
	prefix := []interface{}{dl.headers[level]}
	v = append(prefix, v...)
	selectLogger.Println(v...)

}

func (dl *Golog) Print(level int, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}

	if(dl.headers[level] != nil) {
		prefix := []interface{}{dl.headers[level]}
		v = append(prefix, v...)
	}
	selectLogger.Print(v...)
}

func (dl *Golog) Printf(level int, format string, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}
	if(dl.headers[level] != nil) {
		prefix := []interface{}{dl.headers[level]}
		v = append(prefix, v...)
	}
	selectLogger.Printf(format, v...)
}

func (dl *Golog) Fatal(level int, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}
	if(dl.headers[level] != nil) {
		prefix := []interface{}{dl.headers[level]}
		v = append(prefix, v...)
	}
	selectLogger.Fatal(v...)
}

func (dl *Golog) Fatalf(level int,format string, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}

	if(dl.headers[level] != nil){
		prefix := []interface{}{dl.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Fatalf(format, v...)
}

func (dl *Golog) Fatalln(level int, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}

	if(dl.headers[level] != nil){
		prefix := []interface{}{dl.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Fatalln(v...)
}

func (dl *Golog) Panic(level int, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}

	if(dl.headers[level] != nil){
		prefix := []interface{}{dl.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Panic(v...)
}

func (dl *Golog) Panicf(level int, format string, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}

	if(dl.headers[level] != nil){
		prefix := []interface{}{dl.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Panicf(format, v)
}

func (dl *Golog) Panicln(level int, v ...interface{}) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		selectLogger = dl.loggers[defaultLevel]
	}

	if(dl.headers[level] != nil){
		prefix := []interface{}{dl.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Panicln(v...)
}

//flag can be log.Ldate|log.Ltime|log.Llongfile | log.Lshortfile | log.Lmicroseconds
func (dl *Golog) SetFlags(level int, flag int) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		for _, logger := range dl.loggers{
			logger.SetFlags(flag)
		}

	}else{
		selectLogger.SetFlags(flag)
	}
}

//This prefix is add before the flag --> left most of the log
func (dl *Golog) SetPrefix(level int, prefix string) {
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		for _, logger := range dl.loggers{
			logger.SetPrefix(prefix)
		}
	}else{
		selectLogger.SetPrefix(prefix)
	}

}

//level value other than 1-3 mean "all level"
func (dl *Golog) SetOutput(level int, w ...io.Writer) {
	multi := io.MultiWriter(w...)
	selectLogger := dl.loggers[level]
	if selectLogger == nil{
		for _, logger := range dl.loggers{
			logger.SetOutput(multi)
		}
	}else{
		selectLogger.SetOutput(multi)
	}


}

func (dl *Golog) SetStringPrefix(level int, prefix string)  {
	dl.headers[level] = prefix
}

func (dl *Golog) GetLevelCount() int  {
	return len(dl.loggers)
}

func (dl *Golog) HideLog(levels ...int) {
	for _, lv := range levels {
		if curLogger := dl.loggers[lv]; curLogger != nil{
			curLogger.SetOutput(invisibleWriter)
		}
	}
}

//input 3,  have 5 level, will show 1, 2, 3
func (dl *Golog) ShowLogUptoLevel(level int) {
	totalLevel := len(dl.loggers)
	for i:= level + 1; i <= totalLevel; i++ {
		if logger := dl.loggers[i]; logger != nil{
			logger.SetOutput(invisibleWriter)
		}
	}
}

// Trace Functions
func (dl *Golog) Enter() {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	fmt.Printf("Entering %s\n", fnName)
}

func (dl *Golog) Exit() {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	fmt.Printf("Exiting  %s\n", fnName)
}

func Println(level int, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if standardLogger.headers[level] != nil{
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}
	selectLogger.Println(v...)
}

func Print(level int, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if standardLogger.headers[level] != nil{
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}
	selectLogger.Print(v...)
}

func Printf(level int, format string, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if(standardLogger.headers[level] != nil){
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Printf(format, v...)
}

func Fatal(level int, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if(standardLogger.headers[level] != nil){
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Fatal(v...)
}

func Fatalf(level int,format string, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if(standardLogger.headers[level] != nil){
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Fatalf(format, v...)
}

func Fatalln(level int, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if(standardLogger.headers[level] != nil){
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Fatalln(v...)
}

func Panic(level int, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if(standardLogger.headers[level] != nil){
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Panic(v...)
}

func Panicf(level int, format string, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if(standardLogger.headers[level] != nil){
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Panicf(format, v)
}

func Panicln(level int, v ...interface{}) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		selectLogger = standardLogger.loggers[defaultLevel]
	}

	if(standardLogger.headers[level] != nil){
		prefix := []interface{}{standardLogger.headers[level]}
		v = append(prefix, v...)
	}

	selectLogger.Panicln(v...)
}

//flag can be log.Ldate|log.Ltime|log.Llongfile | log.Lshortfile | log.Lmicroseconds
func SetFlags(level int, flag int) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		for _, logger := range standardLogger.loggers{
			logger.SetFlags(flag)
		}

	}else{
		selectLogger.SetFlags(flag)
	}
}

//This prefix is add before the flag --> left most of the log
func SetPrefix(level int, prefix string) {
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		for _, logger := range standardLogger.loggers{
			logger.SetPrefix(prefix)
		}
	}else{
		selectLogger.SetPrefix(prefix)
	}


}

//level value other than 1-3 mean "all level"
func SetOutput(level int, w ...io.Writer) {
	multi := io.MultiWriter(w...)
	selectLogger := standardLogger.loggers[level]
	if selectLogger == nil{
		for _, logger := range standardLogger.loggers{
			logger.SetOutput(multi)
		}
	}else{
		selectLogger.SetOutput(multi)
	}


}

func SetStringPrefix(level int, prefix string)  {
	standardLogger.headers[level] = prefix
}

func GetLevelCount() int  {
	return len(standardLogger.loggers)
}

func HideLog(levels ...int) {
	for _, lv := range levels {
		if curLogger := standardLogger.loggers[lv]; curLogger != nil{
			curLogger.SetOutput(invisibleWriter)
		}
	}
}

//input 3,  have 5 level, will show 1, 2, 3
func ShowLogUptoLevel(level int) {
	totalLevel := len(standardLogger.loggers)
	for i:= level + 1; i <= totalLevel; i++ {
		if logger := standardLogger.loggers[i]; logger != nil{
			logger.SetOutput(invisibleWriter)
		}
	}
}

// Trace Functions
func Enter() {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	fmt.Printf("Entering %s\n", fnName)
}

func Exit() {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	fmt.Printf("Exiting  %s\n", fnName)
}