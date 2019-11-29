package query

import (
	"os"
	"testing"

	"gotest.tools/assert"
)

// func TestCreateRandomTable(t *testing.T) {
// 	r := NewRandomForwarder("")
// 	err := r.Execute()
// 	assert.Assert(t, !isError(err))
// 	r.Close()
// }

func TestRandomDataGenerator(t *testing.T) {
	defer os.RemoveAll("data")
	ft := NewFileTable("data", "t1")

	//generate a random table
	ft.Create(newRandomSchema())

	is, err := ft.OpenInsertStatement()
	assert.Assert(t, !isError(err))

	rdg := NewRandomData(is)

	rdg.Execute(10000)
}
