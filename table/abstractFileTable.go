package table

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	fileutil "github.com/bisegni/go-c-interface-test/utility"
)

// AbstractFileTable define the base work on a table that is implemented with folder and files
type AbstractFileTable struct {
	// path where store the result
	fullPath string

	// contains the schema of the table/virtual table
	schema []ColDescription

	// Statistic column value
	stat StatisticResult
	// statistic mutex
	statMux sync.Mutex
}

func newAbstractFileTable(fullPath string) AbstractFileTable {
	return AbstractFileTable{fullPath: fullPath}
}

func (aft *AbstractFileTable) folderCheck() (bool, error) {
	return fileutil.CheckFileExists(aft.fullPath)
}

func (aft *AbstractFileTable) ensureFolder() error {
	exists, err := aft.folderCheck()
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	//create folder and metadata
	err = os.MkdirAll(aft.fullPath, os.ModeDir|os.ModePerm)
	if err != nil {
		return ErrTMSchemaMetadataNotFount
	}
	return nil
}

// Allocate the structure column writer
func (aft *AbstractFileTable) allocateColumnWriter() (*[]ColWriter, error) {
	var columnWriter []ColWriter
	if aft.schema == nil {
		return nil, ErrNoSchemaInformation
	}
	// load column writer for write operation
	for _, col := range aft.schema {
		w := NewFileColWriter(aft.fullPath, col.Name, col.Kind)
		err := w.Open()
		if err == nil {
			columnWriter = append(columnWriter, w)
		} else {
			return nil, err
		}
	}
	return &columnWriter, nil
}

// Allocate the structure column reader
func (aft *AbstractFileTable) allocateColumnReader() (*[]ColReader, error) {
	var columnReader []ColReader
	if aft.schema == nil {
		return nil, ErrNoSchemaInformation
	}
	// load column writer for write operation
	for _, col := range aft.schema {
		r := NewFileColReader(aft.fullPath, col.Name, col.Kind)
		err := r.Open()
		if err == nil {
			columnReader = append(columnReader, r)
		} else {
			return nil, err
		}
	}
	return &columnReader, nil
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

func (aft *AbstractFileTable) addStat(col *ColDescription, colValue interface{}) {
	aft.statMux.Lock()
	aft.stat.column = append(aft.stat.column, *col)
	aft.stat.values = append(aft.stat.values, colValue)
	aft.statMux.Unlock()
}

// GetStatistics impl.
func (aft *AbstractFileTable) GetStatistics() *StatisticResult {
	aft.statMux.Lock()
	var stat StatisticResult
	stat.column = []ColDescription{{"AbstractFileTable", reflect.Bool}}
	stat.column = append(stat.column, aft.stat.column...)

	stat.values = []interface{}{true}
	stat.values = append(stat.values, aft.stat.values...)
	aft.statMux.Unlock()
	return &stat
}

// Close all structure for table management
func (aft *AbstractFileTable) Close() {

}
