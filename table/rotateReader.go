package table

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	fileutil "github.com/bisegni/go-c-interface-test/utility"
)

var (
	// ErrRotateReaderNoChunkInfoFile no chunk file info has been found
	ErrRotateReaderNoChunkInfoFile = errors.New("No chunk info file available")
	// ErrRotateReaderChunkNotFound Chunk that need to exists has not been found
	ErrRotateReaderChunkNotFound = errors.New("Chunk not found")
	// ErrRotateReaderNoMoreChunk No more chunk available
	ErrRotateReaderNoMoreChunk = errors.New("No more chunk available")
)

type rotateReader struct {
	lock          sync.Mutex
	path          string
	chunkInfoFile string
	Filename      string `json:"name"`
	TotalIndex    int32  `json:"total_index"`
	currentIndex  int32
	currentFile   *os.File
}

func newRotateReader(path string, filename string) (w *rotateReader, err error) {
	w = &rotateReader{path: path, Filename: filename, chunkInfoFile: filepath.Join(path, fmt.Sprintf("%s.chunk", filename))}
	err = w.updateChunkInfo()
	if err != nil {
		return nil, err
	}
	err = w.switchNextChunk()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func newRotateReaderNoInit(path string, filename string) *rotateReader {
	return &rotateReader{path: path, Filename: filename, chunkInfoFile: filepath.Join(path, fmt.Sprintf("%s.chunk", filename))}
}

func (w *rotateReader) updateChunkInfo() (err error) {
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
		return ErrRotateReaderNoChunkInfoFile
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
	err = json.Unmarshal(byteValue, w)
	return
}

func (w *rotateReader) Read(data interface{}) error {
	err := binary.Read(w.currentFile, binary.LittleEndian, data)
	if err != nil {
		//try to change chunk
		if err = w.switchNextChunk(); err == nil {
			err = binary.Read(w.currentFile, binary.LittleEndian, data)
		}
	}
	return err
}

func (w *rotateReader) switchNextChunk() (err error) {
	if w.currentFile == nil {
		//we need to open first index
		w.currentIndex = 1
	} else {
		w.currentFile.Close()
		if w.currentIndex >= w.TotalIndex {
			return ErrRotateReaderNoMoreChunk
		}
		w.currentIndex++
	}
	// Create a new file.
	nextFilePath := filepath.Join(w.path, fmt.Sprintf("%s.%d", w.Filename, w.currentIndex))
	var chunkExists bool
	if chunkExists, err = fileutil.CheckFileExists(nextFilePath); err != nil {
		return err
	}
	if chunkExists == false {
		return ErrRotateReaderChunkNotFound
	}
	w.currentFile, err = os.Open(nextFilePath)
	return
}

func (w *rotateReader) Close() error {
	if w.currentFile == nil {
		return nil
	}
	return w.currentFile.Close()
}
