package query

import "errors"

// JobStateType represent
type JobStateType int32

const (
	// JobStateCreated identify a job created but still not executed
	JobStateCreated JobStateType = iota
	// JobStateExecuting identify a job that is executing and is or will produce data
	JobStateExecuting
	// JobStateFinished the job is complete his work with no error
	JobStateFinished
	// JobStateFault the job has fault
	JobStateFault
)

var (
	// ErrorJobInstanceAlreadyCreated specify that an instance has been already created
	ErrorJobInstanceAlreadyCreated = errors.New("Instance already created")
)

// Job implementation
/*
	Execute asynchronously a query creating a table with the result
	The job create a folder that can be used as table using @table.FileTable
*/
type Job struct {
	path string
	name string
	// Query define the asynchronous query
	Query string `json:"query"`
	//! identify the state of the job
	State JobStateType `json:"state"`
}

// NewJob create new instance
func NewJob(path string, name string, query string) *Job {
	return &Job{
		path:  path,
		name:  name,
		Query: query}
}

// Execute launch the job
/*
 The job consists in a folder that represent a table that is filled
 with the result of the query. If that table exist this function return an error
*/
func (j *Job) Execute() error {
	return nil
}
