package query

import (
	"log"
	"os"
	"reflect"
)

// FileColReader read the values of a column from a file
/*
	The colType insturct the struct to selecte wich type it need to
	search on binary files
*/
type FileColReader struct {
	path         string
	fileName     string
	colType      reflect.Kind
	rotateReader *rotateReader
}

// check if the file is present
func (r *FileColReader) checkFile() error {
	if r.rotateReader == nil {
		return os.ErrNotExist
	}
	return nil
}

// NewFileColReader allocate new instance
func NewFileColReader(path string, fileName string, kind reflect.Kind) *FileColReader {
	return &FileColReader{
		path:         path,
		fileName:     fileName,
		colType:      kind,
		rotateReader: newRotateReaderNoInit(path, fileName)}
}

// Open the file associated to the column
func (r *FileColReader) Open() (err error) {
	err = r.rotateReader.updateChunkInfo()
	if err != nil {
		log.Printf("Error while opening %s file  with error %s\n", r.fileName, err)
	}
	return err
}

// Close the file
func (r *FileColReader) Close() error {
	var err error
	if r.rotateReader == nil {
		return nil
	}
	err = r.rotateReader.Close()
	if err != nil {
		log.Printf("Error while opening %s file  with error %s\n", r.fileName, err)
	}
	return err
}

// ReadNext read an int32 from file
func (r *FileColReader) ReadNext() (interface{}, error) {
	if r.rotateReader == nil {
		return int32(0), os.ErrNotExist
	}
	//read int32 from file
	var result interface{}
	var err error
	switch r.colType {
	case reflect.Bool:
		var b bool = false
		err = r.rotateReader.Read(&b)
		result = b
	case reflect.Int32:
		var i32 int32 = 0
		err = r.rotateReader.Read(&i32)
		result = i32
	case reflect.Int64:
		var i64 int64 = 0
		err = r.rotateReader.Read(&i64)
		result = i64
	case reflect.Float32:
		var f32 float32 = 0
		err = r.rotateReader.Read(&f32)
		result = f32
	case reflect.Float64:
		var f64 float64 = 0
		err = r.rotateReader.Read(&f64)
		result = f64
	case reflect.String:
		// var i64 int64 = 0
		// binary.Read(&i64)
		// result = i64
	}
	return result, err
}
