package query

import (
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// RandomDataGenerator create a random query result
type RandomDataGenerator struct {
	//path where store the result
	table    Table
	rowCount int64
}

func newRandomSchema() *[]ColDescription {
	var cd []ColDescription
	//generate random number of column
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	//geenrate random type
	colNum := rand.Intn(9) + 1

	for i := 0; i < colNum; i++ {
		colTypeInde := rand.Intn(4)
		switch colTypeInde {
		case 0:
			cd = append(cd, ColDescription{"coll_" + strconv.Itoa(i), reflect.Bool})
		case 1:
			cd = append(cd, ColDescription{"coll_" + strconv.Itoa(i), reflect.Int32})
		case 2:
			cd = append(cd, ColDescription{"coll_" + strconv.Itoa(i), reflect.Int64})
		case 3:
			cd = append(cd, ColDescription{"coll_" + strconv.Itoa(i), reflect.Float32})
		case 4:
			cd = append(cd, ColDescription{"coll_" + strconv.Itoa(i), reflect.Float64})
		}
	}
	return &cd
}

// NewRandomData allocate new instance
func NewRandomData(table Table) *RandomDataGenerator {
	return &RandomDataGenerator{
		table: table}
}

// Execute Generate random query result and metadata
func (rf *RandomDataGenerator) Execute(numberOfRow int32) error {
	var err error
	schema, err := rf.table.GetSchema()

	//generate random number of column
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate random row number
	rf.rowCount = 1000 //rand.Int63n(1000)

	row := make([]interface{}, len(*schema))

	for idx := int32(0); idx < numberOfRow; idx++ {
		for i, col := range *schema {
			switch col.Kind {
			case reflect.Bool:
				row[i] = (rand.Intn(2) != 0)
			case reflect.Int32:
				row[i] = rand.Int31()
			case reflect.Int64:
				row[i] = rand.Int63()
			case reflect.Float32:
				row[i] = rand.Float32()
			case reflect.Float64:
				row[i] = rand.Float64()
			}
		}
		err = rf.table.InsertRow(&row)
		if err != nil {
			return err
		}
	}
	return err
}
