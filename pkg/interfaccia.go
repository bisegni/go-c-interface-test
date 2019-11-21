package pkg

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/lithammer/shortuuid"
)

// QueryExecution is the main reference to the context of an executed query
type QueryExecution struct {
	// (private) uuid is the internal reference id
	uuid string
	// Query conteins the text of the original query
	Query string
	// ColCount is the number of columns in the result set
	ColCount int
	// RowCount is the number of rows in the result set
	RowCount int
}

// QueryExecutor takes the query, execute it ad create a full context to access execution result
func QueryExecutor(sql string) (QueryExecution, error) {
	u := shortuuid.New()
	q := QueryExecution{
		uuid:     u,
		Query:    sql,
		ColCount: 0,
		RowCount: 0,
	}
	return q, nil
}

// GetSchema return the list of column names and types extracted by the executed query
func (q *QueryExecution) GetSchema() ([]string, []*ColumnType, error) {
	n, errN := q.GetResultsetColNames()
	if errN != nil {
		return nil, nil, errN
	}
	t, errT := q.GetResultsetColTypes()
	if errT != nil {
		return nil, nil, errT
	}
	return n, t, nil
}

// GetResultsetColNames return the list of column names extracted by the executed query
// Is the clone of resultset.Columns() provided by the standard mysql drive
func (q *QueryExecution) GetResultsetColNames() ([]string, error) {
	res := make([]string, q.ColCount)
	for index := 0; index < q.ColCount; index++ {
		res[index] = fmt.Sprintf("Colonna_%d", index)
	}
	return res, nil
}

// GetResultsetColTypes return the list of column types extracted by the executed query
// Is the clone of resultset.Columns() provided by the standard mysql drive
func (q *QueryExecution) GetResultsetColTypes() ([]*ColumnType, error) {
	res := make([]*ColumnType, q.ColCount)
	for index := 0; index < q.ColCount; index++ {
		c := ColumnType{
			name: fmt.Sprintf("NOME_%d", index),

			hasNullable:       true,
			hasLength:         false,
			hasPrecisionScale: false,

			nullable:     false,
			length:       0,
			databaseType: "string",
			precision:    0,
			scale:        0,
			scanType:     reflect.TypeOf(""),
		}
		res[index] = &c
	}
	return res, nil
}

// Next is che clone of resultset.Next() provided by the standard mysql drive.
// Every call to Scan, even the first one, must be preceded by a call to Next.
func (q *QueryExecution) Next() bool {
	// Next prepares the next result row for reading with the Scan method. It
	// returns true on success, or false if there is no next result row or an error
	// happened while preparing it. Err should be consulted to distinguish between
	// the two cases.

	ok := false
	// withLock(rs.closemu.RLocker(), func() {
	// 	doClose, ok = rs.nextLocked()
	// })
	// if doClose {
	// 	rs.Close()
	// }
	return ok
}

// Close is che clone of resultset.Close() provided by the standard mysql drive.
func (q *QueryExecution) Close() error {
	// Close closes the Rows, preventing further enumeration. If Next is called
	// and returns false and there are no further result sets,
	// the Rows are closed automatically and it will suffice to check the
	// result of Err. Close is idempotent and does not affect the result of Err.
	return nil
}

