package rscopy

import (
	"github.com/bisegni/go-c-interface-test/adapter"
)

type phase int

// all phases for the table import operation
const (
	download 	phase = itoa
	uncompress 	phase = itoa
	copy 		phase = itoa
	cleanup 	phase = itoa
)

// CopyInto is a job that will copy and external file into a table
type CopyInto struct {
	// job configuration
	config *RsCopy

	// adapter for download the file
	downloader adapter.Dowloader

	//is the initial phase
	currentPhase phase
}

// Execute into the scheduler all the phases for the copy action
func (ci *CopyInto) Execute() {

}

func (ci *CopyInto) name() string {
	return fmt.SPrintf("CopyInto %s", ci.config.Table)
}