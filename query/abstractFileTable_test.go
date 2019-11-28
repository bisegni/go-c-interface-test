package query

import (
	"os"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestAbstractTableCreateFolderCheckAndDelete(t *testing.T) {
	var e error
	at := newAbstractFileTable("data/table_1")

	//check for folder that is not present
	b, e := at.folderCheck()
	assert.Assert(t, !isError(e))
	assert.Equal(t, b, false)

	//ensure folder
	e = at.ensureFolder()
	assert.Assert(t, !isError(e))

	//check for new craeted folder
	b, e = at.folderCheck()
	assert.Assert(t, !isError(e))
	assert.Equal(t, b, true)

	//delete folder
	e = at.Delete()
	assert.Assert(t, !isError(e))

	//check for folder that is not present
	b, e = at.folderCheck()
	assert.Assert(t, !isError(e))
	assert.Equal(t, b, false)

	os.RemoveAll("data")
}

func TestCreateSchema(t *testing.T) {
	schema := []ColDescription{
		ColDescription{
			"col_1",
			reflect.Int32},
		ColDescription{
			"col_2",
			reflect.Int64},
	}
	//create table abstraction
	at := newAbstractFileTable("data/table_1")

	//ensure folder
	assert.Assert(t, !isError(at.ensureFolder()))

	//write scehma
	assert.Assert(t, !isError(at.writeSchema(&schema)))

	//load schema
	assert.Assert(t, !isError(at.loadSchema()))

	//comapre two schema
	assert.Assert(t, reflect.DeepEqual(schema, at.schema))
}

func TestCreateWriter(t *testing.T) {
	schema := []ColDescription{
		ColDescription{
			"col_1",
			reflect.Int32},
		ColDescription{
			"col_2",
			reflect.Int64},
	}
	//create table abstraction
	at := newAbstractFileTable("data/table_1")

	//ensure folder
	assert.Assert(t, !isError(at.ensureFolder()))

	//write scehma
	assert.Assert(t, !isError(at.writeSchema(&schema)))

	//load schema
	assert.Assert(t, !isError(at.loadSchema()))

	//comapre two schema
	assert.Assert(t, reflect.DeepEqual(schema, at.schema))

	//load writer
	assert.Assert(t, !isError(at.allocateColumnWriter()))

	assert.Assert(t, at.columnWriter != nil)
	assert.Assert(t, len(at.columnWriter) == 2)
}

func TestCreateReader(t *testing.T) {
	schema := []ColDescription{
		ColDescription{
			"col_1",
			reflect.Int32},
		ColDescription{
			"col_2",
			reflect.Int64},
	}
	//create table abstraction
	at := newAbstractFileTable("data/table_1")

	//ensure folder
	assert.Assert(t, !isError(at.ensureFolder()))

	//write scehma
	assert.Assert(t, !isError(at.writeSchema(&schema)))

	//load schema
	assert.Assert(t, !isError(at.loadSchema()))

	//comapre two schema
	assert.Assert(t, reflect.DeepEqual(schema, at.schema))

	//load writer
	assert.Assert(t, !isError(at.allocateColumnReader()))

	assert.Assert(t, at.columnReader != nil)
	assert.Assert(t, len(at.columnReader) == 2)
}
