package colreader

import (
	"log"
	"os"
	"reflect"
)

// ColReader is the class that load a file for a column and read all elemen one per time or buffered
type ColReader struct {
	fileName string
	colType  reflect.Kind
	file     *os.File
	err      error
}

// Open the file associated to the column
func (r *ColReader) Open() error {
	var err error
	r.file, err = os.Open(r.fileName)
	if err != nil {
		log.Fatalf("Error while opening %v file  with error %v", r.fileName, err)
	}
	return err
}
