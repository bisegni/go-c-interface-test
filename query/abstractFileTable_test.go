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
	defer func() {
		at.Close()
		os.RemoveAll("data") // change value at the very last moment
	}()
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
	defer func() {
		at.Close()
		os.RemoveAll("data") // change value at the very last moment
	}()
	//ensure folder
	assert.Assert(t, !isError(at.ensureFolder()))

	//write scehma
	assert.Assert(t, !isError(at.writeSchema(&schema)))

	//load schema
	assert.Assert(t, !isError(at.loadSchema()))

	//comapre two schema
	assert.Assert(t, reflect.DeepEqual(schema, at.schema))
}

func TestCreateWriterAndReader(t *testing.T) {
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
	defer func() {
		at.Close()
		os.RemoveAll("data") // change value at the very last moment
	}()
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

	//load writer
	assert.Assert(t, !isError(at.allocateColumnReader()))

	assert.Assert(t, at.columnReader != nil)
	assert.Assert(t, len(at.columnReader) == 2)
}

func TestAbstractTableStatistic(t *testing.T) {
	at := newAbstractFileTable("data/table_1")
	defer func() {
		at.Close()
		os.RemoveAll("data") // change value at the very last moment
	}()

	stat := at.GetStatistics()
	assert.Assert(t, len(stat.column) == 1)
	assert.Assert(t, len(stat.values) == 1)
	at.addStat(&ColDescription{"a", reflect.Bool}, true)
	at.addStat(&ColDescription{"b", reflect.Int32}, 12345)

	stat = at.GetStatistics()
	assert.Assert(t, len(stat.column) == 3)
	assert.Assert(t, stat.column[1].Name == "a")
	assert.Assert(t, stat.column[1].Kind == reflect.Bool)
	assert.Assert(t, stat.column[2].Name == "b")
	assert.Assert(t, stat.column[2].Kind == reflect.Int32)
	assert.Assert(t, len(stat.values) == 3)
	assert.Assert(t, stat.values[1] == true)
	assert.Assert(t, stat.values[2] == 12345)
}
