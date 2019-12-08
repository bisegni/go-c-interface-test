package adapter

// Dowloader interface
type Dowloader interface {
	Download(c *Config, uri string, destPath string) (err error)
}

// Uploader interface
type Uploader interface {
	// Download(c *Config, uri string, destPath string) (err error)
}
