package s3

import "testing"

import "gotest.tools/assert"

import "os"

func TestDownload(t *testing.T) {
	defer os.RemoveAll("download")
	c := Config{Region: "us-east-2",
		Auth: AWSAuth{ID: "AKIAJ3GUCQLZZA7E45WA",
			Secret: "twJMRCqGs3nU/x4dxBD3e5dX9MPnASVphoU57A2e"},
	}
	err := Download(&c, "s3://bisegni/FL_insurance_sample.csv.zip", "download")
	assert.Assert(t, err == nil)
	err = Download(&c, "s3://bisegni/test-folder/FL_insurance_sample.csv.zip", "download/test-folder")
	assert.Assert(t, err == nil)
}
