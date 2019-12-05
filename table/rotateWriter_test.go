package table

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	fileutil "github.com/bisegni/go-c-interface-test/utility"
	"gotest.tools/assert"
)

func TestChunkRotationForWriter(t *testing.T) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	rw, err := newRotateWriter("chunk_test", "filename")
	defer os.RemoveAll("chunk_test")
	assert.Assert(t, !isError(err))

	//try a rotate to check issue
	err = rw.Rotate()
	assert.Assert(t, !isError(err))

	numChunk := rand.Int31n(10)

	//write data to fill up to chunk
	for idx := int32(0); idx < (columnChunkFileSize/8)*numChunk; idx++ {
		tmpWrite := rand.Int63()
		err = rw.Write(&tmpWrite)
		assert.Assert(t, !isError(err))
		err = rw.Rotate()
		assert.Assert(t, !isError(err))
	}
	rw.Rotate()
	assert.Assert(t, rw.TotalIndex == numChunk+1)

	//check file presence
	for idx := 0; int32(idx) < numChunk; idx++ {
		b1, e1 := fileutil.CheckFileExists("chunk_test/filename." + strconv.Itoa(idx+1))
		assert.Assert(t, !isError(e1))
		assert.Assert(t, b1 == true)
	}
}
