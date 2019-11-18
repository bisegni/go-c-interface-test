package main

/*
#cgo CFLAGS: -I . -stdlib=libc++
#cgo LDFLAGS: -L${SRCDIR}/build -ldbengine -L${SRCDIR}/build/boostinstall/lib -lboost_system -lstdc++

#include "src/dbengine.h"

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

int testref(const char * uuid, int *count) {
	int err = columnCount(uuid, count);
	return *count;//
}
*/
import "C"

import (
	"fmt"
	"strconv"
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

	var colCount C.int = 4884;
	var x = C.testref(queryUUID, (*C.int)(unsafe.Pointer(&colCount)));
	// C.columnCount(queryUUID, (*C.int)(&colCount))

	fmt.Println("UUID is: " + C.GoString(queryUUID), " row count = " + strconv.Itoa(int(x)))

	//call library function form dbengine lib
	C.ACFunction()
}
