package query

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestRotateReaderNoChunkNoInfo(t *testing.T) {
	_, err := newRotateReader("chunk_test", "filename")
	defer os.RemoveAll("chunk_test")
	assert.Assert(t, err == ErrRotateReaderNoChunkInfoFile)
}

func TestRotateReaderNoChunkInInfo(t *testing.T) {
	// err := os.MkdirAll("chunk_test", os.ModePerm)
	// assert.Assert(t, !isError(err))

	// f, err := os.Create(filepath.Join("chunk_test", fmt.Sprintf("filename.%d", w.Filename, w.TotalIndex)))
	// assert.Assert(t, !isError(err))
	rr := newRotateReaderNoInit("chunk_test", "filename")
	defer os.RemoveAll("chunk_test")

	err := rr.switchNextChunk()
	assert.Assert(t, err == ErrRotateReaderChunkNotFound)
}

func TestRotateReaderReadMultipleCHunk(t *testing.T) {
	defer os.RemoveAll("chunk_test")
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	rw, err := newRotateWriter("chunk_test", "filename")
	defer rw.Close()
	assert.Assert(t, !isError(err))

	// calculate random number of chunk
	numChunk := rand.Int31n(10) + 1
	elementToWrite := (columnChunkFileSize / 4) * numChunk
	// write data to fill up to chunk
	for idx := int32(0); idx < elementToWrite; idx++ {
		tmpWrite := rand.Int31()
		err = rw.Write(&tmpWrite)
		assert.Assert(t, !isError(err))
		err = rw.Rotate()
		assert.Assert(t, !isError(err))
	}
	//check the number of chunk created
	assert.Assert(t, rw.TotalIndex == numChunk+1)

	//try to read
	rr, err := newRotateReader("chunk_test", "filename")
	defer rr.Close()
	assert.Assert(t, !isError(err))

	var ri32 int32
	for idx := int32(0); idx < elementToWrite; idx++ {
		err := rr.Read(&ri32)
		assert.Assert(t, !isError(err))
	}
	assert.Assert(t, !isError(err))
}
