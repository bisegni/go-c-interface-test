package query

import (
	"encoding/binary"
	"gotest.tools/assert"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

func isError(err error) bool {
	return (err != nil)
}

func TestMain(m *testing.M) {
	//disable log during tests
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
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
