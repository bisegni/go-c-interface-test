package query

// InsertStatement manage a data of a table using file for each column
type InsertStatement struct {
	//path where store the result
	schema *[]ColDescription

	columnWriter []ColWriter
}

// NewFileTable allocate new instance
func newInsertStatement(schema *[]ColDescription, columnWriter []ColWriter) *InsertStatement {
	return &InsertStatement{schema: schema, columnWriter: columnWriter}
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

// SelectStatement impl.
type SelectStatement struct {
	//path where store the result
	schema *[]ColDescription

	columnReader []ColReader
}

func newSelectStatement(schema *[]ColDescription, columnReader []ColReader) *SelectStatement {
	return &SelectStatement{schema: schema, columnReader: columnReader}
}

// SelectAll impl.
func (st *SelectStatement) SelectAll() (*FileResultSet, error) {
	return newFileResultSet(st.schema, st.columnReader), nil
}
