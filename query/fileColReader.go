package query

import (
	"log"
	"os"
	"reflect"
	"encoding/binary"
)

// FileColReader is the class that load a file for a column and
// read all elemen one per time or buffered
type FileColReader struct {
	fileName string
	colType  reflect.Kind
	file     *os.File
	err      error
}

// NewFileColReader allocate new instance
func NewFileColReader(_fileName string) *FileColReader {
	return &FileColReader{fileName: _fileName}
}

// Open the file associated to the column
func (r *FileColReader) Open() error {
	var err error
	r.file, err = os.Open(r.fileName)
	if err != nil {
		log.Printf("Error while opening %s file  with error %s\n", r.fileName, err)
	}
	return err
}

// Close the file
func (r *FileColReader) Close() error {
	var err error
	if r.file == nil {
		return nil
	}
	err = r.file.Close()
	if err != nil {
		log.Printf("Error while opening %s file  with error %s\n", r.fileName, err)
	}
	return err
}

func (r *FileColReader) checkFile() error {
	if r.file == nil {
		return os.ErrNotExist
	}
	return nil
}

// ReadInt32 read an int32 from file
func (r *FileColReader) ReadInt32() (int32, error) {
	if r.file == nil {
		return int32(0), os.ErrNotExist
	}
	//read int32 from file
	var i32 int32 = 0;
	binary.Read(r.file, binary.LittleEndian, &i32)
	return i32, nil
}
