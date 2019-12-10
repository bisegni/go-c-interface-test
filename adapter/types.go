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

// Auth is a set of information form authentication
type Auth struct {
	// IamRole must be in the form "arn:aws:iam::0123456789012:role/MyRedshiftRole"
	IamRole string
	ID      string
	Secret  string
}

// ExternalSource is a set of information to identify the external storage
type ExternalSource struct {
	// SourceURI is the path of external object (ie "s3://mybucket/data/nlTest2.txt")
	SourceURI string
	// ExternalAuth contains the authentication token or password to access the exteranl source
	ExternalAuth Auth
	// Region is the AWS Region in which the buckets resides
	Region string
	// SSH specify if the URI is an ssh manifest
	SSH bool
	// Manifest specify that the URI contains a manifest file
	Manifest bool
	// Encrypted specify that the URI contains an encrypted file
	Encrypted bool
}
