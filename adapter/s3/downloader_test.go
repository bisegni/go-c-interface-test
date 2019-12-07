package s3

import "testing"

import "gotest.tools/assert"

import "os"

func TestDownload(t *testing.T) {
	defer os.RemoveAll("download")
	c := Config{Region: "region",
		Auth: AWSAuth{ID: "ID",
			Secret: "secret"},
	}
	err := Download(&c, "s3://bisegni/FL_insurance_sample.csv.zip", "download")
	assert.Assert(t, err == nil)
	err = Download(&c, "s3://bisegni/test-folder/FL_insurance_sample.csv.zip", "download/test-folder")
	assert.Assert(t, err == nil)
}
