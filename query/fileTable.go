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
func (ft *FileTable) OpenInsertStatement() (*InsertStatement, error) {
	if err := ft.loadSchema(); err != nil {
		return nil, err
	}
	if err := ft.allocateColumnWriter(); err != nil {
		return nil, err
	}
	return newInsertStatement(&ft.schema, ft.columnWriter), nil
}

// OpenSelectStatement impl.
func (ft *FileTable) OpenSelectStatement() (*SelectStatement, error) {
	if err := ft.loadSchema(); err != nil {
		return nil, err
	}
	if err := ft.allocateColumnReader(); err != nil {
		return nil, err
	}
	return newSelectStatement(&ft.schema, ft.columnReader), nil
}
