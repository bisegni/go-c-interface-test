package http

import (
	"log"
	"time"

	"github.com/cavaliercoder/grab"
)

// Download a file via http get protocol
func Download(url string, destPath string) error {
	client := grab.NewClient()
	req, err := grab.NewRequest(destPath, url)
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
