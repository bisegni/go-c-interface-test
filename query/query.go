package query

// ColReader interface for column readedr implementation
type ColReader interface {
	Open() error
	Close() error

	//Private methods
	ReadInt32() (int32, error)
}
