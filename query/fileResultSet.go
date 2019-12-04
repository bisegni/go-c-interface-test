package query

import (
	"errors"
	"io"
)

var (
	// ErrFRSNoSchemaFound No schema has been loaded, perhaps metadata file is missing or empty
	ErrFRSNoSchemaFound = errors.New("No schema has been loaded, perhaps metadata file is missing or empty")

	// ErrFRSNoColumnReader No column reader has been configured
	ErrFRSNoColumnReader = errors.New("No column reader has been configured")

	// ErrFRSColumnReadError Error during read next column value
	ErrFRSColumnReadError = errors.New("Error during read next column value")

	// ErrFRSNoRowFetched No row has been fetched by HasNext method
	ErrFRSNoRowFetched = errors.New("No row has been fetched by HasNext method")
)

// FileResultSet implement a result set usign file as query result
/*
This implementation of result set sue a metadata.json file to read the type of result
and a file for each column that contains the result
*/
type FileResultSet struct {
	//contains the schema of the table/virtual table
	schema *[]ColDescription

	columnReader []ColReader

	//point to current row fetched with HasNext method
	currentRow []interface{}
}

// NewFileResultSet allocate new instance
/*
path param is the path where data is stored
table table is the name of a table or a virtual table that is the result of a query
*/
func newFileResultSet(schema *[]ColDescription, columnReader []ColReader) *FileResultSet {
	return &FileResultSet{schema: schema, columnReader: columnReader}
}

// GetSchema impl.
func (frs *FileResultSet) GetSchema() (*[]ColDescription, error) {
	if len(*frs.schema) == 0 {
		return nil, ErrFRSNoSchemaFound
	}
	return frs.schema, nil
}

// HasNext impl.
func (frs *FileResultSet) HasNext() (bool, error) {
	var err error
	var hasNext bool = true
	var val interface{}
	frs.currentRow = make([]interface{}, len(*frs.schema))
	if frs.columnReader == nil {
		return false, ErrFRSNoColumnReader
	}
	//we can read next row
	for i, cr := range frs.columnReader {
		val, err = cr.ReadNext()
		if err != nil {
			frs.currentRow = nil
			hasNext = false
			if err != io.EOF &&
				err != ErrRotateReaderNoMoreChunk {
				err = ErrFRSColumnReadError
			} else {
				// no error on end of file will be forwarded
				err = nil
			}
			break
		}
		frs.currentRow[i] = val
	}
	return hasNext, err
}

// Next impl.
func (frs *FileResultSet) Next() (*[]interface{}, error) {
	if frs.currentRow == nil {
		return nil, ErrFRSNoRowFetched
	}
	return &frs.currentRow, nil
}

// Close impl.
func (frs *FileResultSet) Close() error {
	for _, cr := range frs.columnReader {
		cr.Close()
	}
	return nil
}
