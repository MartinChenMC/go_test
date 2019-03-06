package main

//#include "windows.h"
//void SayHello(char* s);
import "C"

import (
	"fmt"
)

func main() {
	C.SayHello(C.CString("Hello, World martin \n"))

	ProcessId := C.GetCurrentProcessId()
	fmt.Println("main ProcessId", ProcessId)

	ThreadId := C.GetCurrentThreadId()
	fmt.Println("main ThreadId", ThreadId)
}

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))

	ProcessId := C.GetCurrentProcessId()
	fmt.Println("SayHello ProcessId", ProcessId)

	ThreadId := C.GetCurrentThreadId()
	fmt.Println("SayHello ThreadId", ThreadId)
}