// Scan is che clone of resultset.Scan() provided by the standard mysql drive.
func (q *QueryExecution) Scan(dest ...interface{}) error {
	// Scan copies the columns in the current row into the values pointed
	// at by dest. The number of values in dest must be the same as the
	// number of columns in Rows.
	//
	// Scan converts columns read from the database into the following
	// common Go types and special types provided by the sql package:
	//
	//    *string
	//    *[]byte
	//    *int, *int8, *int16, *int32, *int64
	//    *uint, *uint8, *uint16, *uint32, *uint64
	//    *bool
	//    *float32, *float64
	//    *interface{}
	//    *RawBytes
	//    *Rows (cursor value)
	//    any type implementing Scanner (see Scanner docs)
	//
	// In the most simple case, if the type of the value from the source
	// column is an integer, bool or string type T and dest is of type *T,
	// Scan simply assigns the value through the pointer.
	//
	// Scan also converts between string and numeric types, as long as no
	// information would be lost. While Scan stringifies all numbers
	// scanned from numeric database columns into *string, scans into
	// numeric types are checked for overflow. For example, a float64 with
	// value 300 or a string with value "300" can scan into a uint16, but
	// not into a uint8, though float64(255) or "255" can scan into a
	// uint8. One exception is that scans of some float64 numbers to
	// strings may lose information when stringifying. In general, scan
	// floating point columns into *float64.
	//
	// If a dest argument has type *[]byte, Scan saves in that argument a
	// copy of the corresponding data. The copy is owned by the caller and
	// can be modified and held indefinitely. The copy can be avoided by
	// using an argument of type *RawBytes instead; see the documentation
	// for RawBytes for restrictions on its use.
	//
	// If an argument has type *interface{}, Scan copies the value
	// provided by the underlying driver without conversion. When scanning
	// from a source value of type []byte to *interface{}, a copy of the
	// slice is made and the caller owns the result.
	//
	// Source values of type time.Time may be scanned into values of type
	// *time.Time, *interface{}, *string, or *[]byte. When converting to
	// the latter two, time.RFC3339Nano is used.
	//
	// Source values of type bool may be scanned into types *bool,
	// *interface{}, *string, *[]byte, or *RawBytes.
	//
	// For scanning into *bool, the source may be true, false, 1, 0, or
	// string inputs parseable by strconv.ParseBool.
	//
	// Scan can also convert a cursor returned from a query, such as
	// "select cursor(select * from my_table) from dual", into a
	// *Rows value that can itself be scanned from. The parent
	// select query will close any cursor *Rows if the parent *Rows is closed.

	// rs.closemu.RLock()

	// if rs.lasterr != nil && rs.lasterr != io.EOF {
	// 	rs.closemu.RUnlock()
	// 	return rs.lasterr
	// }
	// if rs.closed {
	// 	err := rs.lasterrOrErrLocked(errRowsClosed)
	// 	rs.closemu.RUnlock()
	// 	return err
	// }
	// rs.closemu.RUnlock()

	if q.ColCount == 0 {
		return errors.New("sql: Scan called without calling Next")
	}
	if len(dest) != q.ColCount {
		return fmt.Errorf("sql: expected %d destination arguments in Scan, not %d", q.ColCount, len(dest))
	}
	for index := 0; index < q.ColCount; index++ {
		var value string
		value = "DA COMPLETARE"
		err := convertAssign(dest[index], value)
		if err != nil {
			return fmt.Errorf(`sql: Scan error on column index %d, name %q: %v`, index, "NOME_COLONNA", err)
		}
	}
	return nil
}

// ScanEntireRow return the entire row of the resultset inside an array
// Returning values are not type converted
func (q *QueryExecution) ScanEntireRow() ([]interface{}, error) {
	res := make([]interface{}, q.ColCount)
	for index := 0; index < q.ColCount; index++ {
		res[index] = nil // DA COMPLETARE
	}
	return res, nil
}

// ****************************************************************************
// ****************************************************************************

var errNilPtr = errors.New("destination pointer is nil") // embedded in descriptive error

// convertAssign copies to dest the value in src, converting it if possible.
// An error is returned if the copy would result in loss of information.
// dest should be a pointer type.
func convertAssign(dest, src interface{}) error {
	// Common cases, without reflect.
	switch s := src.(type) {
	case string:
		switch d := dest.(type) {
		case *string:
			if d == nil {
				return errNilPtr
			}
			*d = s
			return nil
		case *[]byte:
			if d == nil {
				return errNilPtr
			}
			*d = []byte(s)
			return nil
		}
	default:
		return errors.New("Not enabled conversion")
	}
	return nil
}
