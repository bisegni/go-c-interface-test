package query

import (
	"errors"
	"reflect"
)

// ColDescription conains the column specification
type ColDescription struct {
	Name string       `json:"name"`
	Kind reflect.Kind `json:"kind"`
}

// ColReader abstract interface for column readedr implementation
type ColReader interface {
	Open() error
	Close() error
	ReadNext() (interface{}, error)
}

var (
	// ErrCWTBadType wrong tipy passed to column writer
	ErrCWTBadType = errors.New("Wrong type passed to column writer")
)

//chunk base file size
const columnChunkFileSize = 1024 * 1024 // 1Mbyte file size

// ColWriter abstract interface for column writer implementation
type ColWriter interface {
	Open() error
	Close() error
	Write(interface{}) error
}

// ResultSet is the abstraction of a cursor
type ResultSet interface {
	GetSchema() (*[]ColDescription, error)
	HasNext() (bool, error)
	Next() (*[]interface{}, error)
	Close() error
}

var (
	// ErrNoSchemaInformation The table has no schema informatio
	ErrNoSchemaInformation = errors.New("The table has no schema information")
)

var (
	// ErrTMTableAlredyExists The table already exists
	ErrTMTableAlredyExists = errors.New("The table already exists")

	// ErrTMSChemaMetadataNotFount The metadata information has not been found
	ErrTMSChemaMetadataNotFount = errors.New("The metadata information has not been found")
)

// StatisticResult return the statics values for the table
type StatisticResult struct {
	column []ColDescription
	values []interface{}
}

// Table interface to folder management implementation
type Table interface {
	// Create a table
	/*
		if table is alredy preset an error is issue
	*/
	Create(*[]ColDescription) error

	// Delete the table structure
	/*
		Intere table structure will be deleted
	*/
	Delete() error

	// GetSchema return the table schema
	GetSchema() (*[]ColDescription, error)

	// GetStatistics return the statistics for the table
	GetStatistics() *StatisticResult

	//Insert a new row in table
	InsertRow(newRow *[]interface{}) error

	//OpenSelectStatement create new select statement
	OpenSelectStatement() (*SelectStatement, error)

	// Close the table access file
	Close()
}
