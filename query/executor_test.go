package query

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestMain(m *testing.M) {
	//disable log during tests
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestOpenRandoExecutor(t *testing.T) {
	r := NewFileExecutorWithRGA("executor_a")

	//execute query
	err := r.Execute()
	assert.NilError(t, err)

	b, err := r.Wait()
	assert.Assert(t, b)
	assert.Assert(t, !isError(err))

	colDec, err := r.GetSchema()
	assert.NilError(t, err)

	rowToExpect, err := r.GeRowCount()
	assert.NilError(t, err)

	var rowFetcheCount int64 = 0
	for err == nil {
		row, err := r.NextRow()
		assert.NilError(t, err)

		if len(*row) == 0 {
			break
		}

		rowFetcheCount++
		//check row
		assert.Equal(t, len(*colDec), len(*row))

		//check types

		for i := 0; i < len(*colDec); i++ {
			assert.Equal(t, (*colDec)[i].Kind, reflect.ValueOf((*row)[i]).Kind())
		}
	}

	assert.Equal(t, rowToExpect, rowFetcheCount)
	r.Close()
}
