package fileutil

import "os"

// CheckFileExists check if the file exists
func CheckFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CheckForMaxSize check if file has reached the maximun size
func CheckForMaxSize(f *os.File, maxSize int64) (bool, error) {
	fi, err := f.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		return false, err
	}
	return (fi.Size() >= maxSize), nil
}
