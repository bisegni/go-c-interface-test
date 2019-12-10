package s3

import (
	"os"
	"testing"

	"github.com/bisegni/go-c-interface-test/adapter"
	"gotest.tools/assert"
)

func TestDownload(t *testing.T) {
	defer os.RemoveAll("download")
	c := adapter.ExternalSource{Region: "region",
		ExternalAuth: adapter.Auth{ID: "ID",
			Secret: "Secret"},
	}
	err := Download(&c, "s3://bisegni/FL_insurance_sample.csv.zip", "download")
	assert.Assert(t, err == nil)
	err = Download(&c, "s3://bisegni/test-folder/FL_insurance_sample.csv.zip", "download/test-folder")
	assert.Assert(t, err == nil)
}
