package query

import (
	"testing"
	"os"
)
func TestChunkRotation(t *testing.T) {
	rw := newRotateWriter("chunk_test","fileanme")
	defer os.RemoveAll("chunk_test")

	rw.Rotate()
}
