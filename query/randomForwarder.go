package query

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

// RandomForwarder create a random query result
type RandomForwarder struct {
	//path where store the result
	path     string
	colDesc  []ColDescription
	rowCount int64
}

// NewRandomForwarder allocate new instance
func NewRandomForwarder(_path string) *RandomForwarder {
	return &RandomForwarder{
		path: _path}
}

// Execute Generate random query result and metadata
func (rf *RandomForwarder) Execute() error {
	var err error
	//generate random number of column
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	//geenrate random type
	colNum := rand.Intn(9) + 1

	for i := 0; i < colNum; i++ {
		colTypeInde := rand.Intn(4)
		switch colTypeInde {
		case 0:
			rf.colDesc = append(rf.colDesc, ColDescription{"coll_" + strconv.Itoa(i), reflect.Bool})
			break
		case 1:
			rf.colDesc = append(rf.colDesc, ColDescription{"coll_" + strconv.Itoa(i), reflect.Int32})
			break
		case 2:
			rf.colDesc = append(rf.colDesc, ColDescription{"coll_" + strconv.Itoa(i), reflect.Int64})
			break
		case 3:
			rf.colDesc = append(rf.colDesc, ColDescription{"coll_" + strconv.Itoa(i), reflect.Float32})
			break
		case 4:
			rf.colDesc = append(rf.colDesc, ColDescription{"coll_" + strconv.Itoa(i), reflect.Float64})
			break
		}
	}

	// generate random row number
	rf.rowCount = rand.Int63n(1000)

	//generate the column file with random value up to 1000 elements
	for _, colDesc := range rf.colDesc {
		var fileName string
		if len(rf.path) > 0 {
			fileName = fmt.Sprintf("%s/%s", rf.path, colDesc.Name)
		} else {
			fileName = fmt.Sprintf("%s", colDesc.Name)
		}

		f, e := os.Create(fileName)
		if e != nil {
			return e
		}
		var i int64 = 0
		for ; i < rf.rowCount && err == nil; i++ {
			switch colDesc.Kind {
			case reflect.Bool:
				var b bool = (rand.Intn(2) != 0)
				err = binary.Write(f, binary.LittleEndian, &b)
				break
			case reflect.Int32:
				var i32 int32 = rand.Int31()
				err = binary.Write(f, binary.LittleEndian, &i32)
				break
			case reflect.Int64:
				var i64 int64 = rand.Int63()
				err = binary.Write(f, binary.LittleEndian, &i64)
				break
			case reflect.Float32:
				var f32 float32 = rand.Float32()
				err = binary.Write(f, binary.LittleEndian, &f32)
				break
			case reflect.Float64:
				var f64 float64 = rand.Float64()
				err = binary.Write(f, binary.LittleEndian, &f64)
				break
			}
		}
		f.Close()
	}
	return err
}

// GetSchema return the table column schema
func (rf *RandomForwarder) GetSchema() (*[]ColDescription, error) {
	return &rf.colDesc, nil
}

// GetRowCount return the number of row found
func (rf *RandomForwarder) GetRowCount() (int64, error) {
	return rf.rowCount, nil
}

// Close close the excutor and delete all file
func (rf *RandomForwarder) Close() {
	for _, colDesc := range rf.colDesc {
		var fileName string
		if len(rf.path) > 0 {
			fileName = fmt.Sprintf("%s/%s", rf.path, colDesc.Name)
		} else {
			fileName = fmt.Sprintf("%s", colDesc.Name)
		}
		os.Remove(fileName)
	}
}
