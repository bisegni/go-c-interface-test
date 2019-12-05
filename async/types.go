package async

// Runnable define method to cal in a job that need to execute async operation
type Runnable interface {
	execute()
	name() string
}
