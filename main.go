package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lfoo

#include "foo.h"

#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import "unsafe"

func Example() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

func main() {
	Example();

	C.ACFunction();
}