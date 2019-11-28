package query

import (
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestAbstractTableCreateFolderCheckAndDelete(t *testing.T) {
	var e error
	// schema := []ColDescription{
	// 	ColDescription{
	// 		"col_1",
	// 		reflect.Int32},
	// 	ColDescription{
	// 		"col_2",
	// 		reflect.Int64},
	// }
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
