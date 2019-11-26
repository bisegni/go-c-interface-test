package query

import (
	"encoding/binary"
	"log"
	"os"
	"reflect"
)

// FileColWriter read the values of a column from a file
/*
	The colType insturct the struct to selecte wich type it need to
	search on binary files
*/
type FileColWriter struct {
	fileName string
	colType  reflect.Kind
	file     *os.File
}

// NewFileColWriter allocate new instance
func NewFileColWriter(_fileName string, _kind reflect.Kind) *FileColWriter {
	return &FileColWriter{
		fileName: _fileName,
		colType:  _kind}
}

// Open the file associated to the column
func (w *FileColWriter) Open() error {
	var err error
	w.file, err = os.Create(w.fileName)

	if err != nil {
		w.file = nil
		log.Printf("Error while opening %s file  with error %s\n", w.fileName, err)
	}
	return err
}

// Close the file
func (w *FileColWriter) Close() error {
	var err error
	if w.file == nil {
		return nil
	}
	err = w.file.Close()
	if err != nil {
		log.Printf("Error while opening %s file  with error %s\n", w.fileName, err)
	}
	return err
}

// ReadNext read an int32 from file
func (w *FileColWriter) Write(data interface{}) error {
	if w.file == nil {
		return os.ErrNotExist
	}
	//read int32 from file
	// var err error
	if reflect.ValueOf(data).Kind() != w.colType {
		return ErrCWTBadType
	}
	switch w.colType {
	case reflect.Bool:
		b := data.(bool)
		return binary.Write(w.file, binary.LittleEndian, &b)
	case reflect.Int32:
		i32 := data.(int32)
		return binary.Write(w.file, binary.LittleEndian, &i32)
	case reflect.Int64:
		i64 := data.(int64)
		return binary.Write(w.file, binary.LittleEndian, &i64)
	case reflect.Float32:
		f32 := data.(float32)
		return binary.Write(w.file, binary.LittleEndian, &f32)
	case reflect.Float64:
		f64 := data.(float64)
		return binary.Write(w.file, binary.LittleEndian, &f64)
	case reflect.String:
		// var i64 int64 = 0
		// binary.Read(r.file, binary.LittleEndian, &i64)
		// result = i64
	}
	return ErrCWTBadType
}
