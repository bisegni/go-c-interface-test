package query

import (
	"encoding/binary"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	"gotest.tools/assert"
)

func isError(err error) bool {
	return (err != nil)
}

func randByType(f *os.File, t reflect.Kind) []interface{} {
	var generatedIntArray []interface{}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		switch t {
		case reflect.Bool:
		case reflect.Int32:
			generatedIntArray[i] = rand.Int31()
		case reflect.Int64:
			generatedIntArray[i] = rand.Int63()
		case reflect.Float32:
			generatedIntArray[i] = rand.Float32()
		case reflect.Float64:
			generatedIntArray[i] = rand.Float64()
		case reflect.String:
			// var i64 int64 = 0
			// binary.Read(r.file, binary.LittleEndian, &i64)
			// result = i64
		}

		binary.Write(f, binary.LittleEndian, generatedIntArray[i])
	}
	return generatedIntArray
}

func TestOpenNoFoundFile(t *testing.T) {
	defer os.RemoveAll("data")
	r := NewFileColReader("data", "file_not_present.bin", reflect.Bool)
	err := r.Open()
	assert.Assert(t, isError(err))
}

func TestReadBool(t *testing.T) {
	defer os.RemoveAll("data")
	//create file for test
	var generatedIntArray [1000]bool
	fcw := NewFileColWriter("data", "file_bool", reflect.Bool)
	err := fcw.Open()
	assert.Assert(t, !isError(err))

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = (rand.Intn(2) != 0)
		fcw.Write(generatedIntArray[i])
	}
	fcw.Close()

	r := NewFileColReader("data", "file_bool", reflect.Bool)
	err = r.Open()
	assert.Assert(t, !isError(err))

	for i := 0; i < 1000; i++ {
		iread, ie := r.ReadNext()
		assert.Assert(t, !isError(ie))
		assert.Assert(t, generatedIntArray[i] == iread)
	}

	//next call to read nex need to give error
	_, ie := r.ReadNext()
	assert.Assert(t, isError(ie))
}

func TestReadInt32(t *testing.T) {
	defer os.RemoveAll("data")
	//create file for test
	var generatedIntArray [1000]int32

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	fcw := NewFileColWriter("data", "file_int32", reflect.Int32)
	err := fcw.Open()
	assert.Assert(t, !isError(err))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = rand.Int31()
		fcw.Write(generatedIntArray[i])
	}
	fcw.Close()

	r := NewFileColReader("data", "file_int32", reflect.Int32)
	err = r.Open()
	assert.Assert(t, !isError(err))

	for i := 0; i < 1000; i++ {
		iread, ie := r.ReadNext()
		assert.Assert(t, !isError(ie))
		assert.Assert(t, generatedIntArray[i] == iread)
	}

	//next call to read nex need to give error
	_, ie := r.ReadNext()
	assert.Assert(t, isError(ie))
}

func TestReadInt64(t *testing.T) {
	defer os.RemoveAll("data")
	//create file for test
	var generatedIntArray [1000]int64
	fcw := NewFileColWriter("data", "file_int64", reflect.Int64)
	err := fcw.Open()
	assert.Assert(t, !isError(err))

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = rand.Int63()
		fcw.Write(generatedIntArray[i])
	}
	fcw.Close()

	r := NewFileColReader("data", "file_int64", reflect.Int64)
	err = r.Open()
	assert.Assert(t, !isError(err))

	for i := 0; i < 1000; i++ {
		iread, ie := r.ReadNext()
		assert.Assert(t, !isError(ie))
		assert.Assert(t, generatedIntArray[i] == iread)
	}

	//next call to read nex need to give error
	_, ie := r.ReadNext()
	assert.Assert(t, isError(ie))
}

func TestReadFloat32(t *testing.T) {
	defer os.RemoveAll("data")
	//create file for test
	var generatedIntArray [1000]float32
	fcw := NewFileColWriter("data", "file_float32", reflect.Float32)
	err := fcw.Open()
	assert.Assert(t, !isError(err))

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = rand.Float32()
		fcw.Write(generatedIntArray[i])
	}
	fcw.Close()

	r := NewFileColReader("data", "file_float32", reflect.Float32)
	err = r.Open()
	assert.Assert(t, !isError(err))

	for i := 0; i < 1000; i++ {
		iread, ie := r.ReadNext()
		assert.Assert(t, !isError(ie))
		assert.Assert(t, generatedIntArray[i] == iread)
	}

	//next call to read nex need to give error
	_, ie := r.ReadNext()
	assert.Assert(t, isError(ie))
}

func TestReadFloat64(t *testing.T) {
	defer os.RemoveAll("data")
	//create file for test
	var generatedIntArray [1000]float64
	fcw := NewFileColWriter("data", "file_float64", reflect.Float64)
	err := fcw.Open()
	assert.Assert(t, !isError(err))

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = rand.Float64()
		fcw.Write(generatedIntArray[i])
	}
	fcw.Close()

	r := NewFileColReader("data", "file_float64", reflect.Float64)
	err = r.Open()
	assert.Assert(t, !isError(err))

	for i := 0; i < 1000; i++ {
		iread, ie := r.ReadNext()
		assert.Assert(t, !isError(ie))
		assert.Assert(t, generatedIntArray[i] == iread)
	}

	//next call to read nex need to give error
	_, ie := r.ReadNext()
	assert.Assert(t, isError(ie))
}
