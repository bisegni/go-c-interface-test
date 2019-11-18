package main

/*
#cgo CFLAGS: -I . -stdlib=libc++
#cgo LDFLAGS: -L${SRCDIR}/build -ldbengine -L${SRCDIR}/build/boostinstall/lib -lboost_system -lstdc++

#include "src/dbengine.h"

#include <stdio.h>
#include <stdlib.h>

*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	//test call c method in comment
	queryStr := C.CString("select * form dual\n")
	defer C.free(unsafe.Pointer(queryStr))

	queryUUID := C.CString("")
	defer C.free(unsafe.Pointer(queryUUID))

	// C.myprint(queryStr)

	C.submitQuery(queryStr, queryUUID)

	fmt.Println("UUID is: " + C.GoString(queryUUID))

	//call library function form dbengine lib
	C.ACFunction()
}
