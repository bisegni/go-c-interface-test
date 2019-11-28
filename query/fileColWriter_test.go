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

func TestCWOpenNoFoundFile(t *testing.T) {
	w := NewFileColWriter("file_not_present.bin", reflect.Bool)
	defer os.Remove("file_not_present.bin")
	err := w.Open()
	assert.Assert(t, !isError(err))
}

func TestCWOpenFoundFile(t *testing.T) {
	//create file for test
	var file, err = os.Create("file_present.bin")
	assert.Assert(t, !isError(err))
	defer func() {
		file.Close()
		os.Remove("file_present.bin")
	}()
	r := NewFileColWriter("file_present.bin", reflect.Bool)
	err = r.Open()
	assert.Assert(t, !isError(err))
}

func TestCWWriteWrongTypeOnInsert(t *testing.T) {
	//create file for test
	r := NewFileColWriter("file_data.bin", reflect.Bool)
	e := r.Open()
	defer func() {
		r.Close()
		os.Remove("file_data.bin")
	}()

	assert.Assert(t, !isError(e))

	i32 := rand.Int31()
	e = r.Write(i32)
	assert.Assert(t, isError(e))
}

func TestCWWriteBool(t *testing.T) {
	//create file for test
	var generatedIntArray [1000]bool
	defer func() {
		os.Remove("file_data.bin")
	}()

	r := NewFileColWriter("file_data.bin", reflect.Bool)
	e := r.Open()
	assert.Assert(t, !isError(e))

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		b := (rand.Intn(2) != 0)
		ie := r.Write(b)
		assert.Assert(t, !isError(ie))
		generatedIntArray[i] = b
	}

	//read data form file
	file, e := os.Open("file_data.bin")
	assert.Assert(t, !isError(e))
	var b bool = false
	var i int32 = 0
	for i = 0; i < 1000; i++ {
		e = binary.Read(file, binary.LittleEndian, &b)
		if e != nil {
			break
		}
		assert.Equal(t, generatedIntArray[i], b)
	}
	assert.Assert(t, (i == 1000))
	file.Close()
}

func TestCWWriteInt32(t *testing.T) {
	//create file for test
	var generatedArray [1000]int32
	defer func() {
		os.Remove("file_data.bin")
	}()

	r := NewFileColWriter("file_data.bin", reflect.Int32)
	e := r.Open()
	assert.Assert(t, !isError(e))

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		i32 := rand.Int31()
		ie := r.Write(i32)
		assert.Assert(t, !isError(ie))
		generatedArray[i] = i32
	}

	//read data form file
	file, e := os.Open("file_data.bin")
	assert.Assert(t, !isError(e))
	var i32 int32 = 0
	var i int32 = 0
	for i = 0; i < 1000; i++ {
		e = binary.Read(file, binary.LittleEndian, &i32)
		if e != nil {
			break
		}
		assert.Equal(t, generatedArray[i], i32)
	}
	assert.Assert(t, i == 1000)
	file.Close()
}

func TestCWWriteInt64(t *testing.T) {
	//create file for test
	var generatedArray [1000]int64
	defer func() {
		os.Remove("file_data.bin")
	}()

	r := NewFileColWriter("file_data.bin", reflect.Int64)
	e := r.Open()
	assert.Assert(t, !isError(e))

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		i64 := rand.Int63()
		ie := r.Write(i64)
		assert.Assert(t, !isError(ie))
		generatedArray[i] = i64
	}

	//read data form file
	file, e := os.Open("file_data.bin")
	assert.Assert(t, !isError(e))
	var i64 int64 = 0
	var i int32 = 0
	for i = 0; i < 1000; i++ {
		e = binary.Read(file, binary.LittleEndian, &i64)
		if e != nil {
			break
		}
		assert.Equal(t, generatedArray[i], i64)
	}
	assert.Assert(t, i == 1000)
	file.Close()
}

func TestCWWriteFloat32(t *testing.T) {
	//create file for test
	var generatedArray [1000]float32
	defer func() {
		os.Remove("file_data.bin")
	}()

	r := NewFileColWriter("file_data.bin", reflect.Float32)
	e := r.Open()
	assert.Assert(t, !isError(e))

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		f32 := rand.Float32()
		ie := r.Write(f32)
		assert.Assert(t, !isError(ie))
		generatedArray[i] = f32
	}

	//read data form file
	file, e := os.Open("file_data.bin")
	assert.Assert(t, !isError(e))
	var f32 float32 = 0
	var i int32 = 0
	for i = 0; i < 1000; i++ {
		e = binary.Read(file, binary.LittleEndian, &f32)
		if e != nil {
			break
		}
		assert.Equal(t, generatedArray[i], f32)
	}
	assert.Assert(t, i == 1000)
	file.Close()
}

func TestCWWriteFloat64(t *testing.T) {
	//create file for test
	var generatedArray [1000]float64
	defer func() {
		os.Remove("file_data.bin")
	}()

	r := NewFileColWriter("file_data.bin", reflect.Float64)
	e := r.Open()
	assert.Assert(t, !isError(e))

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 1000; i++ {
		f64 := rand.Float64()
		ie := r.Write(f64)
		assert.Assert(t, !isError(ie))
		generatedArray[i] = f64
	}

	//read data form file
	file, e := os.Open("file_data.bin")
	assert.Assert(t, !isError(e))
	var f64 float64 = 0
	var i int32 = 0
	for i = 0; i < 1000; i++ {
		e = binary.Read(file, binary.LittleEndian, &f64)
		if e != nil {
			break
		}
		assert.Equal(t, generatedArray[i], f64)
	}
	assert.Assert(t, i == 1000)
	file.Close()
}

func TestAppend(t *testing.T) {
	r := NewFileColWriter("file_data.bin", reflect.Bool)
	e := r.Open()
	assert.Assert(t, !isError(e))
	r.Write(false)

	r1 := NewFileColWriter("file_data.bin", reflect.Bool)
	defer r1.Close()
	e = r1.Open()
	assert.Assert(t, !isError(e))
	r1.Write(true)

	fileReadTest, e := os.Open("file_data.bin")
	assert.Assert(t, !isError(e))
	defer fileReadTest.Close()

	var readedBool bool
	e = binary.Read(fileReadTest, binary.LittleEndian, &readedBool)
	assert.Assert(t, !isError(e))
	assert.Equal(t, readedBool, false)
	e = binary.Read(fileReadTest, binary.LittleEndian, &readedBool)
	assert.Assert(t, !isError(e))
	assert.Equal(t, readedBool, true)
}
