package query

// func TestCreateRandomTable(t *testing.T) {
// 	r := NewRandomForwarder("")
// 	err := r.Execute()
// 	assert.Assert(t, !isError(err))
// 	r.Close()
// }

// func TestCreateRandomTableAndReadColumn(t *testing.T) {
// 	r := NewRandomForwarder("")
// 	defer r.Close()

// 	err := r.Execute()
// 	assert.Assert(t, !isError(err))

// 	//star reading column one by one
// 	colDesc, err := r.GetSchema()
// 	assert.Assert(t, !isError(err))

// 	var scanErr error
// 	var foundRow int64 = 0
// 	for _, desc := range *colDesc {
// 		foundRow = 0
// 		fcr := NewFileColReader(desc.Name, desc.Kind)
// 		scanErr = fcr.Open()
// 		assert.Assert(t, !isError(scanErr))
// 		for scanErr == nil {
// 			_, scanErr = fcr.ReadNext()
// 			if scanErr == nil {
// 				foundRow++
// 			}
// 		}
// 		r, e := r.GetRowCount()
// 		assert.Assert(t, !isError(e))
// 		assert.Assert(t, r == foundRow)
// 		scanErr = fcr.Close()
// 		assert.Assert(t, !isError(scanErr))
// 	}
// }
