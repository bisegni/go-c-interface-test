package table

import (
	"os"
	"reflect"
)

// FileColWriter read the values of a column from a file
/*
	The colType instruct the struct to selected which type it need to
	search on binary files
*/
type FileColWriter struct {
	path         string
	fileName     string
	colType      reflect.Kind
	rotateWriter *rotateWriter
}

// NewFileColWriter allocate new instance
func NewFileColWriter(path string, fileName string, kind reflect.Kind) *FileColWriter {
	return &FileColWriter{
		path:         path,
		fileName:     fileName,
		colType:      kind,
		rotateWriter: newRotateWriterNoInit(path, fileName)}

}

// Open the file associated to the column
func (w *FileColWriter) Open() error {
	if w.rotateWriter == nil {
		return os.ErrNotExist
	}
	err := w.rotateWriter.init()
	if err != nil {
		return err
	}

	//give a rotate in case we have no file
	return w.rotateWriter.Rotate()
}

// Close the file
func (w *FileColWriter) Close() error {
	if w.rotateWriter == nil {
		return os.ErrNotExist
	}
	return w.rotateWriter.Close()
}

// ReadNext read an int32 from file
func (w *FileColWriter) Write(data interface{}) error {
	if w.rotateWriter == nil {
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
		return w.rotateWriter.Write(b)
	case reflect.Int32:
		i32 := data.(int32)
		return w.rotateWriter.Write(i32)
	case reflect.Int64:
		i64 := data.(int64)
		return w.rotateWriter.Write(i64)
	case reflect.Float32:
		f32 := data.(float32)
		return w.rotateWriter.Write(f32)
	case reflect.Float64:
		f64 := data.(float64)
		return w.rotateWriter.Write(f64)
	case reflect.String:
		// var i64 int64 = 0
		// binary.Read(r.file, binary.LittleEndian, &i64)
		// result = i64
	}
	return ErrCWTBadType
}
