package query

import (
	"encoding/binary"
	"log"
	"os"
	"reflect"
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
func NewFileColReader(_fileName string, _kind reflect.Kind) *FileColReader {
	return &FileColReader{
		fileName: _fileName,
		colType:  _kind}
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

// ReadNext read an int32 from file
func (r *FileColReader) ReadNext() (interface{}, error) {
	if r.file == nil {
		return int32(0), os.ErrNotExist
	}
	//read int32 from file
	var result interface{}
	var err error
	switch r.colType {
	case reflect.Bool:
		var b bool = false
		err = binary.Read(r.file, binary.LittleEndian, &b)
		result = b
		break
	case reflect.Int32:
		var i32 int32 = 0
		err = binary.Read(r.file, binary.LittleEndian, &i32)
		result = i32
		break
	case reflect.Int64:
		var i64 int64 = 0
		err = binary.Read(r.file, binary.LittleEndian, &i64)
		result = i64
		break
	case reflect.Float32:
		var f32 float32 = 0
		err = binary.Read(r.file, binary.LittleEndian, &f32)
		result = f32
		break
	case reflect.Float64:
		var f64 float64 = 0
		err = binary.Read(r.file, binary.LittleEndian, &f64)
		result = f64
		break
	case reflect.String:
		// var i64 int64 = 0
		// binary.Read(r.file, binary.LittleEndian, &i64)
		// result = i64
		break
	}
	return result, err
}
