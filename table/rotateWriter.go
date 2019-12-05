package table

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	fileutil "github.com/bisegni/go-c-interface-test/utility"
)

type rotateWriter struct {
	lock          sync.Mutex
	path          string
	chunkInfoFile string
	Filename      string `json:"name"`
	TotalIndex    int32  `json:"total_index"`
	currentFile   *os.File
}

func newRotateWriter(path string, filename string) (w *rotateWriter, err error) {
	w = &rotateWriter{path: path, Filename: filename, chunkInfoFile: filepath.Join(path, fmt.Sprintf("%s.chunk", filename))}
	err = w.init()
	if err != nil {
		return nil, err
	}
	err = w.Rotate()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func newRotateWriterNoInit(path string, filename string) *rotateWriter {
	return &rotateWriter{path: path, Filename: filename, chunkInfoFile: filepath.Join(path, fmt.Sprintf("%s.chunk", filename))}
}

func (w *rotateWriter) init() (err error) {
	//load data from file
	err = os.MkdirAll(w.path, os.ModeDir|os.ModePerm)
	if err != nil {
		return ErrTMSchemaMetadataNotFount
	}
	exists, err := fileutil.CheckFileExists(w.chunkInfoFile)
	if err != nil {
		return
	}

	if !exists {
		return nil
	}
	jsonFile, err := os.Open(w.chunkInfoFile)
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

	//open file if we have got something
	w.currentFile, err = os.OpenFile(filepath.Join(w.path, fmt.Sprintf("%s.%d", w.Filename, w.TotalIndex)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	return
}

func (w *rotateWriter) persistChunkInfo() (err error) {
	jsonInfo, err := json.Marshal(w)
	if err == nil {
		err = ioutil.WriteFile(filepath.Join(w.path, fmt.Sprintf("%s.chunk", w.Filename)), jsonInfo, os.ModePerm)
	}
	return
}

func (w *rotateWriter) Write(data interface{}) (err error) {
	return binary.Write(w.currentFile, binary.LittleEndian, data)
}

func (w *rotateWriter) Rotate() (err error) {
	var rotate bool = false
	// Close existing file if open
	if w.currentFile != nil {
		rotate, err = fileutil.CheckForMaxSize(w.currentFile, columnChunkFileSize)
		if rotate {
			// we have reached max size
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
	w.TotalIndex = w.TotalIndex + 1
	w.currentFile, err = os.Create(filepath.Join(w.path, fmt.Sprintf("%s.%d", w.Filename, w.TotalIndex)))
	w.persistChunkInfo()
	return
}

func (w *rotateWriter) Close() error {
	if w.currentFile == nil {
		return nil
	}
	return w.currentFile.Close()
}
