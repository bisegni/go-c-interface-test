package query

import (
	"path/filepath"
)

// FileTable manage a table using folder and files
type FileTable struct {
	AbstractFileTable
}

// NewFileTable allocate new instance
func NewFileTable(_path string, _name string) *FileTable {
	return &FileTable{
		AbstractFileTable: newAbstractFileTable(filepath.Join(_path, _name)),
	}
}

// Create impl.
func (ft *FileTable) Create(schema *[]ColDescription) error {
	return ft.writeSchema(schema)
}

// GetSchema impl.
func (ft *FileTable) GetSchema() (*[]ColDescription, error) {
	err := ft.loadSchema()
	return &ft.schema, err
}

// OpenInsertStatement impl.
func (ft *FileTable) OpenInsertStatement() (is *InsertStatement, err error) {
	var cw *[]ColWriter
	if err := ft.loadSchema(); err != nil {
		return nil, err
	}
	if cw, err = ft.allocateColumnWriter(); err != nil {
		return nil, err
	}
	is, err = newInsertStatement(&ft.schema, *cw), nil
	return
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
