package query

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileTableSchemaManagement manage a table using folder and files
type FileTableSchemaManagement struct {
	//path where store the result
	tableFolderPath string
}

// NewFileTableSchemaManagement allocate new instance
func NewFileTableSchemaManagement(_path string, _name string) *FileTableSchemaManagement {
	return &FileTableSchemaManagement{
		tableFolderPath: filepath.Join(_path, _name),
	}
}

func (ft *FileTableSchemaManagement) folderCheck(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// Create impl.
func (ft *FileTableSchemaManagement) Create(schema *[]ColDescription) error {
	exists, err := ft.folderCheck(ft.tableFolderPath)
	if err != nil {
		return err
	}

	if exists {
		return ErrTMTbaleAlredyExists
	}

	//create folder and metadati
	err = os.MkdirAll(ft.tableFolderPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return err
	}

	//create schema file in json
	jsonInfo, err := json.Marshal(*schema)
	if err == nil {
		ioutil.WriteFile(filepath.Join(ft.tableFolderPath, "metadata.json"), jsonInfo, os.ModePerm)
	}
	return err
}

// Delete impl.
func (ft *FileTableSchemaManagement) Delete() error {
	os.RemoveAll(ft.tableFolderPath)
	return nil
}

// GetSchema impl.
func (ft *FileTableSchemaManagement) GetSchema() (*[]ColDescription, error) {
	var result []ColDescription
	jsonFile, err := os.Open(filepath.Join(ft.tableFolderPath, "metadata.json"))
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteValue, &result)
	return &result, nil
}
