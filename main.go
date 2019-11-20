package main

/*
#cgo CFLAGS: -I . -stdlib=libc++
#cgo LDFLAGS: -L${SRCDIR}/build -ldbengine -L${SRCDIR}/build/boostinstall/lib -lboost_system -lstdc++

#include "src/dbengine.h"

#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
*/
// import "C"

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	//test call c method in comment
	queryStr := C.CString("select * form dual\n")
	defer C.free(unsafe.Pointer(queryStr))

	queryUUID := make([]byte, 40)

	// C.myprint(queryStr)

	C.submitQuery(queryStr, (*C.char)(unsafe.Pointer(&queryUUID[0])))
	fmt.Println("UUID is: " + string(queryUUID))

	var colCount C.int = 0;
	C.columnCount((*C.char)(unsafe.Pointer(&queryUUID[0])), (*C.int)(&colCount));

	fmt.Println("UUID is: " + string(queryUUID), " row count = " + strconv.Itoa(int(colCount)))

	//call library function form dbengine lib
	C.ACFunction()
}
