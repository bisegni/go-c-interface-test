package rscopy

import "github.com/bisegni/go-c-interface-test/adapter"

type direction int

const (
	// FROM specify to copy external data into table
	FROM direction = 0
	// TO specify to copy table data into external basket
	TO direction = 1
)

type format int

const (
	// CSV specify the comma separated value format
	CSV format = 0
	// FIXEDWIDTH specify that we have a textual file and each column width is a fixed length
	FIXEDWIDTH format = 1
	// AVRO specify the Apache Avro serialization system
	AVRO format = 2
	// JSON is the javascript object notation
	JSON format = 3
	// PARQUET is the Apache Parquet storage format
	PARQUET format = 4
)

type compression int

const (
	// BZIP2 bzip2 format
	BZIP2 compression = 0
	// GZIP gzip format
	GZIP compression = 1
	// LZOP lzop format
	LZOP compression = 2
	// ZSTD Zstandard format
	ZSTD compression = 3
)

// SourceFormat is a set of information to encode/decode content of external storage
type SourceFormat struct {
	// Format maybe CSV, DELIMITER, FIXEDWIDTH, AVRO, JSON and PARQUET
	Type format
	// ExternalAuth contains the authentication token or password to access the exteranl source
	Compression compression
	// Option is the optional parameter
	Option string
	// Quote is the character used to enclouse fields
	Quote string
	// Delimiter is a single ASCII character that is used to separate fields
	Delimiter string
	// Readratio is a parameter for DynamoDB
	Readratio int
	// Timeformat is the text representation of date and time
	Timeformat string
	// ExplicitIds is an option
	ExplicitIds bool
	// Escape if the backslash is an escaping char
	Escape bool
}

// RsCopy represent the COPY statement (parametrs and options)
type RsCopy struct {
	// Table specify the internal database table involved in the copy operation
	Table string
	// Fields contains the list of fields to write or to read
	Fields []string
	// Direction specify if the operation is COPY FROM or COPY TO
	Direction direction
	// Source is the external storage
	Source adapter.ExternalSource
	// Format specify the content encoding and options
	Format SourceFormat
}
