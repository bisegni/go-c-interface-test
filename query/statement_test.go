package query

import (
	"testing"
)

func TestFileTableNoFolder(t *testing.T) {
	// _, err := NewFileTable("bad-path/inesistent-fodler", nil)
	// assert.Assert(t, isError(err))
}

// func TestFileTableInsert(t *testing.T) {
// 	dirpath := "data/reuslt-1"
// 	defer os.RemoveAll("data")

// 	os.MkdirAll(dirpath, os.ModePerm)

// 	schema := []ColDescription{
// 		ColDescription{
// 			"col_1",
// 			reflect.Int32},
// 		ColDescription{
// 			"col_2",
// 			reflect.Int64},
// 	}
// 	j, err := json.Marshal(schema)
// 	assert.Assert(t, !isError(err))
// 	ioutil.WriteFile(filepath.Join(dirpath, "metadata.json"), j, 0644)

// 	ft, err := NewFileTable(dirpath, &schema)
// 	assert.Assert(t, !isError(err))

// 	for i := 0; i < 100; i++ {
// 		row := []interface{}{int32(123), int64(456)}
// 		err = ft.InsertRow(&row)
// 		assert.Assert(t, !isError(err))
// 	}

// 	//get resultset
// 	rs, err := ft.SelectAll()
// 	assert.Assert(t, !isError(err))
// 	var work bool = true
// 	for work {
// 		work, err = rs.HasNext()
// 		assert.Assert(t, !isError(err))
// 		if !work {
// 			break
// 		}
// 		row, err := rs.Next()
// 		assert.Assert(t, !isError(err))
// 		assert.Equal(t, (*row)[0], int32(123))
// 		assert.Equal(t, (*row)[1], int64(456))
// 	}
// }
