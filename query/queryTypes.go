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

	//Private methods
	ReadNext() (interface{}, error)
}

var (
	// ErrCWTBadType wrong tipy passed to column writer
	ErrCWTBadType = errors.New("Wrong tipy passed to column writer")
)

// ColWriter abstract interface for column writer implementation
type ColWriter interface {
	Open() error
	Close() error
	Writer(interface{}) error
}

var (
	// ErrTMTbaleAlredyExists The table already exists
	ErrTMTbaleAlredyExists = errors.New("The table already exists")
)

// TableManagement interface to folder creation implementation
type TableManagement interface {
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

	// InsertRow add new row within the table
	InsertRow(*[]interface{})
}

// Forwarder abstraction for the query submition implementation to a sublayer
type Forwarder interface {
	// Execute start execution of the query on the backend
	Execute() error

	// GetSchema the schema ofr the reuslt of the query
	GetSchema() (*[]ColDescription, error)

	// GetRowCount return the row that have been found
	GetRowCount() (int64, error)

	// Close the executor
	Close()
}

// Executor abstract interface for query execution
type Executor interface {
	// Execute start execution of the query on the backend
	Execute() error

	// Wait for the result
	Wait() (bool, error)

	// GetSchema the schema ofr the reuslt of the query
	GetSchema() (*[]ColDescription, error)

	// GetRowCount return the number of found row
	GeRowCount() (int64, error)

	// NextRow return next row or error if all row are terminated
	NextRow() (*[]interface{}, error)

	// Close the executor
	Close()
}
