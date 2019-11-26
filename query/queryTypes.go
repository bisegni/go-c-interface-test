package query

import (
	"errors"
	"reflect"
)

// ColReader abstract interface for column readedr implementation
type ColReader interface {
	Open() error
	Close() error

	//Private methods
	ReadNext() (interface{}, error)
}

// TableSchemaManagement interface to folder creation implementation
type TableSchemaManagement interface {
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
}

type TableDataManageent interface {
}

var (
	// ErrTMTbaleAlredyExists the forwarder has not been set
	ErrTMTbaleAlredyExists = errors.New("The table already exists")
)

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

// ColDescription describe a column
type ColDescription struct {
	Name string       `json:"name"`
	Kind reflect.Kind `json:"kind"`
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
