package adapter

// Dowloader interface
type Dowloader interface {
	// Execute the download of the file identified by uri into path
	// the config interface need to be a pointer to a struct that
	// configure the driver for download
	Download(config interface{}, uri string, destPath string) (err error)
}

// Uploader interface
type Uploader interface {
	// Download(c *Config, uri string, destPath string) (err error)
}
