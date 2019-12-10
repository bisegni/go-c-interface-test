package http

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/cavaliercoder/grab"
)

// Download a file via http get protocol
func Download(config interface{}, uri string, destPath string) (err error) {
	//create path directory
	os.MkdirAll(destPath, os.ModePerm)
	//get file name from uri
	fileName := path.Base(uri)

	client := grab.NewClient()
	req, err := grab.NewRequest(filepath.Join(destPath, fileName), uri)
	if err != nil {
		return err
	}
	req.NoCreateDirectories = false
	resp := client.Do(req)

	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			log.Printf("%.02f%% complete\n", resp.Progress())

		case <-resp.Done:
			if err := resp.Err(); err != nil {
				return err
			}
			return nil
		}
	}
}
