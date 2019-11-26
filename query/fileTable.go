package query

import (
	"fmt"
)

// FileTable manage a data of a table using file for each column
type FileTable struct {
	//path where store the result
	tableFolderPath string

	schema *[]ColDescription

	columnWriter []ColWriter
}

// NewFileTable allocate new instance
func NewFileTable(tableFolderPath string, schema *[]ColDescription) (*FileTable, error) {
	ft := FileTable{tableFolderPath: tableFolderPath, schema: schema}
	err := ft.init()
	return &ft, err
}

func (ft *FileTable) init() error {
	if ft.schema == nil {
		return ErrNoSchemaInformation
	}
	// load column writer for write operation
	for _, col := range *ft.schema {
		fileName := fmt.Sprintf("%s/%s", ft.tableFolderPath, col.Name)
		fcw := NewFileColWriter(fileName, col.Kind)
		err := fcw.Open()
		if err == nil {
			ft.columnWriter = append(ft.columnWriter, fcw)
		} else {
			ft.columnWriter = nil
			return err
		}
	}
	return nil
}

// InsertRow impl.
func (ft *FileTable) InsertRow(newRow *[]interface{}) error {
	for i, v := range *newRow {
		err := ft.columnWriter[i].Write(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// SelectAll impl.
func (ft *FileTable) SelectAll() (*FileResultSet, error) {
	return NewFileResultSet(ft.tableFolderPath)
}
