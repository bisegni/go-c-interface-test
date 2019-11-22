package query

// ColReader abstract interface for column readedr implementation
type ColReader interface {
	Open() error
	Close() error

	//Private methods
	ReadNext() (interface{}, error)
}
