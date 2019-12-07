package http

import "testing"

import "gotest.tools/assert"

func TestDownload(t *testing.T) {
	err := Download("https://bisegni.s3.us-east-2.amazonaws.com/FL_insurance_sample.csv.zip", "download/file.csv")
	assert.Assert(t, err == nil)
}
