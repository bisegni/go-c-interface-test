package query

import "testing"

import "gotest.tools/assert"

func TestFRSWitFOlderNotFound(t *testing.T) {
	_, err := NewFileResultSet("badpath", "inesistent-fodler")
	assert.Assert(t, isError(err))
}
