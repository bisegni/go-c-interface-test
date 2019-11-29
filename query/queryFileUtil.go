package query

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

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

type rotateWriter struct {
	lock         sync.Mutex
	path         string
	Filename     string `json:"name"`
	CurrentIdnex int32  `json:"cur_index"`
	currentFile  *os.File
}

func newRotateWriter(path string, filename string) *rotateWriter {
	w := &rotateWriter{path: path, Filename: filename}
	w.init()
	err := w.Rotate()
	if err != nil {
		return nil
	}
	return w
}

func (w *rotateWriter) init() (err error) {
	//load data from file
	err = os.MkdirAll(w.path, os.ModeDir|os.ModePerm)
	if err != nil {
		return ErrTMSChemaMetadataNotFount
	}
	jsonFile, err := os.Open(filepath.Join(w.path, fmt.Sprintf("%s.chunk", w.Filename)))
	// if we os.Open returns an error then handle it
	if err != nil {
		return
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	json.Unmarshal(byteValue, w)

	//open file if we have got somethig
	w.currentFile, err = os.OpenFile(w.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	return
}

func (w *rotateWriter) peristCunkInfo() (err error) {
	jsonInfo, err := json.Marshal(w)
	if err == nil {
		err = ioutil.WriteFile(filepath.Join(w.path, fmt.Sprintf("%s.chunk", w.Filename)), jsonInfo, os.ModePerm)
	}
	return
}

func (w *rotateWriter) Write(data interface{}) error {
	return binary.Write(w.currentFile, binary.LittleEndian, &data)
}

func (w *rotateWriter) Rotate() (err error) {
	var rotate bool = false
	// Close existing file if open
	if w.currentFile != nil {
		rotate, err = checkForMaxSize(w.currentFile, columnChunkFileSize)
		if rotate {
			//we have reacehd max size
			err = w.currentFile.Close()
			w.currentFile = nil
			if err != nil {
				return
			}
		}
	} else {
		rotate = true
	}

	if !rotate {
		return
	}
	// Create a new file.
	w.CurrentIdnex = w.CurrentIdnex + 1
	w.currentFile, err = os.Create(filepath.Join(w.path, fmt.Sprintf("%s.%d", w.Filename, w.CurrentIdnex)))
	w.peristCunkInfo()
	return
}
