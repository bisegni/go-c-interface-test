package table

import (
	"errors"
	"path/filepath"
	"sync"
)

// FileTable manage a table using folder and files
type FileTable struct {
	AbstractFileTable
	writeMutex sync.Mutex
}

var (
	// ErrFileTableRowMIsmatchSchema The table already exists
	ErrFileTableRowMIsmatchSchema = errors.New("The row element number differ from column number")
)

// NewFileTable allocate new instance
func NewFileTable(_path string, _name string) *FileTable {
	return &FileTable{
		AbstractFileTable: newAbstractFileTable(filepath.Join(_path, _name)),
	}
}

// Create impl.
func (ft *FileTable) Create(schema *[]ColDescription) error {
	err := ft.writeSchema(schema)
	if err == nil {
		err = ft.loadSchema()
	}
	return err
}

// GetSchema impl.
func (ft *FileTable) GetSchema() (*[]ColDescription, error) {
	err := ft.loadSchema()
	return &ft.schema, err
}

// InsertRow impl.
func (ft *FileTable) InsertRow(newRow *[]interface{}) (err error) {
	ft.writeMutex.Lock()
	var cw *[]ColWriter

	defer func() {
		ft.writeMutex.Unlock()
		for _, w := range *cw {
			w.Close()
		}
	}()

	if cw, err = ft.allocateColumnWriter(); err != nil {
		return err
	}
	if len(*newRow) != len(*cw) {
		return ErrFileTableRowMIsmatchSchema
	}
	for i, v := range *newRow {
		err := (*cw)[i].Write(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// OpenSelectStatement impl.
func (ft *FileTable) OpenSelectStatement() (ss *SelectStatement, err error) {
	var cr *[]ColReader
	if err := ft.loadSchema(); err != nil {
		return nil, err
	}
	if cr, err = ft.allocateColumnReader(); err != nil {
		return nil, err
	}
	ss = newSelectStatement(&ft.schema, *cr)
	return
}
