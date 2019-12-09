package s3

import (
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/bisegni/go-c-interface-test/adapter"
)

// Download a file via http get protocol
func Download(config interface{}, uri string, destPath string) (err error) {
	c, ok := (config).(*adapter.ExternalSource)
	if !ok {
		return nil
	}
	os.MkdirAll(destPath, os.ModePerm)
	//get file name from uri
	fileName := path.Base(uri)
	currentDestFile, err := os.Create(filepath.Join(destPath, fileName))
	if err != nil {
		return
	}
	defer currentDestFile.Close()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String(c.Region),
		Credentials: credentials.NewStaticCredentials(
			c.ExternalAuth.ID,     // id
			c.ExternalAuth.Secret, // secret
			"")},
	)
	objectInput, err := getObjectByURL(uri)
	if err != nil {
		return
	}

	downloader := s3manager.NewDownloader(session, func(d *s3manager.Downloader) {
		d.PartSize = 64 * 1024 * 1024 // 64MB per part
		d.Concurrency = 6
	})

	numBytes, err := downloader.Download(currentDestFile, objectInput)
	if err != nil {
		return
	}

	log.Println("Downloaded", currentDestFile.Name(), numBytes, "bytes")
	return nil
}

func getObjectByURL(uri string) (*s3.GetObjectInput, error) {
	var obj s3.GetObjectInput
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	log.Printf("proto: %q, bucket: %q, key: %q\n", u.Scheme, u.Host, u.Path)
	obj.Bucket = &u.Host
	obj.Key = &u.Path
	return &obj, nil
}
