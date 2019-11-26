package query

import (
	"fmt"
	"reflect"
)

// FileTable manage a data of a table using file for each column
type FileTable struct {
	//path where store the result
	tableFolderPath string

	schema *[]ColDescription

	columnWriter []*ColWriter
}

// NewFileTable allocate new instance
func NewFileTable(tableFolderPath string, schema *[]ColDescription) (*FileTable, error) {
	ft := FileTable{tableFolderPath: tableFolderPath, schema: schema}
	err := ft.init()
	return &ft, err
}

func (ft *FileTable) init() error {
	// load column writer for write operation
	for _, col := range *ft.schema {
		fileName := fmt.Sprintf("%s/%s", ft.tableFolderPath, col.Name)
		fcw := NewFileColWriter(fileName, col.Kind)
		err := fcw.Open()
		if err == nil {
			ft.columnWriter = append(ft.columnWriter, reflect.ValueOf(fcw).Interface().(*ColWriter))
		} else {
			ft.columnWriter = nil
			return err
		}
	}
	return nil
}

// InsertRow impl.
func (ft *FileTable) InsertRow(newRow *[]interface{}) error {
	return nil
}

// SelectAll impl.
func (ft *FileTable) SelectAll() (*ResultSet, error) {
	return nil, nil
}
