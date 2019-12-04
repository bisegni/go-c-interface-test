package query

import (
	"os"
	"testing"
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

	rdg := NewRandomData(ft)

	rdg.Execute(10000)
}
