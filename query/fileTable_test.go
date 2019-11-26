package query

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestFileTableNoFolder(t *testing.T) {
	_, err := NewFileTable("bad-path/inesistent-fodler", nil)
	assert.Assert(t, isError(err))
}

func TestFileTableInsert(t *testing.T) {
	dirpath := "data/reuslt-1"
	defer os.RemoveAll("data")

	os.MkdirAll(dirpath, os.ModePerm)

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
	ioutil.WriteFile(filepath.Join(dirpath, "metadata.json"), j, 0644)
}
