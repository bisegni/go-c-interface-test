package query

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestFileTableCreation(t *testing.T) {
	r := NewFileTableManagement("data", "table_1")
	schema := []ColDescription{
		ColDescription{
			"col_1",
			reflect.Int32},
		ColDescription{
			"col_2",
			reflect.Int64},
	}
	j, err := json.Marshal(schema)
	assert.Assert(t, !isError(err))

	t.Logf("Create table with schema %v", string(j))
	err = r.Create(&schema)
	t.Logf("Read table schema %v", string(j))
	readedSchema, err := r.GetSchema()
	assert.Assert(t, !isError(err))

	j, err = json.Marshal(schema)
	assert.Assert(t, !isError(err))
	t.Logf("Readed schema %v", string(j))
	assert.Assert(t, reflect.DeepEqual(schema, *readedSchema))

	//delete table
	t.Logf("Delete table %s", r.tableFolderPath)
	err = r.Delete()
	assert.Assert(t, !isError(err))

	os.RemoveAll("data")
}
