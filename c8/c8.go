package main

/*
#include "windows.h"
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);

    printf("printint:ProcessId  %d\n", GetCurrentProcessId());
    printf("printint:ThreadId  %d\n", GetCurrentThreadId());
}
*/
import "C"


import "fmt"

func main() {
	v := 42
	C.printint(C.int(v))

	ProcessId := C.GetCurrentProcessId()
	fmt.Println("main ProcessId", ProcessId)

	ThreadId := C.GetCurrentThreadId()
	fmt.Println("main ThreadId", ThreadId)
}