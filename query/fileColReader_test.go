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

			break
		case reflect.Int32:
			generatedIntArray[i] = rand.Int31()
			break
		case reflect.Int64:
			generatedIntArray[i] = rand.Int63()
			break
		case reflect.Float32:
			generatedIntArray[i] = rand.Float32()
			break
		case reflect.Float64:
			generatedIntArray[i] = rand.Float64()
			break
		case reflect.String:
			// var i64 int64 = 0
			// binary.Read(r.file, binary.LittleEndian, &i64)
			// result = i64
			break
		}

		binary.Write(f, binary.LittleEndian, generatedIntArray[i])
	}
	return generatedIntArray
}

func TestOpenNoFoundFile(t *testing.T) {
	r := NewFileColReader("file_not_present.bin", reflect.Bool)
	err := r.Open()
	assert.Assert(t, isError(err))
}

func TestOpenFoundFile(t *testing.T) {
	//create file for test
	var file, err = os.Create("file_present.bin")
	assert.Assert(t, !isError(err))
	defer func() {
		file.Close()
		os.Remove("file_present.bin")
	}()
	r := NewFileColReader("file_present.bin", reflect.Bool)
	err = r.Open()
	assert.Assert(t, !isError(err))
}

func TestReadBool(t *testing.T) {
	//create file for test
	var generatedIntArray [1000]bool
	var file, err = os.Create("file_bool.bin")
	assert.Assert(t, !isError(err))
	defer func() {
		file.Close()
		os.Remove("file_bool.bin")
	}()

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = (rand.Intn(2) != 0)
		binary.Write(file, binary.LittleEndian, generatedIntArray[i])
	}
	file.Close()

	r := NewFileColReader("file_bool.bin", reflect.Bool)
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
	//create file for test
	var generatedIntArray [1000]int32
	var file, err = os.Create("file_int32.bin")
	assert.Assert(t, !isError(err))
	defer func() {
		file.Close()
		os.Remove("file_int32.bin")
	}()

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = rand.Int31()
		binary.Write(file, binary.LittleEndian, generatedIntArray[i])
	}
	file.Close()

	r := NewFileColReader("file_int32.bin", reflect.Int32)
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
	//create file for test
	var generatedIntArray [1000]int64
	var file, err = os.Create("file_int64.bin")
	assert.Assert(t, !isError(err))
	defer func() {
		file.Close()
		os.Remove("file_int64.bin")
	}()

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = rand.Int63()
		binary.Write(file, binary.LittleEndian, generatedIntArray[i])
	}
	file.Close()

	r := NewFileColReader("file_int64.bin", reflect.Int64)
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
	//create file for test
	var generatedIntArray [1000]float32
	var file, err = os.Create("file_float32.bin")
	assert.Assert(t, !isError(err))
	defer func() {
		file.Close()
		os.Remove("file_float32.bin")
	}()

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = rand.Float32()
		binary.Write(file, binary.LittleEndian, generatedIntArray[i])
	}
	file.Close()

	r := NewFileColReader("file_float32.bin", reflect.Float32)
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
	//create file for test
	var generatedIntArray [1000]float64
	var file, err = os.Create("file_float64.bin")
	assert.Assert(t, !isError(err))
	defer func() {
		file.Close()
		os.Remove("file_float64.bin")
	}()

	//write some intre data
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		generatedIntArray[i] = rand.Float64()
		binary.Write(file, binary.LittleEndian, generatedIntArray[i])
	}
	file.Close()

	r := NewFileColReader("file_float64.bin", reflect.Float64)
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
