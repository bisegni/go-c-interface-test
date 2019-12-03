package query

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// AbstractFileTable define the base work on a table that is implementaed with fodler and files
type AbstractFileTable struct {
	//path where store the result
	fullPath string

	//contains the schema of the table/virtual table
	schema []ColDescription

	//the reader for the column
	columnReader []ColReader

	//the writer for the column
	columnWriter []ColWriter
}

func newAbstractFileTable(fullPath string) AbstractFileTable {
	return AbstractFileTable{fullPath: fullPath}
}

func (aft *AbstractFileTable) folderCheck() (bool, error) {
	return checkFilexExists(aft.fullPath)
}

func (aft *AbstractFileTable) ensureFolder() error {
	exists, err := aft.folderCheck()
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	//create folder and metadati
	err = os.MkdirAll(aft.fullPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return ErrTMSChemaMetadataNotFount
	}
	return nil
}

// Allocate the structure column writer
func (aft *AbstractFileTable) allocateColumnWriter() error {
	if aft.columnWriter != nil {
		return nil
	}
	if aft.schema == nil {
		return ErrNoSchemaInformation
	}
	// load column writer for write operation
	for _, col := range aft.schema {
		w := NewFileColWriter(aft.fullPath, col.Name, col.Kind)
		err := w.Open()
		if err == nil {
			aft.columnWriter = append(aft.columnWriter, w)
		} else {
			aft.columnWriter = nil
			return err
		}
	}
	return nil
}

// Allocate the structur column reader
func (aft *AbstractFileTable) allocateColumnReader() error {
	if aft.schema == nil {
		return ErrNoSchemaInformation
	}
	// load column writer for write operation
	for _, col := range aft.schema {
		r := NewFileColReader(aft.fullPath, col.Name, col.Kind)
		err := r.Open()
		if err == nil {
			aft.columnReader = append(aft.columnReader, r)
		} else {
			aft.columnReader = nil
			return err
		}
	}
	return nil
}

// Create impl.
func (aft *AbstractFileTable) writeSchema(schema *[]ColDescription) error {
	err := aft.ensureFolder()
	if err != nil {
		return err
	}
	//create schema file in json
	jsonInfo, err := json.Marshal(*schema)
	if err == nil {
		ioutil.WriteFile(filepath.Join(aft.fullPath, "metadata.json"), jsonInfo, os.ModePerm)
	}
	return err
}

// GetSchema impl.
func (aft *AbstractFileTable) loadSchema() error {
	if aft.schema != nil {
		return nil
	}
	jsonFile, err := os.Open(filepath.Join(aft.fullPath, "metadata.json"))
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	json.Unmarshal(byteValue, &aft.schema)
	return nil
}

// Delete impl.
func (aft *AbstractFileTable) Delete() error {
	os.RemoveAll(aft.fullPath)
	return nil
}
