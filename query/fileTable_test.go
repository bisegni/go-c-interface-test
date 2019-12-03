package query

import (
	"os"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestManagement(t *testing.T) {
	r := NewFileTable("data", "table_1")
	defer os.RemoveAll("data")

	schema := []ColDescription{
		ColDescription{
			"col_1",
			reflect.Int32},
		ColDescription{
			"col_2",
			reflect.Int64},
	}
	err := r.Create(&schema)
	readedSchema, err := r.GetSchema()
	assert.Assert(t, !isError(err))
	assert.Assert(t, reflect.DeepEqual(schema, *readedSchema))

	//delete table
	err = r.Delete()
	assert.Assert(t, !isError(err))
}

func TestInsert(t *testing.T) {
	r := NewFileTable("data", "table_1")
	defer os.RemoveAll("data")

	schema := []ColDescription{
		ColDescription{
			"col_1",
			reflect.Int32},
		ColDescription{
			"col_2",
			reflect.Int64},
	}
	err := r.Create(&schema)
	assert.Assert(t, !isError(err))

	is, err := r.OpenInsertStatement()
	assert.Assert(t, !isError(err))

	gotSchema := is.GetSchema()
	assert.Assert(t, reflect.DeepEqual(schema, *gotSchema))

	for i := 0; i < 100; i++ {
		row := []interface{}{int32(i), int64(i + 2)}
		err = is.InsertRow(&row)
		assert.Assert(t, !isError(err))
	}

	ss, err := r.OpenSelectStatement()
	assert.Assert(t, !isError(err))

	rs, err := ss.SelectAll()
	assert.Assert(t, !isError(err))

	var work bool = true
	var idx int32 = 0
	for work {
		work, err = rs.HasNext()
		assert.Assert(t, !isError(err))
		if !work {
			break
		}
		row, err := rs.Next()
		assert.Assert(t, !isError(err))
		assert.Equal(t, (*row)[0], int32(idx))
		assert.Equal(t, (*row)[1], int64(idx+2))
		idx++
	}
}
