package query

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	//ErrFENoColumnReader column reader are not be created, perhaps wait method has not been create
	ErrFENoColumnReader = errors.New("No Column reader found")
	// ErrFEColumnReadError error has ben found during fetch data from a column reader
	ErrFEColumnReadError = errors.New("Error during fetch data froma column file")
)

// FileExecutor execute a query using file as column value
type FileExecutor struct {
	//! path where query result are expected
	path string
	//forwarder query
	forwarder Forwarder
	//column reader form file column answer
	columnReader []*FileColReader
}

// NewFileExecutorWithRGA give a query executor that generate a randomly select answer
func NewFileExecutorWithRGA(_path string) *FileExecutor {
	return &FileExecutor{
		path:      _path,
		forwarder: NewRandomForwarder(_path)}
}

// Execute run the query
func (fe *FileExecutor) Execute() error {
	//let path was present befor start
	err := os.MkdirAll(fe.path, os.ModeDir|os.ModePerm)
	if err == nil {
		err = fe.forwarder.Execute()
	}
	return err
}

// Wait until query produce result
func (fe *FileExecutor) Wait() (bool, error) {
	//check for file of the results
	var err error
	var result bool = true
	resultSchema, err := fe.forwarder.GetSchema()
	if err != nil {
		return false, err
	}

	//allocate column reader
	for _, desc := range *resultSchema {
		var fName string
		if len(fe.path) > 0 {
			fName = fmt.Sprintf("%s/%s", fe.path, desc.Name)
		} else {
			fName = fmt.Sprintf("%s", desc.Name)
		}
		fcr := NewFileColReader(fName, desc.Kind)
		err = fcr.Open()
		if err == nil {
			fe.columnReader = append(fe.columnReader, fcr)
		} else {
			result = false
			break
		}
	}
	return result, err
}

// GetSchema return the schema for the found result
func (fe *FileExecutor) GetSchema() (*[]ColDescription, error) {
	return fe.forwarder.GetSchema()
}

// GeRowCount return the number of row found
func (fe *FileExecutor) GeRowCount() (int64, error) {
	return fe.forwarder.GetRowCount()
}

// NextRow return next row or error if all row are terminated
func (fe *FileExecutor) NextRow() (*[]interface{}, error) {
	var err error
	var val interface{}
	// scan each next row of col read and for the row
	schema, err := fe.forwarder.GetSchema()
	if err != nil {
		return nil, err
	}

	row := make([]interface{}, len(*schema))
	if len(fe.columnReader) == 0 {
		return nil, ErrFENoColumnReader
	}

	//we can read next row
	for i, cr := range fe.columnReader {
		val, err = cr.ReadNext()
		if err != nil {
			row = nil
			if err != io.EOF {
				err = ErrFEColumnReadError
			} else {
				// no error on end of file will be forwarded
				err = nil
			}
			break
		}
		row[i] = val
	}
	return &row, err
}

// Close the executor
func (fe *FileExecutor) Close() {
	//close forwarder
	fe.forwarder.Close()
	//close column reader
	for _, fcr := range fe.columnReader {
		fcr.Close()
	}

	//remove executor path
	os.RemoveAll(fe.path)
}
