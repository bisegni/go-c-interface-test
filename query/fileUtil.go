package query

import "os"

func checkFilexExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//! check if file has reached the maximun size
func checkForMaxSize(f *os.File, maxSize int64) (bool, error) {
	fi, err := f.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		return false, err
	}
	return (fi.Size() >= maxSize), nil
}
