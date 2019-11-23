package query

import "reflect"

// ColReader abstract interface for column readedr implementation
type ColReader interface {
	Open() error
	Close() error

	//Private methods
	ReadNext() (interface{}, error)
}

// Forwarder abstractio for the query submition implementation su sublayer
type Forwarder interface {
	// Execute start execution of the query on the backend
	Execute() error

	// Wait for the result
	Wait() (bool, error)

	// GetSchema the schema ofr the reuslt of the query
	GetSchema() (*[]ColDescription, error)

	// Close the executor
	Close()
}

// ColDescription describe a column
type ColDescription struct {
	Name string
	Kind reflect.Kind
}

// Executor abstract interface for query execution
type Executor interface {
	// Execute start execution of the query on the backend
	Execute() error

	// Wait for the result
	Wait() (bool, error)

	// GetSchema the schema ofr the reuslt of the query
	GetSchema() (*[]ColDescription, error)

	// RowCount return the number of found row
	RowCount() (int64, error)

	// NextRow return next row or error if all row are terminated
	NextRow() (*[]interface{}, error)

	// Close the executor
	Close()
}
