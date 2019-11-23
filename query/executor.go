package query

import (
	"reflect"
)

// FileExecutor execute a query using file as column value
type FileExecutor struct {
	//! path where query result are expected
	path string
	//perform the query request to sublayer
	executor Executor
	//column reader
	columnReader []FileColReader
}

// NewQueryFileExecutor allocate new instance
func NewQueryFileExecutor(_fileName string, _kind reflect.Kind) *FileColReader {
	return &FileColReader{
		fileName: _fileName,
		colType:  _kind}
}
